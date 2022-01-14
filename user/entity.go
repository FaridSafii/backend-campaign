package user

import "time"

//Membuat koneksi ke tabel Users (plural) kedalam struct User (tunggal)
//struct bagaikan model dalam bahasa pemrograman lain
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
