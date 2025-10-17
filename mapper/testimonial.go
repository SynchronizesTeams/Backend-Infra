package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func TestiToDTO(t models.Testimonial) dto.TestimonialResponse {
	return dto.TestimonialResponse{
		ID: t.ID,
		Name: t.Name,
		Position: t.Position,
		Photo: t.Photo,
		Testimonial: t.Testimonial,
		CreatedAt: t.CreatedAt,
	}
}

func TestiListToDTO(testis []models.Testimonial) []dto.TestimonialResponse {
	var result []dto.TestimonialResponse
	for _, t := range testis{
		result = append(result, TestiToDTO(t))
	}
	return result
}

func UpdateTesti(t *models.Testimonial, dto dto.TestimonialRequest) {
	if dto.Name != "" {
		t.Name = dto.Name
	}
	if dto.Position != "" {
		t.Position = dto.Position
	}
	if dto.Testimonial != "" {
		t.Testimonial = dto.Testimonial
	}

	t.UpdatedAt = time.Now()
}