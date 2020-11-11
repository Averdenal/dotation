package Models

func (r *Materiel) FindById(id string) error {
	database.First(r, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func (m *Materiels) FindAllMateriel() error {
	result := database.Find(&m.Materiels)
	return result.Error
}
func (r *Materiel) FindByName(n string) error {
	database.Where("name = ?", n).First(&r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Materiel) Saver() error {
	database.Save(r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Materiel) Delete() error {
	database.Delete(&r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
