package Models

import "fmt"

func FindById(id int) (o Ordinateur) {
	database.First(&o, id)
	return
}
func FindAll() (o []Ordinateur, err error) {
	result := database.Find(&o)
	fmt.Println(result)
	err = result.Error
	return o, err
}
func FindByName(n string) (o Ordinateur) {
	database.Where("name = ?", n).First(&o)
	return
}

func (r *Ordinateur) Saver() {
	database.Save(*r)
}

func (r *Ordinateur) Delete() {
	database.Delete(&r)
}
