package helpers

import "strings"

func Slugify(s string) string {
	// Ubah huruf besar → kecil
	s = strings.ToLower(s)
	// Ganti spasi & karakter aneh jadi tanda "-"
	s = strings.ReplaceAll(s, " ", "-")
	// Hapus karakter non-alfanumerik selain "-"
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == '-' {
			return r
		}
		return -1
	}, s)
	return s
}
