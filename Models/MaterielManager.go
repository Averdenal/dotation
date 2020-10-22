package Models

func (m *Materiel) FindById(id string) error {
	database.First(m, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func FindAllCat() (c []Car, err error) {
	result := database.Find(&m)
	err = result.Error
	return c, err
}
func (m *Materiel) FindByName(n string) error {
	database.Where("name = ?", n).First(&m)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Materiel) Saver() error {
	database.Save(m)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Materiel) Delete() error {
	database.Delete(&m)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
