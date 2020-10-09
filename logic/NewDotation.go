package logic

import (
	"strings"

	"github.com/Averdenal/Dotation/Models"
)

//NewUser - Créer un nouveau client
func NewUser(prenom, nom, service string) (Models.Client, error) {

	client := Models.Client{
		Prenom:  prenom,
		Nom:     nom,
		Service: service,
	}
	valideService(&client)
	return client, nil
}

func valideService(c *Models.Client) {
	if !strings.HasSuffix(c.Service, "Service ") {
		c.Service = "Service " + c.Service
	}
}
