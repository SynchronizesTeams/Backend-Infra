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

// 🟢 Create Certification
func CreateCertification(c *fiber.Ctx) error {
	var input dto.CertificationRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/certifications")
	if err != nil {
		filePath = ""
	}

	cert := models.Certification{
		Name:                input.Name,
		Issuer:              input.Issuer,
		Description:         input.Description,
		Image:               filePath,
		CertificationNumber: input.CertificationNumber,
		IssueDate:           input.IssueDate,
		ExpiryDate:          input.ExpiryDate,
	}

	if err := database.DB.Create(&cert).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create certification",
		})
	}

	response := mapper.CertToDTO(cert)

	return c.Status(200).JSON(response)
}

// 🟡 Edit Certification
func EditCertification(c *fiber.Ctx) error {
	id := c.Params("id")
	var cert models.Certification

	if err := database.DB.First(&cert, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch certification",
		})
	}

	var input dto.CertificationRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/certifications", cert.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update image",
			})
		}
		cert.Image = path
	}

	mapper.UpdateCert(&cert, input)

	if err := database.DB.Save(&cert).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save certification",
		})
	}

	response := mapper.CertToDTO(cert)

	return c.Status(200).JSON(response)
}

// 🔵 Show Certification (by ID)
func ShowCertification(c *fiber.Ctx) error {
	id := c.Params("id")
	var cert models.Certification

	if err := database.DB.First(&cert, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch certification",
		})
	}

	response := mapper.CertToDTO(cert)
	return c.Status(200).JSON(response)
}

// 🟣 Show All Certifications
func ShowAllCertification(c *fiber.Ctx) error {
	var certs []models.Certification

	if err := database.DB.Find(&certs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch certifications",
		})
	}

	response := mapper.CertListToDTO(certs)
	return c.Status(200).JSON(response)
}

// 🔴 Delete Certification
func DeleteCertification(c *fiber.Ctx) error {
	id := c.Params("id")
	var cert models.Certification

	if err := database.DB.First(&cert, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch certification",
		})
	}

	if err := os.Remove(cert.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&cert).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete certification",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "certification successfully deleted",
	})
}
