package Models

func (a *Application) FindById(id string) error {
	database.First(a, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func (a *Applications) FindAllApplication() error {
	result := database.Find(&a.apps)
	return result.Error
}
func (a *Application) FindByName(n string) error {
	database.Where("name = ?", n).First(&a)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Application) Saver() error {
	database.Save(a)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Application) Delete() error {
	database.Delete(&a)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func (a *Application) Update(nom, version string) {
	if nom != "" {
		a.Nom = nom
	}
	if version != "" {
		a.Version = version
	}
}
