package personaldata

import "fmt"

type Personal struct {
	Name   string  // имя пользователяё
	Weight float64 // вес пользователя
	Height float64 // рост пользователя
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
