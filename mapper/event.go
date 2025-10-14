package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
)

func EventToDTO(event models.Event) dto.EventResponse {
	return dto.EventResponse {
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Category:    event.Category,
		Visibility:  event.Visibility,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Location:    event.Location,
		Organizer:   event.Organizer,
		Image:       event.Image,
		CreatedAt:   event.CreatedAt,
	}
}

func EventListToDTO(events []models.Event) []dto.EventResponse {
	var result []dto.EventResponse
	for _, e := range events {
		result = append(result, EventToDTO(e))
	}
	return result
}
