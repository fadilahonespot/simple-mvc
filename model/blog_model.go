package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	UserId uint   `json:"userId"`
	Judul  string `json:"judul"`
	Konten string `json:"konten"`
}
