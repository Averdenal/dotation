package logic

import (
	"fmt"
	"github.com/Averdenal/Dotation/db"
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
	ordinateur.ValideNameOrdinateur()
	ordinateur.Save()
	return ordinateur, nil
}

func UpdateOrdinateur(id, nameOrdinateur, codeExpress, os, nbserial, modelOrdinateur, tarif string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	db := db.DbConnnect()
	o := Models.Ordinateur{}
	db.First(&o).Where("id = ?", i)

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
	fmt.Println(o.Name)
	fmt.Println(nameOrdinateur)
	//db.Save(&o)
	return nil
}

func DeleteOrdianteur(id string) error {
	db := db.DbConnnect()
	var o Models.Ordinateur
	db.Where("id = ?", id).Delete(&o)
	return nil
}

func valideTarif(tarif *string) (float64, error) {
	t, err := strconv.ParseFloat(*tarif, 64)
	if err != nil {
		return -1, err
	}
	return t, nil
}
