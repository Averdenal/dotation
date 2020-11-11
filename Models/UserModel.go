package Models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Nom          string
	Prenom       string
	ServiceRefer int
	Service      Service      `gorm:"foreignKey:ServiceRefer"`
	Ordinateur   []Ordinateur `gorm:"foreignKey:UserRefer"`
	Materiel     []Materiel   `gorm:"foreignKey:UserRefer"`
	Pwd          string
}
type Claims struct {
	IdUser   uint   `json:"id_user"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Users struct {
	Users []User
}

func (r *User) AddOrdinateur(o *Ordinateur) {
	r.Ordinateur = append(r.Ordinateur, *o)
}

func (r *User) AddMateriel(m *Materiel) {
	r.Materiel = append(r.Materiel, *m)
}
