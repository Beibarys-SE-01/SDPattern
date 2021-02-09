package week3

import (
	"fmt"
	"sync"
)

type person struct {
	name, gender, race string
}

func (p *person) String() string {
	return "Your name: " + p.name + "\n" + "Your gender: " + p.gender + "\n" + "Your race: " + p.race + "\n"
}

var once2 sync.Once
var instance2 *person
func CreateYourself(name, gender, race string) *person  {
	once.Do(func() {
		you := person{name, gender, race}
		instance2 = &you
		fmt.Println("Hi, you have just been reborn")
	})
	fmt.Println("It's too late to change something, you've already been reborn")
	return instance2
}

func main() {
	i := 1
	fmt.Println("Hello, you were reborn in another world!")
	fmt.Println("Enter your name, gender and race for new live. REMEMBER YOU HAVE ONLY 1 CHANCE")
	for i != 0 {
		var name string
		var gender string
		var race string
		fmt.Scan(&name)
		fmt.Scan(&gender)
		fmt.Scan(&race)
		you := CreateYourself(name, gender, race)
		fmt.Println(you)
		fmt.Println("Do you want to try again? Click 1 if yes, else 0")
		fmt.Scan(&i)
	}
	fmt.Println("Welcome to the world")
	/*you := CreateYourself("Beibarys", "male", "mongoloid")
	fmt.Println(you)
	tset := CreateYourself("te", "dfd", "vdbf")
	fmt.Println(tset)*/
}
