package handlers

import (
	"go-api-infra/database"
	"go-api-infra/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("mysecretkey")

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

	// Generate JWT token
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SecretKey)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"user":  user,	
		"token": t,
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
