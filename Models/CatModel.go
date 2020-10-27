package Models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Type   string
	Models string
}
type Cats struct {
	Cats []Cat
}
