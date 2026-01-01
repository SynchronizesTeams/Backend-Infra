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

func CreateMading(c *fiber.Ctx) error {
	var input dto.MadingRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Invalid formdata",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/madings") 
	if err != nil {
		filePath = ""
	}

	mading := models.Mading{
		Title: input.Title,
		Type: input.Type,
		Content: input.Content,
		Image: filePath,
	}

	if err := database.DB.Create(&mading).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create mading",
		})
	}

	var response = mapper.MadingToDTO(mading)

	return c.Status(200).JSON(response)
}

func EditMading(c *fiber.Ctx) error {
	id := c.Params("id")
	var mading models.Mading

	if err := database.DB.First(&mading, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var input dto.MadingRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/madings", mading.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update mading",
			})
		}
		mading.Image = path
	}

	mapper.UpdateMading(&mading, input)

	if err := database.DB.Save(&mading).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save mading",
		})
	}

	response := mapper.MadingToDTO(mading)

	return c.Status(200).JSON(response)
}

func ShowMading(c *fiber.Ctx) error {
	id := c.Params("id")
	var mading models.Mading

	if err := database.DB.First(&mading, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch mading",
		})
	}

	response := mapper.MadingToDTO(mading)

	return c.Status(200).JSON(response)
}

func ShowAllMadings(c *fiber.Ctx) error {
	var madings []models.Mading

	if err := database.DB.Find(&madings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	response := mapper.MadingListToDTO(madings)

	return c.Status(200).JSON(response)
}

func DeleteMading(c *fiber.Ctx) error {
	id := c.Params("id")
	var mading models.Mading

	if err := database.DB.First(&mading, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch mading",
		})
	}

	if err := os.Remove(mading.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&mading).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete mading",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "mading deleted successfully",
	})
}


func ChangeStatusMading(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.FormValue("status")
	
	var mading models.Mading

	if err := database.DB.First(&mading, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch mading",
		})
	}

	mading.Status = status

	if err := database.DB.Save(&mading).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update status",
		})
	}

	response := mapper.MadingToDTO(mading)

	return c.Status(200).JSON(response)
}