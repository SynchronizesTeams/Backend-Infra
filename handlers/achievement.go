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

func CreateAchievement(c *fiber.Ctx) error {
	var input dto.AchievementRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/achievement")

	if err != nil {
		filePath = ""
	}

	achievement := models.Achievement{
		Title: input.Title,
		Description: input.Description,
		NewsID: input.NewsID,
		Image: filePath,
		AchievementDate: input.AchievementDate,
	}

	if err := database.DB.Create(&achievement).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create achievement",
		})
	}

	database.DB.Preload("News").First(&achievement, achievement.ID)

	response := mapper.AchievementToDTO(achievement)

	return c.Status(200).JSON(response)
}

func EditAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	var achieve models.Achievement

	if err := database.DB.First(&achieve, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var input dto.AchievementRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/achievement", achieve.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update achievement",
			})
		}
		achieve.Image = path
	}

	mapper.UpdateAchievement(&achieve, input)

	if err := database.DB.Save(&achieve).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save achievement",
		})
	}

	database.DB.Preload("News").First(&achieve, achieve.ID)

	response := mapper.AchievementToDTO(achieve)

	return c.Status(200).JSON(response)
}

func ShowAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	var achieve models.Achievement

	if err := database.DB.First(&achieve, id).Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "failed to fetch achieve",
		})
	}

	database.DB.Preload("News").First(&achieve, id)

	response := mapper.AchievementToDTO(achieve)

	return c.Status(200).JSON(response)
}

func ShowAllAchievement(c *fiber.Ctx) error {
	var achievements []models.Achievement

	if err := database.DB.Preload("News").Find(&achievements).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}
	
	var response = mapper.AchievementListToDTO(achievements)

	return c.Status(200).JSON(response)
}

func DeleteAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	var achievement models.Achievement

	if err := database.DB.First(&achievement, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"fiber": "fialed to fetch",
		})
	}

	if err := os.Remove(achievement.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&achievement).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error": "achievement successfully deleted",
	})
}