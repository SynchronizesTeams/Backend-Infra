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

func CreateEskul(c *fiber.Ctx) error {
	var input dto.EskulRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/eskul")

	if err != nil {
		filePath = ""
	}

	eskul := models.Eskul{
		Name: input.Name,
		Description: input.Description,
		PembinaID: input.PembinaID,
		Image: filePath,
	}

	if err := database.DB.Create(&eskul).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create eskul",
		})
	}

	database.DB.Preload("Pembina").First(&eskul, eskul.ID)

	response := mapper.EskulToDTO(eskul)

	return c.Status(200).JSON(response)
}

func EditEskul(c *fiber.Ctx) error {
	id := c.Params("id")
	var eskul models.Eskul

	if err := database.DB.First(&eskul, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var input dto.EskulRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/eskul", eskul.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update eskul",
			})
		}
		eskul.Image = path
	}

	mapper.UpdateEskulFromDTO(&eskul, input)

	if err := database.DB.Save(&eskul).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save mading",
		})
	}

	response := mapper.EskulToDTO(eskul)

	return c.Status(200).JSON(response)
}

func ShowEskul(c *fiber.Ctx) error {
	id := c.Params("id")
	var eskul models.Eskul

	if err := database.DB.First(&eskul, id).Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "failed to fetch eskul",
		})
	}

	database.DB.Preload("Pembina").First(&eskul, id)

	response := mapper.EskulToDTO(eskul)

	return c.Status(200).JSON(response)
}

func ShowAllEskul(c *fiber.Ctx) error {
	var eskuls []models.Eskul

	if err := database.DB.Preload("Pembina").Find(&eskuls).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.EskulListToDTO(eskuls)

	return c.Status(200).JSON(response)
}

func DeleteEskul(c *fiber.Ctx) error {
	id := c.Params("id")
	var eskul models.Eskul

	if err := database.DB.First(&eskul, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"fiber": "fialed to fetch",
		})
	}

	if err := os.Remove(eskul.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&eskul).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error": "eskul successfully deleted",
	})
}