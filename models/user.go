package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NoInduk 			string `json:"no_induk"`
	Name     			string `json:"name"`
	Email    			string `gorm:"unique" json:"email"`
	Password 			string `json:"-"`
	Role				string `json:"role" gorm:"type:enum('admin','guru','siswa','orang_tua');default:'siswa'"`
	PhotoUrl 			string `json:"photo_url"`
	Phone 				string `json:"phone"`
	Alamat 				string `json:"alamat" gorm:"type:text"`
	Jabatan				string `json:"jabatan"`
	TahunAjaranMulai	string `gorm:"type:year" json:"tahun_ajaran_mulai"`

	UserLinks 			[]UserLinks `gorm:"foreignKey:UserID" json:"user_links"`
	News 				[]News `gorm:"foreignKey:AuthorID" json:"news"`
}
