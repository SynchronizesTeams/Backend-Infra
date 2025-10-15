package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
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

func UpdateEventFromDTO(event *models.Event, dto dto.EventRequest) {
	if dto.Title != "" {
		event.Title = dto.Title
	}
	if dto.Description != "" {
		event.Description = dto.Description
	}
	if dto.Category != "" {
		event.Category = dto.Category
	}
	if dto.Location != ""  {
		event.Location = dto.Location
	}
	if dto.Organizer != "" {
		event.Organizer = dto.Organizer
	}
	if !dto.StartDate.IsZero() {
		event.StartDate = dto.StartDate
	}
	if !dto.EndDate.IsZero() {
		event.EndDate = dto.EndDate
	}
	event.UpdatedAt = time.Now()
}
