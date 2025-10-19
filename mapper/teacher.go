package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func TeacherToDTO(t models.Teacher) dto.TeacherResponse {
	return dto.TeacherResponse{
		ID: t.ID,
		NIG: t.NIG,
		FullName: t.FullName,
		Position: t.Position,
		Subject: t.Subject,
		Description: t.Description,
		Photo: t.Photo,
		User: dto.UserTeacherResponse{
			ID: t.User.ID,
			Name: t.User.Name,
		},
	}
}

func TeacherListToDTO(teachers []models.Teacher) []dto.TeacherResponse {
	var result []dto.TeacherResponse
	for _, t := range teachers {
		result = append(result, TeacherToDTO(t))
	}
	return result
}


func UpdateTeacherFromDTO(t *models.Teacher, dto dto.TeacherRequest) {
	if dto.NIG != "" {
		t.NIG = dto.NIG
	}
	if dto.FullName != "" {
		t.FullName = dto.FullName
	}
	if dto.Description != "" {
		t.Description = dto.Description
	}
	if dto.Position != "" {
		t.Position = dto.Position
	}
	if dto.Subject != "" {
		t.Subject = dto.Subject
	}

	t.UpdatedAt = time.Now()
}