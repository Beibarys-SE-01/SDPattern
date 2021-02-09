package week3

import "strings"
type email struct {
	from,to,subject,body string
}
type EmailBuilder struct {
	email email
}
func (b *EmailBuilder) From(from string) *EmailBuilder {
	//Validation
	if !strings.Contains(from,"@"){
		panic("email should contain @")
	}
	b.email.from = from
	return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
	//Validation
	if !strings.Contains(to,"@"){
		panic("email should contain @")
	}
	b.email.to = to
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {

	b.email.subject = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}
func sendEmail(email *email) {
	// smtp
	//
}
type build func(*EmailBuilder)
func SendEmail(action build)  {
	builder := EmailBuilder{}
	action(&builder)
	sendEmail(&builder.email)
}
func ttes()  {
	SendEmail(func(b *EmailBuilder) {
		b.From("test@test.kz").
			To("tset@tset.zk").
			Subject("some Tittle").
			Body("asdasdasd")
	})


}