package logic

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Averdenal/Dotation/Models"
)

//AddOrdinateur ajoute un ordi au client
func CreatedOrdinateur(nameOrdinateur, codeExpress, os, nbserial, modelOrdinateur, tarif string) (Models.Ordinateur, error) {
	s, err := strconv.ParseFloat(tarif, 32)
	if err != nil {
		return Models.Ordinateur{}, err
	}
	fmt.Println(nameOrdinateur)
	ordinateur := Models.Ordinateur{
		Name:        nameOrdinateur,
		CodeExpress: codeExpress,
		Os:          os,

		Cat: Models.Cat{
			Name:   "Ordinateur",
			Models: modelOrdinateur,
		},
		Tarif:    s,
		NbSerial: nbserial,
		Dotation: Models.Dotation{
			Date: time.Now(),
		},
	}
	return ordinateur, nil
}

func UpdateOrdinateur(o *Models.Ordinateur, nameOrdinateur, codeExpress, os, nbserial, modelOrdinateur, tarif string) error {

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

	if modelOrdinateur != "" {
		o.Cat.Models = modelOrdinateur
	}

	if tarif != "" {
		t, _ := valideTarif(&tarif)
		o.Tarif = t
	}

	return nil
}

func valideTarif(tarif *string) (float64, error) {
	t, err := strconv.ParseFloat(*tarif, 64)
	if err != nil {
		return -1, err
	}
	return t, nil
}
