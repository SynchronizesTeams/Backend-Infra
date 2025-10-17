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

func EditEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find event",
		})
	}

	var input dto.EventRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	mapper.UpdateEventFromDTO(&event, input) 

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/events", event.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to save image",
			})
		}
		event.Image = path
	}

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save event",
		})
	}

	response := mapper.EventToDTO(event)

	return c.Status(200).JSON(response)
}

func DeleteEvent (c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find event ",
		})
	}

	if err := os.Remove(event.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete event",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "event successfully deleted",
	})
}


func GetEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find event",
		})
	}

	var response = mapper.EventToDTO(event)

	return c.Status(200).JSON(response)
}

func GetEventByDate(c *fiber.Ctx) error {
	date := c.Params("date")
	var event models.Event

	if err := database.DB.Where("start_date = ?", date).Find(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find event",
		})
	}

	var response = mapper.EventToDTO(event)

	return c.Status(200).JSON(response)
}

func GetAllEvent(c *fiber.Ctx) error {
	var event []models.Event

	if err := database.DB.Find(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch event",
		})
	}

	var response = mapper.EventListToDTO(event)

	return c.Status(200).JSON(response)
}

func ChangeVisibility(c *fiber.Ctx) error {
	id := c.Params("id")
	visibility := c.FormValue("visibility")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find event",
		})
	}

	event.Visibility = visibility

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to change visibility",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "visibility successfuly changed",
		"visibility": visibility,
	})
}