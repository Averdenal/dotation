package Models

import (
	"time"
)

type User struct {
	Nom        string
	Prenom     string
	Service    string
	Ordinateur Ordinateur
	Materiel   []Materiel
	pwd        string
}

type Dotation struct {
	Date time.Time
}

type Materiel struct {
	Cat      Cat
	Tarif    float32
	NbSerial string
	Dotation Dotation
}

type Application struct {
	Nom     string
	Version string
}

type Cat struct {
	Name  string
	Model string
}

type Ordinateur struct {
	Name         string
	CodeExpress  string
	Os           string
	Materiel     Materiel
	Applications []Application
}
