package Models

func (a *Application) FindById(id string) error {
	return database.First(a, id).Error

}
func (a *Applications) FindAllApplication() error {
	return database.Find(&a.apps).Error
}
func (a *Application) FindByName(n string) error {
	return database.Where("name = ?", n).First(&a).Error
}

func (a *Application) Saver() error {
	return database.Save(a).Error
}

func (a *Application) Delete() error {
	return database.Delete(&a).Error
}
func (a *Application) Update(nom, version string) {
	if nom != "" {
		a.Nom = nom
	}
	if version != "" {
		a.Version = version
	}
}
