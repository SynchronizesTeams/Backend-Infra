package dto

import "time"

type PortalRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	URL         string `form:"url"`
	Logo        string `form:"logo"`
	Category    string `form:"category"`
}

type PortalResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	URL          string    `json:"url"`
	Logo         string    `json:"logo"`
	Category     string    `json:"category"`
	IsSSOEnabled bool      `json:"is_sso_enabled"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}