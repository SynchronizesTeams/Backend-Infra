package handlers

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/helpers"
	"go-api-infra/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreateSection(c *fiber.Ctx) error {
	name := c.FormValue("name")
	slug := helpers.Slugify(name)
	description := c.FormValue("description")

	file, err := c.FormFile("icon")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "the icon field is required",
		})
	}

	filePath := fmt.Sprintf("uploads/forum-section/%d_%s", time.Now().Unix(), file.Filename ) 
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save icon",
		})
	}

	forum_section := models.ForumSection {
		Name: name,
		Slug: slug,
		Description: description,
		Icon: filePath,
	}

	if err := database.DB.Create(&forum_section).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create forum section",
		})
	}

	response := dto.ForumSectionResponse{
		ID: forum_section.ID,
		Name: forum_section.Name,
		Slug: forum_section.Slug,
		Description: forum_section.Description,
		Icon: forum_section.Icon,
		CreatedAt: forum_section.CreatedAt,
	}

	return c.Status(201).JSON(response)
}

func EditSection(c *fiber.Ctx) error {
	id := c.Params("id")
	var forum_section models.ForumSection

	if err := database.DB.First(&forum_section, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"erorr": "failed to fetch section",
		})
	}

	name := c.FormValue("name")
	slug := helpers.Slugify(name)
	description := c.FormValue("description")

	if name != "" {
		forum_section.Name = name 
	}

	if slug != "" {
		forum_section.Slug = slug
	}

	if description != "" {
		forum_section.Description = description
	}

	file, err := c.FormFile("icon")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s",time.Now().Unix(), file.Filename)
		path := fmt.Sprintf("uploads/forum-section/%s", filename)

		if err := c.SaveFile(file, path); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to save icon",
			})
		}

		if forum_section.Icon != "" {
			_ = os.Remove(forum_section.Icon)
		}

		forum_section.Icon = path
	}

	if err := database.DB.Save(&forum_section).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update section",
		})
	}

	response := dto.ForumSectionResponse{
		ID: forum_section.ID,
		Name: forum_section.Name,
		Slug: forum_section.Slug,
		Description: forum_section.Description,
		Icon: forum_section.Icon,
		CreatedAt: forum_section.CreatedAt,
	}

	return c.Status(200).JSON(response)

}

func DeleteSection(c *fiber.Ctx) error {
	id := c.Params("id")

	var section models.ForumSection

	if err := database.DB.First(&section, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "cannot found section",
		})
	}

	if err := os.Remove(section.Icon); err != nil {
		if os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete icon",
			})
		}
	}

	if err := database.DB.Delete(&section).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete section record",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "section deleted successfully",
	})
}

func ShowAllSection(c *fiber.Ctx) error {
	var sections []models.ForumSection

	if err := database.DB.Find(&sections).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch section",
		})
	}

	var response []dto.ForumSectionResponse
	for _, section := range sections {
		response = append(response, dto.ForumSectionResponse{
			ID: section.ID,
			Name: section.Name,
			Slug: section.Slug,
			Description: section.Description,
			Icon: section.Icon,
			IsActive: section.IsActive,
		})
	}

	return c.Status(200).JSON(response)
}