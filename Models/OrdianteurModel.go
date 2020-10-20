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

func (r *Ordinateur) AddApplication(a *Application) {
	r.Applications = append(r.Applications, *a)
}

func (r *Ordinateur) Upper() {
	r.Name = strings.TrimSpace(strings.ToUpper(r.Name))
	r.Os = strings.TrimSpace(strings.ToUpper(r.Os))
}
