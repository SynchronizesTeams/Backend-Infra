package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func UserToDTO(u models.User) dto.UserResponse {
	return dto.UserResponse {
		ID:               u.ID,
		NoInduk:          u.NoInduk,
		Name:             u.Name,
		Email:            u.Email,
		Role:             u.Role,
		PhotoUrl:         u.PhotoUrl,
		Phone:            u.Phone,
		Alamat:           u.Alamat,
		Jabatan:          u.Jabatan,
		TahunAjaranMulai: u.TahunAjaranMulai,
	}
}

func UpdateUser(u *models.User, dto dto.UserRequest) {
	if dto.NoInduk != "" {
		u.NoInduk = dto.NoInduk
	}
	if dto.Name != "" {
		u.Name = dto.Name
	}
	if dto.Email != "" {
		u.Email = dto.Email
	}
	if dto.Role != "" {
		u.Role = dto.Role
	}
	if dto.Phone != "" {
		u.Phone = dto.Phone
	}
	if dto.Alamat != "" {
		u.Alamat = dto.Alamat
	}
	if dto.Jabatan != "" {
		u.Jabatan = dto.Jabatan
	}
	if dto.TahunAjaranMulai != "" {
		u.TahunAjaranMulai = dto.TahunAjaranMulai
	}
	if dto.Phone != "" {
		u.Phone = dto.Phone
	}
	u.UpdatedAt = time.Now()
}