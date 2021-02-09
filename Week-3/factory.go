package week3

import (
	"fmt"
	"strconv"
)

type Car struct {
	Brand string
	Model string
	Year int
	Cost string
}

func NewCarFactory(brand string) func(string, int, string) *Car {
	return func(model string,year int, cost string) *Car {
		return &Car{brand,model,year, cost}
	}
}
func (c *Car) String() string {
	return c.Brand + " " + c.Model + " " + strconv.Itoa(c.Year) + " will cost " + c.Cost
}

func factory()  {
	BMWFactory := NewCarFactory("BMW")
	x6 := BMWFactory("x6", 2020, "13 million")
	fmt.Println(x6)
	toyotaFactory := NewCarFactory("TOYOTA")
	camry := toyotaFactory("50", 2012, "5 million")
	fmt.Println(camry)
}