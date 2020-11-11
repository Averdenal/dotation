package Models

func (r *Service) Saver() error {
	database.Save(r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Service) FindById(id string) error {
	err := database.First(r, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Services) FindAll() error {
	err := database.Find(&r.Services).Error
	if err != nil {
		return err
	}
	return nil
}
