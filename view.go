package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type View interface {
	Body() View
}

type viewWriter interface {
	render(w io.Writer)
}

type HTMLView interface {
	Body() View
	apply(node *html.Node)
}

func buildHTMLNode(view HTMLView) *html.Node {
	node := &html.Node{}
	view.apply(node)
	return node
}

// Render takes a View and renders it and its tree to w
func Render(w io.Writer, view View) {
	switch view := view.(type) {
	case HTMLView:
		html.Render(w, buildHTMLNode(view))
	case viewWriter:
		view.render(w)
	case View:
		body := view.Body()
		Render(w, body)
	}
}

// HTMLText represents an HTML text node
type HTMLText struct {
	Text string
}

// Text makes a html text node with the given content
func Text(text string) HTMLText {
	return HTMLText{text}
}

func (_ HTMLText) Body() View {
	return nil
}

func (text HTMLText) apply(node *html.Node) {
	node.Type = html.TextNode
	node.Data = text.Text
}

// ClassNames is a slice of class names
type ClassNames []string

func (classNames ClassNames) Class(additions ...string) ClassNames {
	return append(classNames, additions...)
}

// Strings converts the class names to a single, space-separated string
func (classNames ClassNames) String() string {
	return strings.Join(classNames, " ")
}

// H lets you render h1, h2, h3, etc
type H struct {
	Level int
	Child HTMLView
}

func (_ H) Body() View {
	return nil
}

func (h H) apply(node *html.Node) {
	node.Type = html.ElementNode
	switch h.Level {
	case 1:
		node.Data = "h1"
		node.DataAtom = atom.H1
	case 2:
		node.Data = "h2"
		node.DataAtom = atom.H2
	case 3:
		node.Data = "h3"
		node.DataAtom = atom.H3
	case 4:
		node.Data = "h4"
		node.DataAtom = atom.H4
	case 5:
		node.Data = "h5"
		node.DataAtom = atom.H5
	case 6:
		node.Data = "h6"
		node.DataAtom = atom.H6
	default:
		panic(fmt.Sprintf("Unsupported heading level %v", h.Level))
	}

	node.AppendChild(buildHTMLNode(h.Child))
}

// ButtonView makes <button>
type ButtonView struct {
	Child      HTMLView
	buttonType string
	classNames ClassNames
}

func Button(child HTMLView) ButtonView {
	return ButtonView{Child: child}
}

func (button ButtonView) AddClasses(additions ClassNames) ButtonView {
	button.classNames = append(button.classNames, additions...)
	return button
}

func (_ ButtonView) Body() View { return nil }

func (button ButtonView) Submit() ButtonView {
	button.buttonType = "submit"
	return button
}

func (button ButtonView) apply(node *html.Node) {
	buttonType := button.buttonType
	if buttonType == "" {
		buttonType = "button"
	}

	node.Type = html.ElementNode
	node.Data = "button"
	node.DataAtom = atom.Button
	node.Attr = []html.Attribute{{Key: "type", Val: buttonType}}

	node.AppendChild(buildHTMLNode(button.Child))
}

// HTMLElementView can be adapted to many types of HTML elements
type HTMLElementView struct {
	base       html.Node
	children   []HTMLView
	classNames ClassNames
}

func (_ HTMLElementView) Body() View { return nil }

func (basic HTMLElementView) Class(classNames ...string) HTMLElementView {
	basic.classNames = append(basic.classNames, classNames...)
	// basic = append(*basic, classNames...)
	return basic
}

func (basic HTMLElementView) AddClasses(additions ClassNames) HTMLElementView {
	basic.classNames = append(basic.classNames, additions...)
	return basic
}

func (basic HTMLElementView) ChangeClasses(changer func(classNames ClassNames) ClassNames) HTMLElementView {
	basic.classNames = changer(basic.classNames)
	return basic
}

func (basic HTMLElementView) apply(node *html.Node) {
	node.Type = html.ElementNode
	node.Data = basic.base.Data
	node.DataAtom = basic.base.DataAtom
	node.Attr = basic.base.Attr

	for _, child := range basic.children {
		switch child := child.(type) {
		case HTMLAttrView:
			child.apply(node)
		case HTMLView:
			childNode := &html.Node{}
			child.apply(childNode)
			node.AppendChild(childNode)
		}
	}

	if len(basic.classNames) > 0 {
		node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: basic.classNames.String()})
	}
}

func Header(children ...HTMLView) HTMLView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "header",
			DataAtom: atom.Header,
		},
		children: children,
	}
}

func Div(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "div",
			DataAtom: atom.Div,
		},
		children: children,
	}
}

func DivWithClasses(classNames ClassNames, children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "div",
			DataAtom: atom.Div,
		},
		children:   children,
		classNames: classNames,
	}
}

func Nav(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "nav",
			DataAtom: atom.Nav,
		},
		children: children,
	}
}

func Link(url string, children ...HTMLView) HTMLElementView {
	children = append(children, CustomAttr("href", url))

	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "a",
			DataAtom: atom.A,
		},
		children: children,
	}
}

//

// HTMLAttrView allows setting HTML attributes
type HTMLAttrView struct {
	Key   string
	Value string
}

func (_ HTMLAttrView) Body() View {
	return nil
}

func (attrView HTMLAttrView) apply(node *html.Node) {
	node.Attr = append(node.Attr, html.Attribute{Key: attrView.Key, Val: attrView.Value})
}

// AriaAttr is for aria attributes such as aria-label or aria-current
func AriaAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: "aria-" + key, Value: value}
}

// DataAttr is for data attributes such as data-testid or data-anything
func DataAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: "data-" + key, Value: value}
}

// CustomAttr is for data attributes such as href or src
func CustomAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: key, Value: value}
}
