package models

import (
	"time"

	"gorm.io/gorm"
)


type Event struct {
	gorm.Model
	Title 		string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Category 	string `json:"category" gorm:"type:enum('kunjungan_industri', 'tamu', 'acara_sekolah', 'ujian', 'libur', 'lainnya');default:'tamu'"`
	Visibility 	string 	`json:"visibility" gorm:"type:enum('public', 'private');default:'private'"`
	StartDate 	time.Time 	`json:"start_date" gorm:"type:date"`
	EndDate 	time.Time 	`json:"end_date" gorm:"type:date"`
	Location 	string `json:"location"`
	Organizer 	string `json:"organizer"`
	Image		string `json:"image"`
}