package Models

func (c *Cat) FindById(id string) error {
	return database.First(c, id).Error
}
func (c *Cats) FindAllCat() error {
	return database.Find(&c.Cats).Error
}
func (c *Cat) FindByName(n string) error {
	return database.Where("name = ?", n).First(&c).Error
}

func (c *Cat) Saver() error {
	return database.Save(c).Error
}

func (c *Cat) Delete() error {
	return database.Delete(&c).Error
}

func (c *Cat) Update(typeCat, model string) {
	if typeCat != "" {
		c.Type = typeCat
	}
	if model != "" {
		c.Models = model
	}
}
