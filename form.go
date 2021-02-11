package dovetail

import (
	"strconv"

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

type FieldInputProps struct {
	name         string
	inputType    string
	defaultValue string
	rows         int
	core         HTMLElementCore
}

type FieldHTMLView struct {
	inputProps     FieldInputProps
	labelInnerView HTMLView
	labelCore      HTMLElementCore
	inputCore      HTMLElementCore
}

type FieldOption interface {
	applyToField(field FieldHTMLView) FieldHTMLView
}

func FieldLabelled(labelText string, option FieldOption, children ...HTMLEnhancer) FieldHTMLView {
	field := FieldHTMLView{labelInnerView: Text(labelText)}
	field = option.applyToField(field)
	for _, child := range children {
		field.labelCore.children = append(field.labelCore.children, child)
	}
	return field
}

type FieldTextInputOption func(FieldHTMLView) FieldHTMLView

func Textbox(inputName string, options ...FieldTextInputOption) FieldTextInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field.inputProps.name = inputName
		for _, option := range options {
			field = option(field)
		}
		return field
	}
}

func (option FieldTextInputOption) Use(enhancers ...HTMLEnhancer) FieldTextInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field = option(field)
		field.inputProps.core = field.inputProps.core.Use(enhancers...)
		return field
	}
}

func (option FieldTextInputOption) Rows(rows int) FieldTextInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field = option(field)
		field.inputProps.rows = rows
		return field
	}
}

func (option FieldTextInputOption) DefaultValue(value string) FieldTextInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field = option(field)
		field.inputProps.defaultValue = value
		return field
	}
}

func (option FieldTextInputOption) applyToField(field FieldHTMLView) FieldHTMLView {
	return option(field)
}

type FieldFileInputOption func(FieldHTMLView) FieldHTMLView

func FileInput(inputName string, options ...FieldFileInputOption) FieldFileInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field.inputProps.inputType = "file"
		field.inputProps.name = inputName
		for _, option := range options {
			field = option(field)
		}
		return field
	}
}

func (option FieldFileInputOption) Use(enhancers ...HTMLEnhancer) FieldFileInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field = option(field)
		field.inputProps.core = field.inputProps.core.Use(enhancers...)
		return field
	}
}

func (option FieldFileInputOption) applyToField(field FieldHTMLView) FieldHTMLView {
	return option(field)
}

type FieldNumberInputOption func(FieldHTMLView) FieldHTMLView

func NumberInput(inputName string, options ...FieldNumberInputOption) FieldNumberInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field.inputProps.inputType = "number"
		field.inputProps.name = inputName
		for _, option := range options {
			field = option(field)
		}
		return field
	}
}

func (option FieldNumberInputOption) Use(enhancers ...HTMLEnhancer) FieldNumberInputOption {
	return func(field FieldHTMLView) FieldHTMLView {
		field = option(field)
		field.inputProps.core = field.inputProps.core.Use(enhancers...)
		return field
	}
}

func (option FieldNumberInputOption) applyToField(field FieldHTMLView) FieldHTMLView {
	return option(field)
}

func (field FieldHTMLView) Class(className string) FieldHTMLView {
	field.labelCore.classNames = field.labelCore.classNames.Class(className)
	return field
}

func (field FieldHTMLView) apply(node *html.Node) {
	inputType := field.inputProps.inputType
	if inputType == "" {
		inputType = "text"
	}

	var inputEl *html.Node
	if field.inputProps.rows > 0 {
		inputEl = &html.Node{
			Type:     html.ElementNode,
			Data:     "textarea",
			DataAtom: atom.Textarea,
			Attr:     []html.Attribute{{Key: "name", Val: field.inputProps.name}, {Key: "rows", Val: strconv.Itoa(field.inputProps.rows)}},
		}

		if field.inputProps.defaultValue != "" {
			inputEl.AppendChild(&html.Node{
				Type: html.TextNode,
				Data: field.inputProps.defaultValue,
			})
		}
	} else {
		inputEl = &html.Node{
			Type:     html.ElementNode,
			Data:     "input",
			DataAtom: atom.Input,
			Attr:     []html.Attribute{{Key: "type", Val: inputType}, {Key: "name", Val: field.inputProps.name}},
		}

		if field.inputProps.defaultValue != "" {
			inputEl.Attr = append(inputEl.Attr, html.Attribute{Key: "value", Val: field.inputProps.defaultValue})
		}
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

	field.inputProps.core.applyToNode(inputEl)
	node.AppendChild(inputEl)

	field.labelCore.applyToNode(node)
}
