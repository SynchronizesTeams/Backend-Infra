package models

import "gorm.io/gorm"

type Certification struct {
	gorm.Model
	Name 					string
	Issuer 					string
	Description 			string
	Image 					string
	CertificationNumber  	string
	IssueDate 				string
	ExpiryDate 				string
}