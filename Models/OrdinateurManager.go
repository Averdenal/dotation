package Models

func (o *Ordinateur) FindById(id string) error {
	return database.First(o, id).Error
}
func (o *Ordinateurs) FindAllOrdinateur() error {
	return database.Find(&o.Ordinateurs).Error
}
func (o *Ordinateur) FindByName(n string) error {
	return database.Where("name = ?", n).First(&o).Error
}

func (o *Ordinateur) Saver() error {
	o.Upper()
	return database.Save(o).Error
}

func (o *Ordinateur) Delete() error {
	return database.Delete(&o).Error
}
