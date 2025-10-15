package handlers

import (
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/mapper"
	"go-api-infra/models"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreateGuest(c *fiber.Ctx) error {
	var input dto.GuestBookRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	guestbook := models.GuestBook{
		Title: input.Title,
		InstanceName: input.InstanceName,
		ContactPerson: input.ContactPerson,
		Email: input.Email,
		Phone: input.Phone,
		Description: input.Description,
	}

	if input.RequestDate != "" {
		parsed, err := time.Parse("2006-01-02", input.RequestDate)
		if err == nil {
			guestbook.RequestDate = parsed
		}
	}


	if err := database.DB.Create(&guestbook).Error; err != nil {	
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create guest book",
		})
	}

	response := mapper.GuestBookToDTO(&guestbook) 

	return c.Status(200).JSON(response)
}

func EditGuest(c *fiber.Ctx) error {
	id := c.Params("id")
	var guest models.GuestBook

	if err := database.DB.First(&guest, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find guest",
		})
	}

	var input dto.GuestBookRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	mapper.UpdateGuestBookFromDTO(&guest, input)

	if err := database.DB.Save(&guest).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"erros": "failed to save guest",
		})
	}

	response := mapper.GuestBookToDTO(&guest)
	
	return c.Status(200).JSON(response)
}

func GetAllGuestBook (c *fiber.Ctx) error {
	var guests []models.GuestBook	

	if err := database.DB.Find(&guests).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch guestbook",
		})
	}

	response := mapper.GuestBookListToDTO(guests)

	return c.Status(200).JSON(response)
}