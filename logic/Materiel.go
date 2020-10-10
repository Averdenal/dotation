package logic

import (
	"time"

	"github.com/Averdenal/Dotation/Models"
)

// CreatedMateriel Création de Mat
func CreatedMateriel(tarif float32, nbSerial, nom, model string) (Models.Materiel, error) {
	materiel := Models.Materiel{
		Cat: Models.Cat{
			Model: model,
			Name:  nom,
		},
		Tarif:    tarif,
		NbSerial: nbSerial,
		Dotation: Models.Dotation{
			Date: time.Now(),
		},
	}
	return materiel, nil
}

//AddMateriel Add Materiel au Client
func AddMateriel(c *Models.User, m Models.Materiel) {
	c.Materiel = append(c.Materiel, m)
}
