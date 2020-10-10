package logic

import (
	"errors"

	"github.com/Averdenal/Dotation/Models"
)

// CreatedApplication - création des applications
func CreatedApplication(nom, version string) (Models.Application, error) {
	app := Models.Application{
		Nom:     nom,
		Version: version,
	}
	return app, nil
}

//AddApplication Ajouter des applications à l'ordinateur
func AddApplication(c *Models.User, app Models.Application) error {
	if c.Ordinateur.Name == "" {
		return errors.New("Ordinateur vide")
	}
	c.Ordinateur.Applications = append(c.Ordinateur.Applications, app)
	return nil
}
