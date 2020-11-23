package Models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Cat    string
	Type   string
	Models string
}
type Cats struct {
	Cats []Cat
}
