package logic

import (
	"github.com/Averdenal/Dotation/Models"
	"strconv"
)

func UpdateOrdinateur(o *Models.Ordinateur, creds Models.CredsOrdinateur) error {

	if creds.NameOrdinateur != "" {
		o.Name = creds.NameOrdinateur
	}

	if creds.Codeexpress != "" {
		o.CodeExpress = creds.Codeexpress
	}
	if creds.Nbserial != "" {
		o.NbSerial = creds.Nbserial
	}

	if creds.Tarif != "" {
		o.Tarif, _ = ValideTarif(creds.Tarif)
	}

	return nil
}

func ValideTarif(tarif string) (float64, error) {
	t, err := strconv.ParseFloat(tarif, 64)
	if err != nil {
		return -1, err
	}
	return t, nil
}
