package Models

import (
	"gorm.io/gorm"
	"strings"
)

type Ordinateur struct {
	gorm.Model
	Name        string
	CodeExpress string
	Os          string
	CatRefer    int
	Cat         Cat `gorm:"foreignKey:CatRefer"`
	UserRefer   uint
	Tarif       float64
	NbSerial    string
	Dotation
	Applications []Application `gorm:"foreignKey:OrdinateurRefer"`
}

func (o *Ordinateur) AddApplication(a *Application) {
	o.Applications = append(o.Applications, *a)
}

func (o *Ordinateur) Upper() {
	o.Name = strings.TrimSpace(strings.ToUpper(o.Name))
	o.Os = strings.TrimSpace(strings.ToUpper(o.Os))
}
