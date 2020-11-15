package Models

import "github.com/Averdenal/Dotation/db"

var database = db.DbConnnect()

type ModelInterface interface {
	Saver() error
	Delete() error
	FindById(id string) error
	FindByName(name string) error
}
