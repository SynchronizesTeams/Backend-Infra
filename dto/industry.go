package dto

import "time"

type IndustryRequest struct {
	Name        string `form:"name"`
	Logo        string `form:"logo"`
	Website     string `form:"website"`
	Description string `form:"description"`
}

type IndustryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Website     string    `json:"website"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}