package handlers

import (
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/mapper"
	"go-api-infra/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateTeacher(c *fiber.Ctx) error {
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	userToken, ok := userData.(*jwt.Token)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, _ := claims["id"].(float64)
	userID := uint(idVal)

	var input dto.TeacherRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	teacher := models.Teacher{
		NIG: input.NIG,
		FullName: input.FullName,
		Position: input.Position,
		Subject: input.Subject,
		Description: input.Description,
		UserID: userID,
	}

	if err := database.DB.Create(&teacher).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create teacher",
		})
	}

	database.DB.Preload("User").First(&teacher, teacher.ID)

	response := mapper.TeacherToDTO(teacher)

	return c.Status(201).JSON(response)
}

func EditTeacher(c *fiber.Ctx) error {
	id := c.Params("id")
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var input dto.TeacherRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	database.DB.Preload("User").First(&teacher, teacher.ID)

	mapper.UpdateTeacherFromDTO(&teacher, input)

	response := mapper.TeacherToDTO(teacher)

	return c.Status(200).JSON(response)
} 


func DeleteTeacher(c *fiber.Ctx) error {
	id := c.Params("id")
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch ",
		})
	}

	if err := os.Remove(teacher.Photo); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&teacher).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete teacher",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "teacher deleted successfully",
	})
}

func ShowAllTeachers(c *fiber.Ctx) error {
	var teachers []models.Teacher

	if err := database.DB.Preload("User").Find(&teachers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch teachers",
		})
	}

	response := mapper.TeacherListToDTO(teachers)

	return c.Status(200).JSON(response)
}

func ShowTeacher(c *fiber.Ctx) error {
	id := c.Params("id")
	var teacher models.Teacher

	if err := database.DB.Preload("User").First(&teacher, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	response := mapper.TeacherToDTO(teacher)

	return c.Status(200).JSON(response)
}
