package Models

func (r *Service) Saver() error {
	return database.Save(r).Error
}

func (r *Service) FindById(id string) error {
	return database.First(r, id).Error
}

func (r *Services) FindAll() error {
	return database.Find(&r.Services).Error
}
