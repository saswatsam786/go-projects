package models

import (
	"github.com/jinzhu/gorm"
	"github.com/saswatsam786/go-projects/go-mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config
}