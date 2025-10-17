package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func PortalToDTO(p models.Portal) dto.PortalResponse {
	return dto.PortalResponse{
		ID:          p.ID,
		Name:        p.Name,
		URL:         p.URL,
		Logo: 		 p.Logo,
		Category: 	 p.Category,
		IsSSOEnabled: p.IsSSOEnabled,
		IsActive: p.IsActive,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
	}
}

func PortalListToDTO(portals []models.Portal) []dto.PortalResponse {
	var result []dto.PortalResponse
	for _, p := range portals {
		result = append(result, PortalToDTO(p))
	}
	return result
}

func UpdatePortalFromDTO(p *models.Portal, dto dto.PortalRequest) {
	if dto.Name != "" {
		p.Name = dto.Name
	}
	if dto.URL != "" {
		p.URL = dto.URL
	}
	if dto.Logo != "" {
		p.Logo = dto.Logo
	}
	if dto.Category != "" {
		p.Category = dto.Category
	}
	if dto.Description != "" {
		p.Description = dto.Description
	}

	p.UpdatedAt = time.Now()
}