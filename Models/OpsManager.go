package Models

func (r *Ops) FindById(id string) error {
	err := database.First(r, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Ops) Saver() error {
	err := database.Save(r).Error
	return err
}
func (r *Opss) name() error {
	err := database.Find(r).Error
	return err
}
