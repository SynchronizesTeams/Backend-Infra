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

// 🟢 Create Testimonial
func CreateTestimonial(c *fiber.Ctx) error {
	var input dto.TestimonialRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "photo", "uploads/testimonial")
	if err != nil {
		filePath = ""
	}

	testi := models.Testimonial{
		Name:        input.Name,
		Position:    input.Position,
		Photo:       filePath,
		Testimonial: input.Testimonial,
	}

	if err := database.DB.Create(&testi).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create testimonial",
		})
	}

	response := mapper.TestiToDTO(testi)
	return c.Status(200).JSON(response)
}

// 🟡 Edit Testimonial
func EditTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	var testi models.Testimonial

	if err := database.DB.First(&testi, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch testimonial",
		})
	}

	var input dto.TestimonialRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("photo")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/testimonial", testi.Photo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update photo",
			})
		}
		testi.Photo = path
	}

	mapper.UpdateTesti(&testi, input)

	if err := database.DB.Save(&testi).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save testimonial",
		})
	}

	response := mapper.TestiToDTO(testi)
	return c.Status(200).JSON(response)
}

// 🔵 Show Testimonial (by ID)
func ShowTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	var testi models.Testimonial

	if err := database.DB.First(&testi, id).Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "failed to fetch testimonial",
		})
	}

	response := mapper.TestiToDTO(testi)
	return c.Status(200).JSON(response)
}

// 🟣 Show All Testimonials
func ShowAllTestimonial(c *fiber.Ctx) error {
	var testimonials []models.Testimonial

	if err := database.DB.Find(&testimonials).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch testimonials",
		})
	}

	response := mapper.TestiListToDTO(testimonials)
	return c.Status(200).JSON(response)
}

// 🔴 Delete Testimonial
func DeleteTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	var testi models.Testimonial

	if err := database.DB.First(&testi, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch testimonial",
		})
	}

	if err := os.Remove(testi.Photo); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete photo",
			})
		}
	}

	if err := database.DB.Delete(&testi).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete testimonial",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "testimonial successfully deleted",
	})
}
