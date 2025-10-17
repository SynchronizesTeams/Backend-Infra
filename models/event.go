package models

import (
	"gorm.io/gorm"
)


type Event struct {
	gorm.Model
	Title 		string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Category 	string `json:"category" gorm:"type:enum('kunjungan_industri', 'tamu', 'acara_sekolah', 'ujian', 'libur', 'lainnya');default:'tamu'"`
	Visibility 	string 	`json:"visibility" gorm:"type:enum('public', 'private');default:'private'"`
	StartDate 	string 	`json:"start_date"`
	EndDate 	string 	`json:"end_date"`
	Location 	string `json:"location"`
	Organizer 	string `json:"organizer"`
	Image		string `json:"image"`
}