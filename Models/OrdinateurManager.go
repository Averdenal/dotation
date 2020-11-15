package Models

func (o *Ordinateur) FindById(id string) error {
	err := database.First(o, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (o *Ordinateurs) FindAllOrdinateur() error {
	err := database.Find(&o.Ordinateurs).Error
	return err
}
func (o *Ordinateur) FindByName(n string) error {
	err := database.Where("name = ?", n).First(&o).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Ordinateur) Saver() error {
	o.Upper()
	err := database.Save(o).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Ordinateur) Delete() error {
	err := database.Delete(&o).Error
	if err != nil {
		return err
	}
	return nil
}
