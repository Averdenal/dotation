package logic

import (
	"strings"
	"time"

	"github.com/Averdenal/Dotation/Models"
)

//AddOrdinateur ajoute un ordi au client
func AddOrdinateur(c *Models.User, nameOrdinateur, codeExpress, os, nbserial, modelOrdinateur string, tarif float32) {
	valideNameOrdinateur(&nameOrdinateur)
	oridnateur := Models.Ordinateur{
		Name:        nameOrdinateur,
		CodeExpress: codeExpress,
		Os:          os,
		Materiel: Models.Materiel{
			Cat: Models.Cat{
				Name:  "Ordinateur",
				Model: modelOrdinateur,
			},
			Tarif:    tarif,
			NbSerial: nbserial,
			Dotation: Models.Dotation{
				Date: time.Now(),
			},
		},
	}

	c.Ordinateur = oridnateur
}

func valideNameOrdinateur(name *string) {
	*name = strings.TrimSpace(strings.ToUpper(*name))
}
