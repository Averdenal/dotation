package Models

import (
	"github.com/Averdenal/Dotation/db"
	"gorm.io/gorm"
	"strings"
)

type Ordinateur struct {
	gorm.Model
	Name          string
	CodeExpress   string
	Os            string
	CatRefer      int
	Cat           Cat `gorm:"foreignKey:CatRefer"`
	UserRefer     uint
	Tarif         float64
	NbSerial      string
	DotationRefer int
	Dotation      Dotation      `gorm:"foreignKey:DotationRefer"`
	Applications  []Application `gorm:"foreignKey:OrdinateurRefer"`
}

func (r *Ordinateur) Save() {
	db := db.DbConnnect()
	db.Save(*r)
}

func (r *Ordinateur) AddApplication(a *Application) {
	r.Applications = append(r.Applications, *a)
}

func (r *Ordinateur) ValideNameOrdinateur() {
	r.Name = strings.TrimSpace(strings.ToUpper(r.Name))
}
