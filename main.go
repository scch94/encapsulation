package main

import (
	"fmt"

	"github.com/scch94/encapsulation/course"
)

func main() {
	//Go := Course{
	Go := course.Course{
		Name:    "go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIds: []uint{12, 56, 89},
		Classses: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	}
	Go.ChangePrice(11.2)
	fmt.Println(Go.Price)
	Go.PrintClasses()
}
