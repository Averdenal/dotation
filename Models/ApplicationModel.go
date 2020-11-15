package Models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	OrdinateurRefer uint
	Ordinateur      []*Ordinateur `gorm:"many2many:application_ordinateur;"`
	Nom             string
	Version         string
}

type Applications struct {
	apps []Application
}
