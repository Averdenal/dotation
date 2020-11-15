package Models

func (r *Materiel) FindById(id string) error {
	return database.First(r, id).Error
}
func (m *Materiels) FindAllMateriel() error {
	return database.Find(&m.Materiels).Error
}
func (r *Materiel) FindByName(n string) error {
	return database.Where("name = ?", n).First(&r).Error
}

func (r *Materiel) Saver() error {
	return database.Save(r).Error
}

func (r *Materiel) Delete() error {
	return database.Delete(&r).Error
}
