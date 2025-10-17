package dto

import "time"

type TestimonialRequest struct {
	Name        string `form:"name"`
	Position    string `form:"position"`
	Testimonial string `form:"testimonial"`
}

type TestimonialResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Photo       string    `json:"photo"`
	Testimonial string    `json:"testimonial"`
	CreatedAt   time.Time `json:"created_at"`
}