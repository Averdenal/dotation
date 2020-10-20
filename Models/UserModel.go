package Models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nom        string
	Prenom     string
	Service    string
	Ordinateur []Ordinateur `gorm:"foreignKey:UserRefer"`
	Materiel   []Materiel   `gorm:"foreignKey:UserRefer"`
	pwd        string
}

func (r *User) AddOrdinateur(o *Ordinateur) {
	r.Ordinateur = append(r.Ordinateur, *o)
}

func (r *User) AddMateriel(m *Materiel) {
	r.Materiel = append(r.Materiel, *m)
}
