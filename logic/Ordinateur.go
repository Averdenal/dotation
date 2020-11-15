package logic

import (
	"github.com/Averdenal/Dotation/Models"
	"strconv"
)

func UpdateOrdinateur(o *Models.Ordinateur, nameOrdinateur, codeExpress, nbserial, tarif string) error {

	if nameOrdinateur != "" {
		o.Name = nameOrdinateur
	}

	if codeExpress != "" {
		o.CodeExpress = codeExpress
	}
	if nbserial != "" {
		o.NbSerial = nbserial
	}

	if tarif != "" {
		t, _ := ValideTarif(&tarif)
		o.Tarif = t
	}

	return nil
}

func ValideTarif(tarif *string) (float64, error) {
	t, err := strconv.ParseFloat(*tarif, 64)
	if err != nil {
		return -1, err
	}
	return t, nil
}
