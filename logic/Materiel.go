package logic

import (
	"github.com/Averdenal/Dotation/Models"
)

// CreatedMateriel Création de Mat
func CreatedMateriel(tarif, nbSerial string, cat Models.Cat) (Models.Materiel, error) {
	t, err := ValideTarif(tarif)
	if err != nil {

	}
	materiel := Models.Materiel{
		Cat:      cat,
		Tarif:    t,
		NbSerial: nbSerial,
	}
	return materiel, nil
}

func UpdateMateriel(m *Models.Materiel, tarif, nbSerial string, cat Models.Cat) {

}
