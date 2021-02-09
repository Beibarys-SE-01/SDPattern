package week3

import "fmt"
type Person struct {
	//Address
	StreetAddress, Postcode, City string
	//Work
	CompanyName,Position string
	AnnualIncome int
}
type PersonBuilder struct {
	person *Person
}
func (b *PersonBuilder) Lives() *PersonAddressBuilder{
	return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder{
	return &PersonJobBuilder{*b}
}
func (b *PersonBuilder) Build() *Person {
	return b.person
}
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
type PersonAddressBuilder struct {
	PersonBuilder
}
func (ab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	ab.person.StreetAddress = streetAddress
	return ab
}
func (ab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	ab.person.City = city
	return ab
}
func (ab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	ab.person.Postcode = postcode
	return ab
}

type PersonJobBuilder struct {
	PersonBuilder
}
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}
func test()  {
	pb := NewPersonBuilder()
	pb.Lives().In("Astana").WithPostcode("100600").Works().Earning(1000)

	pb.Works().At("telekom").AsA("developer").Lives().At("Gagarina 100")

	p := pb.Build()
	fmt.Println(*p)
}