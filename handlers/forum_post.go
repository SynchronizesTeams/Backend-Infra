package handlers

import (
	"go-api-infra/database"
	"go-api-infra/dto"
	"go-api-infra/helpers"
	"go-api-infra/mapper"
	"go-api-infra/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)


func CreatePost(c *fiber.Ctx) error {
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

	var input dto.ForumPostRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid formdata",
		})
	}

	filePath, err := helpers.SaveUploadedFile(c, "image", "uploads/forum-posts") 
	if err != nil {
		filePath = ""
	}

	post := models.ForumPost{
		Title: input.Title,
		Content: input.Content,
		Image: filePath,
		SectionID: input.SectionID,
		UserID: userID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		if post.Image != "" {
			_ = os.Remove(post.Image)
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create post",
		})
	}

	database.DB.Preload("User").Preload("Section").First(&post, post.ID)

	response := mapper.ForumPostToDTO(post)

	return c.Status(201).JSON(response)
}

func EditPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.ForumPost

	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "failed to fetch post",
		})
	}

	var input dto.ForumPostRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	helpers.UpdateForumPostFromDTO(&post, input)

	file, _ := c.FormFile("image")
	if file != nil {
		path, err := helpers.SaveUpdatedFile(c, file, "uploads/forum-posts", post.Image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update post",
			})
		}
		post.Image = path
	}

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save forum post",
		})
	}

	response := mapper.ForumPostToDTO(post)

	return c.Status(200).JSON(response)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.ForumPost

	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "cannot find post",
		})
	}

	if err := os.Remove(post.Image); err != nil {
		if !os.IsNotExist(err) {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to delete image",
			})
		}
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete post",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "post deleted successfully",
	})
}

func GetAllPost(c *fiber.Ctx) error {
	var posts []models.ForumPost

	if err := database.DB.Find(&posts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch post data",
		})
	} 

	database.DB.Preload("User").Preload("Section").Find(&posts)
	
	response := mapper.ForumPostListToDTO(posts)
	return c.Status(200).JSON(response)
}

func ChangePostStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.FormValue("status")

	var post models.ForumPost

	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cannot find post",
		})
	}

	post.Status = status

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to change status",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "status successfully changed",
		"status": status,
	})
}

func ShowPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.ForumPost

	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot find post",
		})
	}

	database.DB.Preload("User").Preload("Section").Find(&post)

	response := mapper.ForumPostToDTO(post)
	return c.Status(200).JSON(response)
}


func UpvotePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.ForumPost

	if err :=  database.DB.First(&post, id).Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "failed to fetch post",
		})
	}

	post.Upvote += 1

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to add upvote",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "post successfully upvote",
		"upvote": post.Upvote,
	})
}

func DownvotePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.ForumPost

	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "failed to fetch",
		})
	}

	post.Downvote += 1

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to save downvote",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "successfully downvote post",
		"downvote": post.Downvote,
	})
}

func GetPostWithReplies(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.ForumPost
	if err := database.DB.Preload("User").Preload("Section").Preload("Replies", func (db *gorm.DB) *gorm.DB  {
		return db.Where("parent_id IS NULL").Preload("Children")
	}).First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	response := mapper.ForumPostToDTO(post)

	return c.JSON(response)
}
