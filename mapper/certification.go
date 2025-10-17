package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func CertToDTO(c models.Certification) dto.CertificationResponse{
	return dto.CertificationResponse{
		ID: c.ID,
		Name: c.Name,
		Issuer: c.Issuer,
		Description: c.Description,
		Image: c.Image,
		CertificationNumber: c.CertificationNumber,
		IssueDate: c.IssueDate,
		ExpiryDate: c.ExpiryDate,
		CreatedAt: c.CreatedAt,
	}
}

func CertListToDTO(certificates []models.Certification) []dto.CertificationResponse {
	var result []dto.CertificationResponse
	for _, c := range certificates{
		result = append(result, CertToDTO(c))
	}
	return result
}

func UpdateCert(c *models.Certification, dto dto.CertificationRequest) {
	if dto.Name != "" {
		c.Name = dto.Name
	}
	if dto.Description != "" {
		c.Description = dto.Description
	}
	if dto.Issuer != "" {
		c.Issuer = dto.Issuer
	}
	if dto.CertificationNumber != "" {
		c.CertificationNumber = dto.CertificationNumber
	}
	if dto.IssueDate != "" {
		c.IssueDate = dto.IssueDate
	}
	if dto.ExpiryDate != "" {
		c.ExpiryDate = dto.ExpiryDate
	}
	c.UpdatedAt = time.Now()
}