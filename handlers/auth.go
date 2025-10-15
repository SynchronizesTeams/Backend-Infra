package handlers

import (
	"go-api-infra/config"
	"go-api-infra/database"
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(fiber.Map{
		"user_id": claims["id"],
		"email":   claims["email"],
	})
}
