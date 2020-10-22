package Models

import "fmt"

func (a *Application) FindById(id string) error {
	database.First(a, id)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}
func FindAllOordinateur() (a []Application, err error) {
	result := database.Find(&a)
	fmt.Println(result)
	err = result.Error
	return o, err
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
