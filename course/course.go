package course

import "fmt"

type Course struct {
	Name     string
	Price    float64
	IsFree   bool
	UserIds  []uint
	Classses map[uint]string
}

func (c *Course) PrintClasses() {
	text := "estos son los cursos a crusar de " + c.Name + " "
	for i, value := range c.Classses {
		if i < uint(len(c.Classses)) {
			text += fmt.Sprintf("%v: %v, ", i, value)
		} else {
			text += fmt.Sprintf("%v: %v.", i, value)
		}

	}

	fmt.Println(text)

}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
