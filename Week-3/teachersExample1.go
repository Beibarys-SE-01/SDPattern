package week3

import (
	"fmt"
	"strings"
)
const indentSize = 2

type HtmlElement struct {
	name, text string
	elements []HtmlElement
}
func (e *HtmlElement) String() string {
	return e.string(0)
}
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent+1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}
// step by step call API of builder
type HtmlBuilder struct {
	root HtmlElement
	rootName string
}
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	b:=HtmlBuilder{rootName: rootName, root: HtmlElement{rootName,"",[]HtmlElement{}}}
	return &b
}
func (b *HtmlBuilder) String() string {
	return b.root.String()
}
// Fluent Interface
func (b *HtmlBuilder) AddChildFluent(childName,childText string) *HtmlBuilder {
	e := HtmlElement{name: childName,text: childText, elements: []HtmlElement{}}
	b.root.elements = append(b.root.elements,e)
	return b
}

func asd() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li","AITU").
		AddChildFluent("li","KBTU").
		AddChildFluent("li","IITU")

	fmt.Println(b)
}