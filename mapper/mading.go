package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func MadingToDTO(mading models.Mading) dto.MadingResponse {
	return dto.MadingResponse{
		ID: mading.ID,
		Title: mading.Title,
		Type: mading.Type,
		Content: mading.Content,
		Image: mading.Image,
		Status: mading.Status,
		IsActive: mading.IsActive,
	}
}

func MadingListToDTO(madings []models.Mading) []dto.MadingResponse {
	var result []dto.MadingResponse
	for _, mading := range madings {
		result = append(result, MadingToDTO(mading))
	}
	return result
}

func UpdateMading(mading *models.Mading, dto dto.MadingRequest) {
	if dto.Title != "" {
		mading.Title = dto.Title
	}
	if dto.Type != "" {
		mading.Type = dto.Type
	}
	if dto.Content != "" {
		mading.Content = dto.Content
	}
	mading.UpdatedAt = time.Now()
}