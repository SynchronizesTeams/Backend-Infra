package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func IndustryToDTO(i models.Industry) dto.IndustryResponse {
	return dto.IndustryResponse{
		ID: i.ID,
		Name: i.Name,
		Logo: i.Logo,
		Website: i.Website,
		Description: i.Description,
		CreatedAt: i.CreatedAt,
	}
}

func IndustryListToDTO(industries []models.Industry) []dto.IndustryResponse {
	var result []dto.IndustryResponse
	for _, i := range industries{
		result = append(result, IndustryToDTO(i))
	}
	return result
}

func UpdateIndustry(i *models.Industry, dto dto.IndustryRequest) {
	if dto.Name != "" {
		i.Name = dto.Name
	}
	if dto.Website != "" {
		i.Website = dto.Website
	}
	if dto.Description != "" {
		i.Description = dto.Description
	}
	i.UpdatedAt = time.Now()
}