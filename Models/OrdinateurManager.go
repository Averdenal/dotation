package Models

import "fmt"

func (o *Ordinateur) FindById(id string) error {
	database.First(o, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func FindAllOordinateur() (o []Ordinateur, err error) {
	result := database.Find(&o)
	fmt.Println(result)
	err = result.Error
	return o, err
}
func (o *Ordinateur) FindByName(n string) error {
	database.Where("name = ?", n).First(&o)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Ordinateur) Saver() error {
	o.Upper()
	database.Save(o)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Ordinateur) Delete() error {
	database.Delete(&o)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
