package logic

import (
	"strings"

	"github.com/Averdenal/Dotation/Models"
)

//NewUser - Créer un nouveau client
func NewUser(prenom, nom, service string) (Models.User, error) {

	client := Models.User{
		Prenom:  prenom,
		Nom:     nom,
		Service: service,
	}
	valideService(&client)
	return client, nil
}

func valideService(c *Models.User) {
	if !strings.HasSuffix(c.Service, "Service ") {
		c.Service = "Service " + c.Service
	}
}
