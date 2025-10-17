package handlers

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/models"
	"os"

	// "path"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddLink(c *fiber.Ctx) error {
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	userID := uint(idVal)

	title := c.FormValue("title")
	url := c.FormValue("url")
	
	file, err := c.FormFile("icon")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "icon file required"})
	}

	os.MkdirAll("uploads/user_links", os.ModePerm)

	savePath := fmt.Sprintf("uploads/user_links/%d_%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save icons"})
	}

	user_links := models.UserLinks{
		Icon: savePath,
		Title: title,
		Url: url,
		UserID: userID,
	}

	if err := database.DB.Create(&user_links).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faile to save user links",
		})
	}	

	database.DB.Preload("User").First(&user_links, user_links.ID)
	
	return c.Status(200).JSON(fiber.Map{
		"id": user_links.ID,
		"title": user_links.Title,
		"url": user_links.Url,
		"icon": user_links.Icon,
		"user": fiber.Map{
			"id": user_links.UserID,
			"name": user_links.User.Name,
		},
	})
}

func EditLinks(c *fiber.Ctx) error {
	id := c.Params("id")
	var user_links models.UserLinks

	if err := database.DB.First(&user_links, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch user links",
		})
	}

	title := c.FormValue("title")
	url := c.FormValue("url")

	if title != "" {
		user_links.Title = title
	}

	if url != "" {
		user_links.Url = url
	}

	file, err := c.FormFile("icon")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s",time.Now().Unix(), file.Filename)
		path := fmt.Sprintf("uploads/user_links/%s", filename)

		if err := c.SaveFile(file, path); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to save icon",
			})
		}

		if user_links.Icon != "" {
			_ = os.Remove(user_links.Icon)
		}

		user_links.Icon = path
	}

	if err := database.DB.Save(&user_links).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update user links",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "user links update successfully",
		"data": fiber.Map{
			"id": user_links.ID,
			"title": user_links.Title,
			"url": user_links.Url,
			"icon": user_links.Icon,
		},
	})
}

func ShowUserLinksSelf(c *fiber.Ctx) error {
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	userID := uint(idVal)

	var user_links []models.UserLinks

	if err := database.DB.Where("user_id = ?", userID).Find(&user_links).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fectch user links",
		})
	}

	if len(user_links) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "cannot find link for this user",
		})
	}

	var response []dto.UserLinksResponse
	for _, link := range user_links {
		response = append(response, dto.UserLinksResponse{
			ID: link.ID,
			Title: link.Title,
			Url: link.Url,
			Icon: link.Icon,
		})
	}

	return c.Status(200).JSON(response)
}

func DeleteLinks(c *fiber.Ctx) error {
	id := c.Params("id")

	var user_links models.UserLinks

	if err := database.DB.First(&user_links, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch user links",
		})
	}

	if err := os.Remove(user_links.Icon); err != nil {
		if os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete icon",
			})
		}
	}


	if err := database.DB.Delete(&user_links).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user link",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "user links delete successfully",
	})
}

func ShowUserLinks(c *fiber.Ctx) error {
	id := c.Params("id")

	var user_links []models.UserLinks

	if err := database.DB.Where("user_id = ?", id).Find(&user_links).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fectch user links",
		})
	}

	if len(user_links) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "cannot find link for this user",
		})
	}

	var response []dto.UserLinksResponse
	for _, link := range user_links {
		response = append(response, dto.UserLinksResponse{
			ID: link.ID,
			Title: link.Title,
			Url: link.Url,
			Icon: link.Icon,
		})
	}

	return c.Status(200).JSON(response)
}