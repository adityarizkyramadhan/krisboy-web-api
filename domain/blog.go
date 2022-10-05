package domain

import "gorm.io/gorm"

//Admin
type Admin struct {
	Username    string
	Password    string
	Name        string
	Address     string
	Description string
}

//Entitas perblog
type Blog struct {
	gorm.Model
	Photos Photo `gorm:"foreignkey:BlogId"`
}

//Entitas Gambar
type Photo struct {
	gorm.Model
	BlogId     uint
	LinkPhotos uint
}
