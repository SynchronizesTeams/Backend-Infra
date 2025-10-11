package handlers

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)


func AddImage(c *fiber.Ctx) error {
	category := c.FormValue("category")
	title := c.FormValue("title")

	file, err := c.FormFile("file_path")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "image required",
		})
	}

	filePath := fmt.Sprintf("uploads/images/%d_%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save image",
		})
	}

	image := models.Image {
		Category: category,
		FilePath: filePath,
		Title: title,
	}

	if err := database.DB.Create(&image).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save image",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "image succesfully uploaded",
		"image": image,
	})
}

func DeleteImage (c *fiber.Ctx) error {
	id := c.Params("id")

	var image models.Image
	if err := database.DB.First(&image, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "image not found",
		})
	}

	if err := os.Remove(image.FilePath); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&image).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete image record",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Image deleted successfully",
	})
}

func ShowAllImages(c *fiber.Ctx) error {
	var images []models.Image

	if err := database.DB.Find(&images).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "failed to fetch images",
		})
	}

	if len(images) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "no images found",
		})
	}

	return c.Status(200).JSON(images)
}


func ShowImagesByCategory(c *fiber.Ctx) error {
	category := c.Query("category")
	var image []models.Image

	if err := database.DB.Where("category = ?", category).Find(&image).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "failed to fetch image",
		})
	}

	if len(image) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "no images found for this category",
		})
	}

	return c.Status(200).JSON(image)
}