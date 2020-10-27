package logic

import "github.com/Averdenal/Dotation/Models"

func CreatedCat(typeCat, model string) (Models.Cat, error) {
	cat := Models.Cat{
		Type:   typeCat,
		Models: model,
	}
	return cat, nil
}
