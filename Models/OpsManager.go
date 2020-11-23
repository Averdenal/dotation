package Models

func (r *Ops) FindById(id string) error {
	return database.First(r, id).Error
}
func (r *Ops) Saver() error {
	return database.Save(r).Error
}
func (r *Opss) FindAllOs() error {
	return database.Find(&r.Ops).Error
}
