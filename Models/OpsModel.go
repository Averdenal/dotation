package Models

import "gorm.io/gorm"

type Ops struct {
	gorm.Model
	Name    string
	Version string
}

type Opss struct {
	Ops []Ops
}
