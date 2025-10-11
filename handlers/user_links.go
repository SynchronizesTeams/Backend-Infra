package handlers

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/models"
	"os"
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