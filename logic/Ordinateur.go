package logic

import (
	"database/sql"
	"fmt"
	"github.com/Averdenal/Dotation/Models"
	"strconv"
)

//AddOrdinateur ajoute un ordi au client
func CreatedOrdinateur(nameOrdinateur, codeExpress, os, nbserial, tarif string, cat Models.Cat) (Models.Ordinateur, error) {
	s, err := strconv.ParseFloat(tarif, 32)
	if err != nil {
		return Models.Ordinateur{}, err
	}
	fmt.Println(nameOrdinateur)
	ordinateur := Models.Ordinateur{
		Name:        nameOrdinateur,
		CodeExpress: codeExpress,
		Os:          os,
		Cat:         cat,
		Tarif:       s,
		NbSerial:    nbserial,
		Dotation: Models.Dotation{
			Date: sql.NullTime{Valid: false},
		},
	}
	return ordinateur, nil
}

func UpdateOrdinateur(o *Models.Ordinateur, nameOrdinateur, codeExpress, os, nbserial, tarif string) error {

	if nameOrdinateur != "" {
		o.Name = nameOrdinateur
	}

	if codeExpress != "" {
		o.CodeExpress = codeExpress
	}

	if os != "" {
		o.Os = os
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
