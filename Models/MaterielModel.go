package Models

import "gorm.io/gorm"

type Materiel struct {
	gorm.Model
	UserRefer uint
	CatRefer  int
	Cat       Cat `gorm:"foreignKey:CatRefer"`
	Tarif     float64
	NbSerial  string
	Dotation
}

type Materiels struct {
	Materiels []Materiel
}

func (r *Materiel) AddCat(c *Cat) {
	r.Cat = *c
}
