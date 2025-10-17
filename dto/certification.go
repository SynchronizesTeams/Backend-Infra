package dto

import "time"

type CertificationRequest struct {
	Name                string `form:"name"`
	Issuer              string `form:"issuer"`
	Description         string `form:"description"`
	Image               string `form:"image"`
	CertificationNumber string `form:"certification_number"`
	IssueDate           string `form:"issue_date"`
	ExpiryDate          string `form:"expiry_date"`
}

type CertificationResponse struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	Issuer              string `json:"issuer"`
	Description         string `json:"description"`
	Image               string `json:"image"`
	CertificationNumber string `json:"certification_number"`
	IssueDate           string `json:"issue_date"`
	ExpiryDate          string `json:"expiry_date"`
	CreatedAt           time.Time `json:"created_at"`
}