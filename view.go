package dovetail

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

// HTMLEnhancer adds attributes but doesn’t add children
type HTMLEnhancer interface {
	HTMLView
	enhances() bool
}

// Build takes an HTMLView and creates an html.Node
func Build(view HTMLView) *html.Node {
	node := &html.Node{}
	view.apply(node)
	return node
}

// Render takes an HTMLView and renders it and its tree to w
func Render(w io.Writer, views ...HTMLView) {
	for _, view := range views {
		html.Render(w, Build(view))
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

// Concat adds all the class names from another instance
func (classNames ClassNames) Concat(additions ClassNames) ClassNames {
	return append(classNames, additions...)
}

// Strings converts the class names to a single, space-separated string
func (classNames ClassNames) String() string {
	return strings.Join(classNames, " ")
}

// HTMLElementCore is shared by various components to perform much of the work of creating an HTML element node
type HTMLElementCore struct {
	classNames       ClassNames
	children         []HTMLView
	childTransformer func(node *html.Node) *html.Node
}

// Use the provided enhancers
func (core HTMLElementCore) Use(enhancers ...HTMLEnhancer) HTMLElementCore {
	for _, enhancer := range enhancers {
		core.children = append(core.children, enhancer)
	}
	return core
}

func (core HTMLElementCore) applyToNode(node *html.Node) {
	classNames := core.classNames

	for _, child := range core.children {
		switch child := child.(type) {
		case HTMLAttrView:
			child.apply(node)
		case HTMLClassNameView:
			classNames = classNames.Concat(child.classNames)
		case HTMLView:
			childNode := &html.Node{}
			child.apply(childNode)
			if core.childTransformer != nil {
				childNode = core.childTransformer(childNode)
			}
			node.AppendChild(childNode)
		}
	}

	if len(classNames) > 0 {
		node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: classNames.String()})
	}
}

// Heading lets you render h1, h2, h3, etc
type Heading struct {
	level       int
	elementCore HTMLElementCore
}

// H can be used to create <h1>, <h2>, etc. The first argument is the number
func H(level int, children ...HTMLView) HTMLView {
	return Heading{level: level, elementCore: HTMLElementCore{children: children}}
}

func (h Heading) apply(node *html.Node) {
	node.Type = html.ElementNode
	switch h.level {
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
		panic(fmt.Sprintf("Unsupported heading level %v", h.level))
	}

	h.elementCore.applyToNode(node)
}

// ButtonView makes <button>
type ButtonView struct {
	buttonType  string
	elementCore HTMLElementCore
}

func Button(children ...HTMLView) ButtonView {
	return ButtonView{elementCore: HTMLElementCore{children: children}}
}

func ButtonSubmit(input ButtonView) ButtonView {
	input.buttonType = "submit"
	return input
}

func SpecialButton(options ...func(input ButtonView) ButtonView) func(children ...HTMLView) ButtonView {
	return func(children ...HTMLView) ButtonView {
		button := Button(children...)
		for _, option := range options {
			button = option(button)
		}
		return button
	}
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

	button.elementCore.applyToNode(node)
}

//

// HTMLElementView can be adapted to many types of HTML elements
type HTMLElementView struct {
	tagName     string
	tagAtom     atom.Atom
	elementCore HTMLElementCore
}

func (el HTMLElementView) Use(enhancers ...HTMLEnhancer) HTMLElementView {
	el.elementCore = el.elementCore.Use(enhancers...)
	return el
}

func (basic HTMLElementView) Class(classNames ...string) HTMLElementView {
	basic.elementCore.classNames = append(basic.elementCore.classNames, classNames...)
	return basic
}

func (basic HTMLElementView) AddClasses(additions ClassNames) HTMLElementView {
	basic.elementCore.classNames = append(basic.elementCore.classNames, additions...)
	return basic
}

func (basic HTMLElementView) ChangeClasses(changer func(classNames ClassNames) ClassNames) HTMLElementView {
	basic.elementCore.classNames = changer(basic.elementCore.classNames)
	return basic
}

func (basic HTMLElementView) apply(node *html.Node) {
	node.Type = html.ElementNode
	node.Data = basic.tagName
	node.DataAtom = basic.tagAtom

	basic.elementCore.applyToNode(node)
}

func HTMLElementViewOf(tagName string, tagAtom atom.Atom, children []HTMLView) HTMLElementView {
	return HTMLElementView{
		tagName:     tagName,
		tagAtom:     tagAtom,
		elementCore: HTMLElementCore{children: children},
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
		tagName: "div",
		tagAtom: atom.Div,
		elementCore: HTMLElementCore{
			children:   children,
			classNames: classNames,
		},
	}
}

func Ul(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("ul", atom.Ul, children)
}

func Li(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("li", atom.Li, children)
}

func List(children ...HTMLView) HTMLElementView {
	return HTMLElementView{
		tagName: "ul",
		tagAtom: atom.Ul,
		elementCore: HTMLElementCore{
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
		},
	}
}

func Nav(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("nav", atom.Nav, children)
}

func P(children ...HTMLView) HTMLElementView {
	return HTMLElementViewOf("p", atom.P, children)
}

func Link(url string, children ...HTMLView) HTMLElementView {
	children = append(children, CustomAttr("href", url))

	return HTMLElementViewOf("a", atom.A, children)
}

func Img(srcURL string, alt string, enhancers ...HTMLEnhancer) HTMLElementView {
	view := HTMLElementViewOf("img", atom.Img, []HTMLView{CustomAttr("src", srcURL), CustomAttr("alt", alt)})
	view.elementCore = view.elementCore.Use(enhancers...)
	return view
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

func (HTMLAttrView) enhances() bool { return true }

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

// HTMLClassNameView allows adding to the class attribute
type HTMLClassNameView struct {
	classNames ClassNames
}

// This method is not actually used, instead the class names are all merged before setting the class attribute
func (view HTMLClassNameView) apply(node *html.Node) {
	node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: view.classNames.String()})
}

func (HTMLClassNameView) enhances() bool { return true }

// ClassName adds a class name
func ClassName(classNames ...string) HTMLClassNameView {
	return HTMLClassNameView{classNames: classNames}
}

// Class adds a class name
func Class(classNames ...string) HTMLClassNameView {
	return HTMLClassNameView{classNames: classNames}
}
