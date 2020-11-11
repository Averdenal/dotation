package Models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name string
	Code string
}

type Services struct {
	Services []Service
}
