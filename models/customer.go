package models

type Customer struct {
	ID           int     `gorm:"primaryKey" json:"id"`
	NIK          string  `gorm:"not null" json:"nik"`
	FullName     string  `gorm:"not null" json:"full_name"`
	LegalName    string  `gorm:"not null" json:"legal_name"`
	PlaceOfBirth string  `gorm:"not null" json:"place_of_birth"`
	BirthDate    string  `gorm:"not null" json:"birth_date"`
	Salary       float64 `gorm:"not null" json:"salary"`
	PhotoKTP     string  `gorm:"not null" json:"photo_ktp"`
	PhotoSelfie  string  `gorm:"not null" json:"photo_selfie"`
}
