package Models

func (o *Ordinateur) FindById(id string) error {
	database.First(o, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func (o *Ordinateurs) FindAllOrdinateur() error {
	result := database.Find(&o.Ordinateurs)
	return result.Error
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
