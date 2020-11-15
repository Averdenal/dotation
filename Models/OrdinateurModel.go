package Models

import (
	"database/sql"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Ordinateur struct {
	gorm.Model
	Name        string
	CodeExpress string
	OpsRefer    int
	Ops         Ops `gorm:"foreignKey:OpsRefer"`
	CatRefer    int
	Cat         Cat `gorm:"foreignKey:CatRefer"`
	UserRefer   uint
	Tarif       float64
	NbSerial    string
	Dotation
	Applications []*Application `gorm:"many2many:application_ordinateur;"`
}

type Ordinateurs struct {
	Ordinateurs []Ordinateur
}

type CredsOrdinateur struct {
	NameOrdinateur string
	Codeexpress    string
	Idos           string
	Nbserial       string
	Idcat          string
	Tarif          string
}

func (o *Ordinateur) Created(ordinateur CredsOrdinateur) error {
	var cat Cat
	var ops Ops

	_ = cat.FindById(ordinateur.Idcat)
	_ = ops.FindById(ordinateur.Idos)

	s, err := strconv.ParseFloat(ordinateur.Tarif, 32)
	if err != nil {
		return err
	}
	o = &Ordinateur{
		Name:        ordinateur.NameOrdinateur,
		CodeExpress: ordinateur.Codeexpress,
		Ops:         ops,
		Cat:         cat,
		Tarif:       s,
		NbSerial:    ordinateur.Nbserial,
		Dotation: Dotation{
			Date: sql.NullTime{Valid: false},
		},
	}
	return nil
}
func (o *Ordinateur) AddApplication(a *Application) {
	o.Applications = append(o.Applications, a)
}
func (o *Ordinateur) Upper() {
	o.Name = strings.TrimSpace(strings.ToUpper(o.Name))
}
