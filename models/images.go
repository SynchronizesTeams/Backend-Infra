package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	FilePath 	string `json:"file_path"`
	Category 	string `json:"category" gorm:"type:enum('berita', 'mading', 'galeri', 'teacher', 'eskul', 'achievement', 'testimonial', 'saprol', 'certification', 'portal', 'mitra', 'profile');default:'galeri'"`
	Title 		string `json:"title"`
}