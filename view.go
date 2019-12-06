package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// HTMLView applies changes to an html.Node, such as making it into an element or text node, or adding attributes
type HTMLView interface {
	apply(node *html.Node)
}

func buildHTMLNode(view HTMLView) *html.Node {
	node := &html.Node{}
	view.apply(node)
	return node
}

// Render takes an HTMLView and renders it and its tree to w
func Render(w io.Writer, views ...HTMLView) {
	for _, view := range views {
		html.Render(w, buildHTMLNode(view))
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

func (text HTMLText) apply(node *html.Node) {
	node.Type = html.TextNode
	node.Data = text.Text
}

// ClassNames is a slice of class names
type ClassNames []string

// Class adds the passed class names
func (classNames ClassNames) Class(additions ...string) ClassNames {
	return append(classNames, additions...)
}

// Strings converts the class names to a single, space-separated string
func (classNames ClassNames) String() string {
	return strings.Join(classNames, " ")
}

// Heading lets you render h1, h2, h3, etc
type Heading struct {
	Level int
	Child HTMLView
}

// H can be used to create <h1>, <h2>, etc. The first argument is the number
func H(level int, child HTMLView) HTMLView {
	return Heading{Level: level, Child: child}
}

func (h Heading) apply(node *html.Node) {
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
	base             html.Node
	children         []HTMLView
	classNames       ClassNames
	childTransformer func(node *html.Node) *html.Node
}

func (basic HTMLElementView) Class(classNames ...string) HTMLElementView {
	basic.classNames = append(basic.classNames, classNames...)
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

	classNames := basic.classNames

	for _, child := range basic.children {
		switch child := child.(type) {
		case HTMLAttrView:
			child.apply(node)
		case HTMLClassNameView:
			classNames = classNames.Class(child.ClassName)
		case HTMLView:
			childNode := &html.Node{}
			child.apply(childNode)
			if basic.childTransformer != nil {
				childNode = basic.childTransformer(childNode)
			}
			node.AppendChild(childNode)
		}
	}

	if len(classNames) > 0 {
		node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: classNames.String()})
	}
}

func HTMLElementViewOf(tagName string, tagAtom atom.Atom, children []HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     tagName,
			DataAtom: tagAtom,
		},
		children: children,
	}
}

func Main(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("main", atom.Main, children)
}

func Header(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("header", atom.Header, children)
}

func Footer(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("footer", atom.Footer, children)
}

func Section(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("section", atom.Section, children)
}

func Article(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("article", atom.Article, children)
}

func Aside(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("aside", atom.Aside, children)
}

func Div(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("div", atom.Div, children)
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

func Ul(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "ul",
			DataAtom: atom.Ul,
		},
		children: children,
	}
}

func Li(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "li",
			DataAtom: atom.Li,
		},
		children: children,
	}
}

func List(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		base: html.Node{
			Type:     html.ElementNode,
			Data:     "ul",
			DataAtom: atom.Ul,
		},
		children: children,
		childTransformer: func(node *html.Node) *html.Node {
			li := &html.Node{
				Type:     html.ElementNode,
				Data:     "li",
				DataAtom: atom.Li,
			}
			li.AppendChild(node)
			return li
		},
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

func (attrView HTMLAttrView) apply(node *html.Node) {
	node.Attr = append(node.Attr, html.Attribute{Key: attrView.Key, Val: attrView.Value})
}

// AriaAttr is for aria attributes such as aria-label or aria-current
func AriaAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: "aria-" + key, Value: value}
}

// AriaLabel sets the aria-label attribute
func AriaLabel(value string) HTMLAttrView {
	return HTMLAttrView{Key: "aria-label", Value: value}
}

// DataAttr is for data attributes such as data-testid or data-anything
func DataAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: "data-" + key, Value: value}
}

// CustomAttr is for data attributes such as href or src
func CustomAttr(key string, value string) HTMLAttrView {
	return HTMLAttrView{Key: key, Value: value}
}

// TODO? replace HTMLClassNameView with HTMLAttrView with AppendWithSpace strategy
// FIXME: this completely replaces existing class attributes
// HTMLClassNameView allows adding to the class attribute
type HTMLClassNameView struct {
	ClassName string
}

func (view HTMLClassNameView) apply(node *html.Node) {
	node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: view.ClassName})
}

// ClassName adds a class name
func ClassName(className string) HTMLClassNameView {
	return HTMLClassNameView{ClassName: className}
}
