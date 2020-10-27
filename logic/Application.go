package logic

import (
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

func UpdateApplication(nom, version string, app Models.Application) Models.Application {
	if nom != "" {
		app.Nom = nom
	}
	if version != "" {
		app.Version = version
	}

	return app
}
