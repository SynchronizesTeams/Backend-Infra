package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func AchievementToDTO(a models.Achievement) dto.AchievementResponse {
	return dto.AchievementResponse{
		ID: a.ID,
		Title: a.Title,
		Description: a.Description,
		Image: a.Image,
		AchievementDate: a.AchievementDate,
		News: dto.NewsAchievementResponse{
			ID: a.News.ID,
			Title: a.News.Title,
			Slug: a.News.Slug,
			Excerpt: a.News.Excerpt,
		},
	}
}


func AchievementListToDTO(as []models.Achievement) []dto.AchievementResponse {
	var result []dto.AchievementResponse
	for _, a := range as{
		result = append(result, AchievementToDTO(a))
	}
	return result
}

func UpdateAchievement(a *models.Achievement, dto dto.AchievementRequest) {
	if dto.Title != "" {
		a.Title = dto.Title
	}
	if dto.Description != "" {
		a.Description = dto.Description
	}
	if dto.AchievementDate != "" {
		a.AchievementDate = dto.AchievementDate
	}
	if dto.NewsID != 0 {
		a.NewsID = dto.NewsID
	}
	a.UpdatedAt = time.Now()
}