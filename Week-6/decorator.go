package main

import "fmt"

type salary interface {
	getSalary() int
}

type human struct {
	salary int
}

func newHuman() *human {
	return &human{50000}
}

func (h *human) getSalary() int {
	return h.salary
}

type student struct {
	human *human
}

func (s *student) getSalary() int {
	return s.human.getSalary()
}

type junior struct {
	human *human
}

func (j *junior) getSalary() int {
	return j.human.getSalary() * 2
}

type middle struct {
	human *human
}

func (m *middle) getSalary() int {
	return m.human.getSalary() * 3
}

type senior struct {
	human *human
}

func (s *senior) getSalary() int {
	return s.human.getSalary() * 5
}

func decorator() {
	human := newHuman()

	student := student{human}
	fmt.Printf("Beacuse of you are a student and have low experinece, we will pay you minimum %d\n", student.getSalary())

	junior := junior{human}
	fmt.Printf("Beacuse of you are a junior and have 1 year of experinece, we will pay %d\n", junior.getSalary())

	middle := middle{human}
	fmt.Printf("Beacuse of you are a middle and have 3 years of experinece, we will pay %d\n", middle.getSalary())

	senior := senior{human}
	fmt.Printf("Beacuse of you are a senior and have 5 years of experinece, we will pay %d\n", senior.getSalary())
}