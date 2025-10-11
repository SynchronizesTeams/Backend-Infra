package handlers

import (
	"fmt"
	"go-api-infra/database"
	"go-api-infra/models"
	// "strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "gorm.io/gorm"
	"os"
)

func slugify(s string) string {
	// Ubah huruf besar → kecil
	s = strings.ToLower(s)
	// Ganti spasi & karakter aneh jadi tanda "-"
	s = strings.ReplaceAll(s, " ", "-")
	// Hapus karakter non-alfanumerik selain "-"
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == '-' {
			return r
		}
		return -1
	}, s)
	return s
}

func CreateNews(c *fiber.Ctx) error {
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
	authorID := uint(idVal)

	title := c.FormValue("title")
	content := c.FormValue("content")
	excerpt := c.FormValue("excerpt")
	status := c.FormValue("status")
	rawTags := c.FormValue("tags")
	slug := slugify(title)

	file, err := c.FormFile("thumbnail")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Thumbnail file required"})
	}

	os.MkdirAll("uploads/news", os.ModePerm)

	savePath := fmt.Sprintf("uploads/news/%d_%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save thumbnail"})
	}

	news := models.News{
		Title:     title,
		Slug:      slug,
		Content:   content,
		Excerpt:   excerpt,
		Status:    status,
		Thumbnail: savePath,
		AuthorID:  authorID,
	}

	var tags []models.NewsTag
	if rawTags != "" {
		tagNames := strings.Split(rawTags, ",")
		for _, tagName := range tagNames {
			tagName = strings.TrimSpace(tagName)
			if tagName == "" {
				continue
			}

			var tag models.NewsTag
			if err := database.DB.Where("name = ?", tagName).First(&tag).Error; err != nil {
				tag = models.NewsTag{Name: tagName}
				database.DB.Create(&tag)
			}
			tags = append(tags, tag)
		}
		news.Tags = tags
	}

	if err := database.DB.Create(&news).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create news"})
	}

	database.DB.Preload("Author").First(&news, news.ID)

	return c.Status(201).JSON(fiber.Map{
		"id":        news.ID,
		"title":     news.Title,
		"slug":      news.Slug,
		"excerpt":   news.Excerpt,
		"thumbnail": news.Thumbnail,
		"status":    news.Status,
		"tags":      tags,
		"author": fiber.Map{
			"id":   news.AuthorID,
			"name": news.Author.Name,
		},
	})
}


func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil data news lama
	var news models.News
	if err := database.DB.Preload("Tags").First(&news, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "News not found",
		})
	}

	// Ambil data input langsung dari form
	title := c.FormValue("title")
	content := c.FormValue("content")
	excerpt := c.FormValue("excerpt")
	status := c.FormValue("status")
	tagsInput := c.FormValue("tags") // bisa jadi: "golang,tech,update"

	// Update kolom text
	if title != "" {
		news.Title = title
		news.Slug = slugify(title)
	}
	if content != "" {
		news.Content = content
	}
	if excerpt != "" {
		news.Excerpt = excerpt
	}
	if status != "" {
		news.Status = status
	}

	// Handle thumbnail (opsional)
	file, err := c.FormFile("thumbnail")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		path := fmt.Sprintf("uploads/news/%s", filename)

		// Simpan file baru
		if err := c.SaveFile(file, path); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to upload thumbnail",
			})
		}

		// Hapus file lama (kalau ada)
		if news.Thumbnail != "" {
			_ = os.Remove(news.Thumbnail)
		}

		news.Thumbnail = path
	}

	// Handle tag (bisa kosong, satu, atau banyak)
	if tagsInput != "" {
		tagsArr := strings.Split(tagsInput, ",") // misal "golang,tech,update"

		var tags []models.NewsTag
		for _, tagName := range tagsArr {
			tagName = strings.TrimSpace(tagName)
			if tagName == "" {
				continue
			}

			var tag models.NewsTag
			if err := database.DB.Where("name = ?", tagName).First(&tag).Error; err != nil {
				tag = models.NewsTag{Name: tagName}
				database.DB.Create(&tag)
			}
			tags = append(tags, tag)
		}

		// Replace relasi lama dengan yang baru
		if err := database.DB.Model(&news).Association("Tags").Replace(tags); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update tags",
			})
		}
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update news",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "News updated successfully",
		"data":    news,
	})
}

func GetNews(c *fiber.Ctx) error {
	id := c.Params("id")
	var news []models.News
	if err := database.DB.Preload("Author").Preload("Tags").First(&news, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch news",
		})
	}

	// Bentuk response ringkas
	var response []fiber.Map
	for _, n := range news {
		var tagNames []string
		for _, t := range n.Tags {
			tagNames = append(tagNames, t.Name)
		}

		response = append(response, fiber.Map{
			"id":        n.ID,
			"title":     n.Title,
			"slug":      n.Slug,
			"excerpt":   n.Excerpt,
			"thumbnail": n.Thumbnail,
			"status":    n.Status,
			"viewCount": n.ViewCount,
			"tags":      tagNames,
			"author": fiber.Map{
				"id":   n.Author.ID,
				"name": n.Author.Name,
			},
			"created_at": n.CreatedAt,
		})
	}

	return c.JSON(response)
}

func DeleteNews(c *fiber.Ctx) error {
	id := c.Params("id")

	var news models.News

	// Cari news + preload tags agar bisa hapus pivot-nya
	if err := database.DB.Preload("Tags").First(&news, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "News not found",
		})
	}

	// --- Hapus file thumbnail jika ada ---
	if news.Thumbnail != "" {
		if err := os.Remove(news.Thumbnail); err != nil && !os.IsNotExist(err) {
			// abaikan error kalau file sudah tidak ada, tapi laporkan error lain
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to delete thumbnail: %v", err),
			})
		}
	}

	// --- Hapus relasi tag di tabel pivot ---
	if err := database.DB.Model(&news).Association("Tags").Clear(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to clear news tags",
		})
	}

	// --- Hapus record berita (soft delete) ---
	if err := database.DB.Delete(&news).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete news",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "News successfully deleted",
	})
}


func GetAllNews(c *fiber.Ctx) error {
	var newsList []models.News
	if err := database.DB.Preload("Author").Preload("Tags").Find(&newsList).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch news",
		})
	}

	// Bentuk response ringkas
	var response []fiber.Map
	for _, n := range newsList {
		var tagNames []string
		for _, t := range n.Tags {
			tagNames = append(tagNames, t.Name)
		}

		response = append(response, fiber.Map{
			"id":        n.ID,
			"title":     n.Title,
			"slug":      n.Slug,
			"excerpt":   n.Excerpt,
			"thumbnail": n.Thumbnail,
			"status":    n.Status,
			"viewCount": n.ViewCount,
			"tags":      tagNames,
			"author": fiber.Map{
				"id":   n.Author.ID,
				"name": n.Author.Name,
			},
			"created_at": n.CreatedAt,
		})
	}

	return c.JSON(response)
}

func AddView (c *fiber.Ctx) error {
	id := c.Params("id")
	var news models.News
	
	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "News not found",
		})
	}

	news.ViewCount += 1

	if err := database.DB.Save(&news).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update view count",
		})
	}

	return c.JSON(fiber.Map{
		"view_count": news.ViewCount,
	})
}

func ChangeStatus (c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.FormValue("status")
	var news models.News

	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cannot find news",
		})
	}

	news.Status = status

	if err := database.DB.Model(&news).Update("status", status).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to change status"})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": news.Status,
	})
}