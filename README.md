# Dovetail

_CURRENTLY IN ALPHA, API WILL CHANGE_

[Documentation](https://godoc.org/github.com/RoyalIcing/dovetail)

User- and developer-friendly HTML Components for Go, inspired by SwiftUI/React/Elm.

- Produce accessible HTML markup easily.
- Make components with functions.
- Supports ARIA and data attributes.
- Conveniences for forms and lists.

```go
Render(
  w,
  Div(
    Header(
      Nav(
        AriaLabel("Primary"),
        List(
          Link("/", Text("Home")),
          Link("/about", Text("About")),
          Link("/pricing", Text("Pricing")).Use(AriaAttr("current", "page")),
          Link("/sign-in", Text("Sign In")),
          Link("/join", Text("Join")),
        ),
      ),
    ),
    Main(
      Article(
        H(1, Text("Welcome")),
        P("Render HTML using components with Go"),
      ),
      FormTo("/newsletter").With(
        FieldLabelled("Email", TextInput("email")),
        SubmitButton(Text("Sign up for the newsletter")),
      ),
    ),
  )
)
```

## Provided components

Type: `HTMLElementView`

### Landmarks

- `Main(children ...HTMLView)` — [`<main>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main)
- `Nav(children ...HTMLView)` — [`<nav>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav)
- `Header(children ...HTMLView)` — [`<header>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header)
- `Footer(children ...HTMLView)` — [`<footer>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer)
- `Section(children ...HTMLView)` — [`<section>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section)
- `Article(children ...HTMLView)` — [`<article>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article)
- `Aside(children ...HTMLView)` — [`<aside>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside)

### Structure

- `Div(children ...HTMLView)` — [`<div>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div)
- `List(children ...HTMLView)` — `<ul><li>{ children[0] }</li>…<li>{ children[n] }</li></ul>`
- `Ul(children ...HTMLView)` — [`<ul>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul)
- `Li(children ...HTMLView)` — [`<li>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li)
- Coming soon: `<ol>`, `<dl>`

### Elements

- `Link(url string, children ...HTMLView)` — [`<a href="{ url }">{ children }</a>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a)
- `Button(children ...HTMLView)` — [`<button type="button">{ children }</button>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button)
- `Img(srcUrl string, alt string, enhancers ...HTMLEnhancer)` — `<img src="{ srcUrl }" alt="{ alt }" {...enhancers}>`
- `P(children ...HTMLView)` — `<p>`
- `TextWith(text string, enhancers ...HTMLEnhancer)` — `<span {...enhancers}>{ text }</span>`
- `Noscript(children ...HTMLView)` — `<noscript>{ children }</noscript>`

### Forms

- `FormTo(action string, options ...func(form FormHTMLView) FormHTMLView)` — [`<form>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form)
  - `Multipart(form FormHTMLView) FormHTMLView` — [`<form enctype="multipart/form-data">`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form#attr-enctype)
- `SubmitButton(children ...HTMLView)` — [`<button type="submit">{ children }</button>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button#attr-type)

### Text nodes

- `Text(text string)` — [HTML text node](https://developer.mozilla.org/en-US/docs/Web/API/Text)

### Logic

- `When(when bool, view HTMLView)` — renders the provided `view` only if `when` is `true`

### Custom

- `HTMLElementViewOf(tagName string, tagAtom atom.Atom, children []HTMLView)` — custom html element

## Attributes

- `AriaAttr` — [`aria-*`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/ARIA_Techniques#States_and_properties)
- `AriaLabel` — [`aria-label`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/ARIA_Techniques/Using_the_aria-label_attribute)
- `CustomAttr` — custom HTML attributes
- `DataAttr` — [`data-*` attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/data-*)

## Define components

Components are defined using functions. These functions can take any number of arguments, and return a composite of other components.

```go
func PrimaryButton(text string) View {
	return Button(Text(text)).Class("btn btn-primary")
}
```

## Performance

While not trying to be the fastest HTML producer possible, Dovetail aims to be faster than `html/template` to parse and execute.

Run `make test_bench` to see how Dovetail performs to produce a variety of HTML components. Here are results on a 2016 15″ MacBook Pro:

```
go test -p 1 -timeout 30s -bench="Bench" -benchmem -v -run "Bench" ./...
goos: darwin
goarch: amd64
pkg: github.com/RoyalIcing/dovetail
BenchmarkTailwindJustDiv-8                  	 4758601	       247 ns/op	     272 B/op	       2 allocs/op
BenchmarkTailwind0Classes-8                 	 4620313	       249 ns/op	     272 B/op	       2 allocs/op
BenchmarkTailwind2Classes-8                 	 2074034	       580 ns/op	     400 B/op	       6 allocs/op
BenchmarkTailwind4Classes-8                 	 1713938	       694 ns/op	     464 B/op	       7 allocs/op
BenchmarkTailwind8Classes-8                 	 1349778	       887 ns/op	     640 B/op	       8 allocs/op
BenchmarkTailwindAddClasses2Classes-8       	 1990402	       601 ns/op	     416 B/op	       6 allocs/op
BenchmarkTailwindAddClasses4Classes-8       	 1771840	       677 ns/op	     480 B/op	       6 allocs/op
BenchmarkTailwindAddClasses8Classes-8       	 1504893	       783 ns/op	     656 B/op	       6 allocs/op
BenchmarkTailwindDivWithClasses8Classes-8   	 1757682	       729 ns/op	     528 B/op	       5 allocs/op
BenchmarkTailwindChangeClasses2Classes-8    	 1797513	       643 ns/op	     464 B/op	       8 allocs/op
BenchmarkTailwindChangeClasses4Classes-8    	 1533936	       790 ns/op	     560 B/op	       9 allocs/op
BenchmarkTailwindChangeClasses8Classes-8    	 1234028	       970 ns/op	     800 B/op	      10 allocs/op
BenchmarkText-8                             	 7639869	       154 ns/op	     128 B/op	       2 allocs/op
BenchmarkHeader-8                           	 4320638	       251 ns/op	     272 B/op	       2 allocs/op
BenchmarkDiv-8                              	 4858587	       242 ns/op	     272 B/op	       2 allocs/op
BenchmarkDivWithClasses1-8                  	 2697412	       445 ns/op	     336 B/op	       4 allocs/op
BenchmarkDivWithClasses2Together-8          	 2296304	       519 ns/op	     368 B/op	       5 allocs/op
BenchmarkDivWithClasses2-8                  	 2066386	       576 ns/op	     384 B/op	       6 allocs/op
BenchmarkDivWithChildClassNames-8           	 1734169	       672 ns/op	     448 B/op	       9 allocs/op
BenchmarkButton-8                           	 2446404	       478 ns/op	     352 B/op	       5 allocs/op
BenchmarkButtonSubmit-8                     	 2451301	       480 ns/op	     352 B/op	       5 allocs/op
BenchmarkH1-8                               	 2267946	       521 ns/op	     448 B/op	       6 allocs/op
PASS
ok  	github.com/RoyalIcing/dovetail	39.216s
```

## Why?

Because I want to create server-rendered web apps that I can host cheaply on GCP App Engine Standard.

I want something that is both user friendly (quick to load, accessible) while also being developer friendly.
