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

func CreatePortal(c *fiber.Ctx) error {
	var input dto.PortalRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "logo", "uploads/portal")

	if err != nil {
		filePath = ""
	}

	portal := models.Portal{
		Name:       input.Name,
		Description: input.Description,
		URL:     input.URL,
		Logo:       filePath,
		Category: input.Category,
	}

	if err := database.DB.Create(&portal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create portal",
		})
	}

	response := mapper.PortalToDTO(portal)

	return c.Status(200).JSON(response)
}

func EditPortal(c *fiber.Ctx) error {
	id := c.Params("id")
	var portal models.Portal

	if err := database.DB.First(&portal, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var input dto.PortalRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("logo")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/portal", portal.Logo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update portal",
			})
		}
		portal.Logo = path
	}

	mapper.UpdatePortalFromDTO(&portal, input)

	if err := database.DB.Save(&portal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save achievement",
		})
	}


	response := mapper.PortalToDTO(portal)

	return c.Status(200).JSON(response)
}

func ShowAllPortal(c *fiber.Ctx) error {
	var portal []models.Portal

	if err := database.DB.Find(&portal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.PortalListToDTO(portal)

	return c.Status(200).JSON(response)
}


func ShowPortal(c *fiber.Ctx) error {
	id := c.Params("id")
	var portal models.Portal

	if err := database.DB.First(&portal, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.PortalToDTO(portal)

	return c.Status(200).JSON(response)
}

func DeletePortal(c *fiber.Ctx) error {
	id := c.Params("id")
	var portal models.Portal

	if err := database.DB.First(&portal, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	if err := os.Remove(portal.Logo); err != nil {
		if os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete logo",
			})
		}
	}

	if err := database.DB.Delete(&portal).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "portal successfully deleted",
	})
}