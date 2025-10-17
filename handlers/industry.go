package handlers

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/helpers"
	"go-api-infra/mapper"
	"go-api-infra/models"
)

// 🧩 CREATE
func CreateIndustry(c *fiber.Ctx) error {
	var input dto.IndustryRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "logo", "uploads/industry")
	if err != nil {
		filePath = ""
	}

	industry := models.Industry{
		Name:        input.Name,
		Logo:        filePath,
		Website:     input.Website,
		Description: input.Description,
	}

	if err := database.DB.Create(&industry).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create industry",
		})
	}

	response := mapper.IndustryToDTO(industry)
	return c.Status(200).JSON(response)
}

// 🧩 READ (All)
func ShowAllIndustry(c *fiber.Ctx) error {
	var industries []models.Industry

	if err := database.DB.Find(&industries).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch industries"})
	}

	response := mapper.IndustryListToDTO(industries)
	return c.Status(200).JSON(response)
}

// 🧩 READ (Single)
func ShowIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	var industry models.Industry

	if err := database.DB.First(&industry, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "industry not found"})
	}

	response := mapper.IndustryToDTO(industry)
	return c.Status(200).JSON(response)
}

// 🧩 UPDATE
func EditIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	var industry models.Industry

	if err := database.DB.First(&industry, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "industry not found"})
	}

	var input dto.IndustryRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid form data"})
	}

	file, _ := c.FormFile("logo")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/industry", industry.Logo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to update logo"})
		}
		industry.Logo = path
	}

	mapper.UpdateIndustry(&industry, input)

	if err := database.DB.Save(&industry).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to save industry"})
	}

	response := mapper.IndustryToDTO(industry)
	return c.Status(200).JSON(response)
}

// 🧩 DELETE
func DeleteIndustry(c *fiber.Ctx) error {
	id := c.Params("id")
	var industry models.Industry

	if err := database.DB.First(&industry, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "industry not found"})
	}

	// Hapus file logo jika ada
	if industry.Logo != "" {
		if err := os.Remove(industry.Logo); err != nil && !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{"error": "failed to delete logo"})
		}
	}

	if err := database.DB.Delete(&industry).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete industry"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "industry successfully deleted"})
}