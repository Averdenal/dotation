package Models

func (u *User) FindById(id string) error {
	database.First(u, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func (u *Users) FindAllUser() error {
	result := database.Find(&u.Users)
	return result.Error
}
func (u *User) FindByName(n string) error {
	database.Where("name = ?", n).First(&u)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Saver() error {
	database.Save(u)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	database.Delete(&u)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
