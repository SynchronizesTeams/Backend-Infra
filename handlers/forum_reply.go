package handlers

import (
	"go-api-infra/database"
	"go-api-infra/mapper"
	"go-api-infra/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func CreateReply(c *fiber.Ctx) error {
	postIDParam := c.Params("id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid post ID",
		})
	}

	content := c.FormValue("content")
	parentIDParam := c.FormValue("parent_id") // optional
	var parentID *uint

	if parentIDParam != "" {
		parentInt, err := strconv.Atoi(parentIDParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid parent ID",
			})
		}
		pid := uint(parentInt)
		parentID = &pid
	}

	// Ambil user dari JWT token
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	userID := uint(idVal)

	reply := models.ForumReply{
		PostID:   uint(postID),
		UserID:   userID,
		Content:  content,
		ParentID: parentID, // bisa null (parent), atau pointer (child)
	}

	if err := database.DB.Create(&reply).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save reply"})
	}

	// Ambil data user setelah insert
	database.DB.Preload("User").First(&reply, reply.ID)

	return c.JSON(reply)
}

func GetNestedReplies(c *fiber.Ctx) error {
	id := c.Params("id")

	var replies []models.ForumReply
	if err := database.DB.
	Preload("User").
	Preload("Children.User").
	Preload("Children.Children.User").
	Preload("Children.Children.Children.User").
	Where("post_id = ? AND parent_id IS NULL", id).
	Find(&replies).Error; err != nil {
	return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	response := mapper.ForumReplyListToDTO(replies)

	return c.Status(200).JSON(response)
}