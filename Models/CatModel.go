package Models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name   string
	Models string
}
