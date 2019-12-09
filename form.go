package dovetail

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// FormHTMLView makes <form> with the provided Method and Action
// <form method="post" action="/pictures" enctype="multipart/form-data" class="mb-4 flex flex-row items-end">
type FormHTMLView struct {
	Method      string
	Action      string
	encType     string
	elementCore HTMLElementCore
}

// FormTo an action URL
func FormTo(action string, options ...func(form FormHTMLView) FormHTMLView) FormHTMLView {
	form := FormHTMLView{Method: "post", Action: action}
	for _, option := range options {
		form = option(form)
	}
	return form
}

func Multipart(form FormHTMLView) FormHTMLView {
	form.encType = "multipart/form-data"
	return form
}

// Multipart sets the `enctype` attribute to "multipart/form-data"
func (form FormHTMLView) Multipart() FormHTMLView {
	form.encType = "multipart/form-data"
	return form
}

// Add the provided views as children
func (form FormHTMLView) Add(view ...HTMLView) FormHTMLView {
	form.elementCore.children = append(form.elementCore.children, view...)
	return form
}

// With adds the provided views as children
func (form FormHTMLView) With(view ...HTMLView) FormHTMLView {
	form.elementCore.children = append(form.elementCore.children, view...)
	return form
}

func (form FormHTMLView) apply(node *html.Node) {
	node.Type = html.ElementNode
	node.Data = "form"
	node.DataAtom = atom.Form
	node.Attr = []html.Attribute{{Key: "method", Val: form.Method}, {Key: "action", Val: form.Action}}

	if form.encType != "" {
		node.Attr = append(node.Attr, html.Attribute{Key: "enctype", Val: form.encType})
	}

	form.elementCore.applyToNode(node)
}

type FieldHTMLView struct {
	formName       string
	inputType      string
	labelInnerView HTMLView
	elementCore    HTMLElementCore
}

func FieldLabelled(labelText string, options ...func(FieldHTMLView) FieldHTMLView) FieldHTMLView {
	field := FieldHTMLView{labelInnerView: Text(labelText)}
	for _, option := range options {
		field = option(field)
	}
	return field
}

func FieldNamed(formName string, options ...func(FieldHTMLView) FieldHTMLView) FieldHTMLView {
	field := FieldHTMLView{formName: formName}
	for _, option := range options {
		field = option(field)
	}
	return field
}

func InputType(inputType string) func(FieldHTMLView) FieldHTMLView {
	return func(field FieldHTMLView) FieldHTMLView {
		field.inputType = inputType
		return field
	}
}

var FileInput = InputType("file")

func InputNamed(inputName string, options ...func(FieldHTMLView) FieldHTMLView) func(FieldHTMLView) FieldHTMLView {
	return func(field FieldHTMLView) FieldHTMLView {
		field.formName = inputName
		for _, option := range options {
			field = option(field)
		}
		return field
	}
}

func (field FieldHTMLView) Type(inputType string) FieldHTMLView {
	field.inputType = inputType
	return field
}

func (field FieldHTMLView) Label(view HTMLView) FieldHTMLView {
	field.labelInnerView = view
	return field
}

func (field FieldHTMLView) Class(className string) FieldHTMLView {
	field.elementCore.classNames = field.elementCore.classNames.Class(className)
	return field
}

func (field FieldHTMLView) apply(node *html.Node) {
	inputType := field.inputType
	if inputType == "" {
		inputType = "text"
	}

	inputEl := &html.Node{
		Type:     html.ElementNode,
		Data:     "input",
		DataAtom: atom.Input,
		Attr:     []html.Attribute{{Key: "type", Val: inputType}, {Key: "name", Val: field.formName}},
	}

	spanEl := &html.Node{
		Type:     html.ElementNode,
		Data:     "span",
		DataAtom: atom.Span,
	}
	spanEl.AppendChild(Build(field.labelInnerView))

	node.Type = html.ElementNode
	node.Data = "label"
	node.DataAtom = atom.Label

	node.AppendChild(spanEl)
	node.AppendChild(inputEl)

	field.elementCore.applyToNode(node)
}
