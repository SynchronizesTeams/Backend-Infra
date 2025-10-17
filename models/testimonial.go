package models

import "gorm.io/gorm"

type Testimonial struct {
	gorm.Model
	Name string
	Position string
	Photo  string
	Testimonial string
}