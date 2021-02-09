package week3
import "fmt"
type Employee struct {
	Name,Position string
	AnnualIncome int
}
// functional
func NewEmployeeFactory(position string,annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name,position,annualIncome}
	}
}
type EmployeeFactory struct {
	position string
	annualIncome int
}
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position,f.annualIncome}
}
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position,annualIncome}
}

func ex()  {
	developerFactory := NewEmployeeFactory("developer",850000)
	managerFactory := NewEmployeeFactory("manager",95000)
	bossFactory := NewEmployeeFactory2("CEO",150000)
	developer := developerFactory("Adam")
	developer2 := developerFactory("Mark")
	manager := managerFactory("Bob")
	bossFactory.annualIncome = 300000
	boss := bossFactory.Create("Sam")
	fmt.Println(developer)
	fmt.Println(developer2)
	fmt.Println(manager)
	fmt.Println(boss)
}