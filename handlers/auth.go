package handlers

import (
	"go-api-infra/config"
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/helpers"
	"go-api-infra/mapper"
	"go-api-infra/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register user baru
func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse json"})
	}

	// Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create user"})
	}

	return c.JSON(fiber.Map{
		"message": "user created",
		"user": user,
	})

}

// Login
func Login(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse json"})
	}

	var user models.User
	database.DB.Where("email = ?", req.Email).First(&user)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "wrong password"})
	}

	// Generate JWT token (pindah ke helper di config/jwt.go)
	token, err := config.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

// Protected route
func Profile(c *fiber.Ctx) error {
	var user models.User
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	UserID := uint(idVal)

	if err := database.DB.First(&user, UserID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.UserToDTO(user)

	return c.Status(200).JSON(response)
}

func EditProfile(c *fiber.Ctx) error {
	var user models.User
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken, ok := userData.(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	idVal, ok := claims["id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID in token"})
	}
	UserID := uint(idVal)

	if err := database.DB.First(&user, UserID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find user",
		})
	}

	var input dto.UserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid form data",
		})
	}

	file, _ := c.FormFile("photo_url")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/user", user.PhotoUrl)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update image",
			})
		}
		user.PhotoUrl = path
	}

	mapper.UpdateUser(&user, input)

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save user",
		})
	}

	var response = mapper.UserToDTO(user)

	return c.Status(200).JSON(response)
}

func ProfilePublic(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	var response = mapper.UserToDTO(user)

	return c.Status(200).JSON(response)
}