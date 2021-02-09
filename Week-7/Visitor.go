package main

import "fmt"

type students interface {
	getName() string
	accept(teacher)
}

type Miras struct {
	name string
}

func (s *Miras) accept(t teacher) {
	t.visitForMiras(s)
}

func (s *Miras) getName() string {
	return s.name
}

type Beibarys struct {
	name string
}

func (s *Beibarys) accept(t teacher) {
	t.visitForBeibarys(s)
}

func (s *Beibarys) getName() string {
	return s.name
}

type Sanat struct {
	name string
}

func (s *Sanat) accept(t teacher) {
	t.visitForSanat(s)
}

func (s *Sanat) getName() string {
	return s.name
}

type teacher interface {
	visitForBeibarys(*Beibarys)
	visitForSanat(*Sanat)
	visitForMiras(*Miras)
}

type Olki struct {
	list []string
}

func (t *Olki) visitForBeibarys(s *Beibarys)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Olki) visitForSanat(s *Sanat)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Olki) visitForMiras(s *Miras)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Olki) showList() {
	for i, building := range t.list{
		fmt.Println(i, building)
	}
}

func visi() {
	miras := &Miras{"Miras"}
	beibarys := &Beibarys{"Beibarys"}
	sanat := &Sanat{"Sanat"}
	olki := Olki{[]string{}}
	olki.visitForBeibarys(beibarys)
	olki.visitForSanat(sanat)
	olki.visitForMiras(miras)
	olki.showList()
}