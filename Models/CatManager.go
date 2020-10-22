package Models

func (c *Cat) FindById(id string) error {
	database.First(c, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func FindAllCat() (c []Car, err error) {
	result := database.Find(&c)
	err = result.Error
	return c, err
}
func (c *Cat) FindByName(n string) error {
	database.Where("name = ?", n).First(&c)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Cat) Saver() error {
	database.Save(c)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Cat) Delete() error {
	database.Delete(&c)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
