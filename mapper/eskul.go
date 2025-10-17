package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func EskulToDTO(e models.Eskul) dto.EskulResponse{
	return dto.EskulResponse{
		ID: e.ID,
		Name: e.Name,
		Description: e.Description,
		Image: e.Image,
		PembinaID: e.PembinaID,
		Pembina: dto.TeacherEskulResponse{
			ID: e.Pembina.ID,
			FullName: e.Pembina.FullName,
		},
	}
}


func EskulListToDTO(eskuls []models.Eskul) []dto.EskulResponse {
	var result []dto.EskulResponse
	for _, e := range eskuls {
		result = append(result, EskulToDTO(e))
	}
	return result
}

func UpdateEskulFromDTO(e *models.Eskul, dto dto.EskulRequest) {
	if dto.Name != "" {
		e.Name = dto.Name
	}
	if dto.Description != "" {
		e.Description = dto.Description
	}
	if dto.PembinaID != 0 {
		e.PembinaID = dto.PembinaID
	}

	e.UpdatedAt = time.Now()
}