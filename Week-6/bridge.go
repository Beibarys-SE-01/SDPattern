package main

import "fmt"

type human1 interface {
	setFood(food)
	eat()
	checkFood() error
}

type baby struct {
	food food
}

func (b *baby) setFood(f food) {
	b.food = f
}

func (b *baby) eat() {
	b.food.takeFood()
	fmt.Println("Let's eat it")
}

func (b baby) checkFood() error {
	if b.food == nil {
		return fmt.Errorf("There is no food!")
	}
	return nil
}

type teenager struct {
	food food
}

func (t *teenager) setFood(f food) {
	t.food = f
}

func (t *teenager) eat() {
	t.food.takeFood()
	fmt.Println("Let's eat it")
}

func (t *teenager) checkFood() error {
	if t.food == nil {
		return fmt.Errorf("There is no food!")
	}
	return nil
}

type food interface {
	takeFood()
}

type milk struct {}

func (m *milk) takeFood() {
	fmt.Println("Milk was taken")
}

type pizza struct {}

func (p *pizza) takeFood() {
	fmt.Println("Pizza was taken")
}

type client struct {}

func (c client) giveNewFoodToHuman(human human1, food food) {
	err := human.checkFood()
	if err == nil {
		human.setFood(food)
		human.eat()
	} else {
		human.setFood(food)
		human.eat()
	}
}

func (c client) giveFoodToHuman(human human1) {
	err := human.checkFood()
	if err == nil {
		human.eat()
	} else {
		fmt.Println(err)
	}
}

func bridge() {
	client := client{}

	pizza := &pizza{}
	teenager := &teenager{}
	client.giveNewFoodToHuman(teenager, pizza)

	milk := &milk{}
	baby := &baby{}

	client.giveFoodToHuman(baby)

	client.giveNewFoodToHuman(baby, milk)

}