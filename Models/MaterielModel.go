package Models

import "gorm.io/gorm"

type Materiel struct {
	gorm.Model
	UserRefer uint
	CatRefer  int
	Cat       Cat `gorm:"foreignKey:CatRefer"`
	Tarif     float32
	NbSerial  string
	Dotation
}

func (r *Materiel) AddCat(c *Cat) {
	r.Cat = *c
}
