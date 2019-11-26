# Dovetail

HTML Components for Go

```go
Render(
  Header(
    Nav(
      NavLink("/", Text("Home")),
      NavLink("/about", Text("About")),
      NavLink("/pricing", Text("Pricing")),
      NavLink("/sign-in", Text("Sign In")),
      NavLink("/join", Text("Join")),
    ),
  ),
  Article(
    H(1, Text("Welcome")),
    Markdown(content),
  ),
)
```

## Define components

Simple components can be defined using functions:

```go
func PrimaryButton(text string) View {
	return Button(Text(text)).Class("btn btn-primary")
}
```

More complex components can be defined using structs that hold state:

```go
type PrimaryNavView struct {
  CurrentURL string
}

func PrimaryNav(currentURL string) HTMLText {
	return PrimaryNavView{CurrentURL: currentURL}
}

func (view PrimaryNavView) item(to string, text string) HTMLView {
  NavLink(to, Text(text), When(view.CurrentURL == to, Attr("aria-current", "page")))
}

func (view PrimaryNavView) Body() View {
	return Nav(
    view.item("/", "Home"),
    view.item("/about", "About"),
    view.item("/pricing", "Pricing"),
    view.item("/sign-in", "Sign In"),
    view.item("/join", "Join"),
  )
}
```

## Tailwind

Dovetail has native support for [Tailwind](https://tailwindcss.com/) classes

```go
Div().AddClasses(Tailwind(Pt8, Pb8, Text2XL, FontBold))
```