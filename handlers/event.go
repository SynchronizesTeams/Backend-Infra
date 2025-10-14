package handlers

import (
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/helpers"
	"go-api-infra/mapper"
	"go-api-infra/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	var input dto.EventRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/events")
	if err != nil {
		filePath = ""
	} 

	event := models.Event{
		Title: input.Title,
		Category: input.Category,
		Description: input.Description,
		StartDate: input.StartDate,
		EndDate: input.EndDate,
		Location: input.Location,
		Organizer: input.Organizer,
		Image: filePath,
	}

	if err := database.DB.Create(&event).Error; err != nil {
		if event.Image != "" {
			_ = os.Remove(event.Image)
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create event",
		})
	}

	response := mapper.EventToDTO(event)

	return c.Status(200).JSON(response)
}