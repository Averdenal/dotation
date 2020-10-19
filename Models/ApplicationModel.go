package Models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	OrdinateurRefer uint
	Nom             string
	Version         string
}
