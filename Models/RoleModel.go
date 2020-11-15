package Models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
	Code int
	User []*User `gorm:"many2many:user_role;"`
}
