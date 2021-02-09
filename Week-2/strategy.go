package week2

import (
	"fmt"
	"strings"
)

type WeaponBehavior interface {
	useWeapon() string
}
type KnifeBehavior struct {}
type BowAndArrowBehavior struct {}
type AxeBehavior struct {}
type SwordBehavior struct {}

func (k *KnifeBehavior) useWeapon() string {
	return "knife... Striking"
}
func (b *BowAndArrowBehavior) useWeapon() string {
	return "bow and arrow... Striking"
}
func (a *AxeBehavior) useWeapon() string {
	return "axe... Striking"
}
func (b *SwordBehavior) useWeapon() string {
	return "sword... Striking"
}

type Character struct {
	hero Class
	weapon WeaponBehavior
}

func NewCharacter(hero Class, weapon WeaponBehavior) *Character {
	return &Character{hero, weapon}
}

func Weapon(weapon string) string {
	beg := int(strings.IndexAny(weapon,"."))
	end := int(strings.Index(weapon,"Behavior"))
	return weapon[beg+1:end]
}

func Hero(class string) string {
	beg := int(strings.IndexAny(class,"."))
	end := int(strings.Index(class,"{"))
	return class[beg+1:end]
}

func (c *Character) SetWeapon(weapon WeaponBehavior) {
	c.weapon = weapon
	str := fmt.Sprintf("%#v", weapon)
	str = Weapon(str)
	fmt.Println(c.hero.Display() + " is taking " + str + ". Everyone retreat! Now!!!")
}
func (c *Character) SetClass(class Class) {
	c.hero = class
	str := fmt.Sprintf("%#v", class)
	str = Hero(str)
	fmt.Println("The hero has just improved, now he is " + str)
}

type Class interface {
	Display() string
}
type Queen struct {}
func (q *Queen) Display() string {
	return "Queen"
}
type King struct {}
func (k *King) Display() string {
	return "King"
}
type Troll struct {}
func (t *Troll) Display() string {
	return "Troll"
}
type Knight struct {}
func (k *Knight) Display() string {
	return "Knight"
}

func (c *Character) Fight() string{
	return c.hero.Display() + " is using " + c.weapon.useWeapon()
}

func (c *Character) String() string {
	return c.hero.Display() + "is using" + c.weapon.useWeapon();
}

func stategy() {
	hero := Character{&Troll{}, &AxeBehavior{}}
	fmt.Println(hero.Fight())
	hero.SetWeapon(&BowAndArrowBehavior{})
	hero.SetClass(&Knight{})
	hero.SetWeapon(&SwordBehavior{})
	queen := Character{&Queen{}, &KnifeBehavior{}}
	fmt.Println(queen.Fight())
	hero.SetClass(&King{})
}