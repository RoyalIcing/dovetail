package main

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestRender(t *testing.T) {
	t.Run("Rendering Example in Readme", func(t *testing.T) {
		s := subjectAsString(Div(
			Header(
				Nav(
					AriaLabel("Primary"),
					List(
						Link("/", Text("Home")),
						Link("/about", Text("About")),
						Link("/pricing", Text("Pricing")),
						Link("/sign-in", Text("Sign In")),
						Link("/join", Text("Join")),
					),
				),
			),
			Main(
				Article(
					H(1, Text("Welcome")),
					Div(Text("markdown")),
				),
			),
		))

		t.Run(`it renders`, func(t *testing.T) {
			assert.Equal(t, s, `<div><header><nav aria-label="Primary"><ul><li><a href="/">Home</a></li><li><a href="/about">About</a></li><li><a href="/pricing">Pricing</a></li><li><a href="/sign-in">Sign In</a></li><li><a href="/join">Join</a></li></ul></nav></header><main><article><h1>Welcome</h1><div>markdown</div></article></main></div>`)
		})
	})

	t.Run("Rendering Link", func(t *testing.T) {
		s := subjectAsString(Link("https://example.org/", Text("Hello")))

		t.Run(`it renders <a> with href and child text`, func(t *testing.T) {
			assert.Equal(t, s, `<a href="https://example.org/">Hello</a>`)
		})
	})

	// t.Run("Rendering LinkView", func(t *testing.T) {
	// 	s := subjectAsString(LinkView{To: "https://example.org/", Text: "Hello"})

	// 	t.Run(`it renders <a> with href and child text`, func(t *testing.T) {
	// 		assert.Equal(t, s, `<a href="https://example.org/">Hello</a>`)
	// 	})
	// })

	t.Run("Rendering Nav with items and aria label Primary", func(t *testing.T) {
		s := subjectAsString(Nav(AriaLabel("Primary"), Ul(Li(Link("/about", Text("About"))))))

		t.Run(`it renders <nav> with aria-label and one link in a list`, func(t *testing.T) {
			assert.Equal(t, s, `<nav aria-label="Primary"><ul><li><a href="/about">About</a></li></ul></nav>`)
		})
	})

	t.Run("Rendering Ul with Li then Link and Button children", func(t *testing.T) {
		s := subjectAsString(
			Ul(
				Li(Link("/about", Text("About"))),
				Li(Button(Text("Sign Out"))),
			),
		)

		t.Run(`it renders <ul> with 2 <li> with <a> and <button> inside`, func(t *testing.T) {
			assert.Equal(t, s, `<ul><li><a href="/about">About</a></li><li><button type="button">Sign Out</button></li></ul>`)
		})
	})

	t.Run("Rendering List with Link and Button children", func(t *testing.T) {
		s := subjectAsString(
			List(
				Link("/about", Text("About")),
				Button(Text("Sign Out")),
			),
		)

		t.Run(`it renders <ul> with 2 <li> with <a> and <button> inside`, func(t *testing.T) {
			assert.Equal(t, s, `<ul><li><a href="/about">About</a></li><li><button type="button">Sign Out</button></li></ul>`)
		})
	})

	t.Run("Rendering H1", func(t *testing.T) {
		s := subjectAsString(H(1, Text("Hello")))

		t.Run(`it renders <h1> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h1>Hello</h1>`)
		})
	})

	t.Run("Rendering H2", func(t *testing.T) {
		s := subjectAsString(H(2, Text("Hello")))

		t.Run(`it renders <h2> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h2>Hello</h2>`)
		})
	})

	t.Run("Rendering H3", func(t *testing.T) {
		s := subjectAsString(H(3, Text("Hello")))

		t.Run(`it renders <h3> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h3>Hello</h3>`)
		})
	})

	t.Run("Rendering H4", func(t *testing.T) {
		s := subjectAsString(H(4, Text("Hello")))

		t.Run(`it renders <h4> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h4>Hello</h4>`)
		})
	})

	t.Run("Rendering H5", func(t *testing.T) {
		s := subjectAsString(H(5, Text("Hello")))

		t.Run(`it renders <h5> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h5>Hello</h5>`)
		})
	})

	t.Run("Rendering H6", func(t *testing.T) {
		s := subjectAsString(H(6, Text("Hello")))

		t.Run(`it renders <h6> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h6>Hello</h6>`)
		})
	})

	// t.Run("Rendering H{1, LinkView{}}", func(t *testing.T) {
	// 	s := subjectAsString(H{1, LinkView{To: "https://example.org/", Text: "click me"}})

	// 	t.Run(`it renders <h1> with text Hello`, func(t *testing.T) {
	// 		assert.Equal(t, s, `<h1><a href="https://example.org/">click me</a></h1>`)
	// 	})
	// })

	t.Run(`Rendering Button(Text("Click me"))`, func(t *testing.T) {
		s := subjectAsString(Button(Text("Click me")))

		t.Run(`it renders <button> with text "Click me"`, func(t *testing.T) {
			assert.Equal(t, s, `<button type="button">Click me</button>`)
		})
	})

	t.Run(`Rendering Button with Submit`, func(t *testing.T) {
		s := subjectAsString(Button(Text("Click me")).Submit())

		t.Run(`it renders <button> with text "Click me"`, func(t *testing.T) {
			assert.Equal(t, s, `<button type="submit">Click me</button>`)
		})
	})

	t.Run(`Rendering Div with two buttons`, func(t *testing.T) {
		s := subjectAsString(Div(Button(Text("First")), Button(Text("Second"))))

		t.Run(`it renders <div> with two buttons`, func(t *testing.T) {
			assert.Equal(t, s, `<div><button type="button">First</button><button type="button">Second</button></div>`)
		})
	})

	t.Run(`Rendering Header with two buttons`, func(t *testing.T) {
		s := subjectAsString(Header(Button(Text("First")), Button(Text("Second"))))

		t.Run(`it renders <header> with two buttons`, func(t *testing.T) {
			assert.Equal(t, s, `<header><button type="button">First</button><button type="button">Second</button></header>`)
		})
	})
}

type CustomView struct{}

func (v CustomView) Body() View {
	return H(5, Text("Hello"))
}

type CustomViewNested struct{}

func (v CustomViewNested) Body() View {
	return CustomView{}
}

func TestClassNames(t *testing.T) {
	t.Run("Class()", func(t *testing.T) {
		// subject := new(ClassNames)
		var subject ClassNames

		t.Run(`when adding two classes`, func(t *testing.T) {
			after := subject.Class("first", "second")

			// expected := &ClassNames{"first", "second"}

			t.Run(`it has a length of 2`, func(t *testing.T) {
				assert.Equal(t, len(after), 2)
			})

			// t.Run(`it is a slice of two strings`, func(t *testing.T) {
			// 	assert.Equal(t, *subject, []string{"first", "second"})
			// })
		})
	})
}

func TestCustomViews(t *testing.T) {
	t.Run("Render CustomView", func(t *testing.T) {
		s := subjectAsString(CustomView{})

		t.Run(`it renders <h5> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h5>Hello</h5>`)
		})
	})

	t.Run("Render CustomViewNested", func(t *testing.T) {
		s := subjectAsString(CustomViewNested{})

		t.Run(`it renders <h5> with text Hello`, func(t *testing.T) {
			assert.Equal(t, s, `<h5>Hello</h5>`)
		})
	})
}

func TestViewClassNames(t *testing.T) {
	t.Run("Class func can chain", func(t *testing.T) {
		s := subjectAsString(Div().Class("first").Class("second"))

		t.Run(`it renders <div> with class "first second"`, func(t *testing.T) {
			assert.Equal(t, s, `<div class="first second"></div>`)
		})
	})
}

func TestViewAriaAttributes(t *testing.T) {
	t.Run("Adding aria-label", func(t *testing.T) {
		s := subjectAsString(Div(AriaAttr("label", "descriptive label")))

		t.Run(`it renders <div> with attribute "aria-label"`, func(t *testing.T) {
			assert.Equal(t, s, `<div aria-label="descriptive label"></div>`)
		})
	})
}

func TestViewDataAttributes(t *testing.T) {
	t.Run("Adding data-custom", func(t *testing.T) {
		s := subjectAsString(Div(DataAttr("custom", "some value")))

		t.Run(`it renders <div> with attribute "data-custom"`, func(t *testing.T) {
			assert.Equal(t, s, `<div data-custom="some value"></div>`)
		})
	})
}

func TestViewClassNameView(t *testing.T) {
	t.Run("Adding two classes", func(t *testing.T) {
		s := subjectAsString(Div(ClassName("first"), ClassName("second")))

		t.Run(`it renders <div> with class "first"`, func(t *testing.T) {
			assert.Equal(t, s, `<div class="first second"></div>`)
		})
	})
}

func TestViewNilChild(t *testing.T) {
	t.Run("Div with text and nil child", func(t *testing.T) {
		s := subjectAsString(Div(Text("first"), nil, Text("second")))

		t.Run(`it renders <div> with just the text`, func(t *testing.T) {
			assert.Equal(t, s, `<div>firstsecond</div>`)
		})
	})
}

// func TestModifiedViews(t *testing.T) {
// 	t.Run("Class names", func(t *testing.T) {
// 		s := subjectAsString(Modify(H{5, Text("Hello")}).Class("first").Class("second"))

// 		t.Run(`it renders <h5> with class "first second"`, func(t *testing.T) {
// 			assert.Equal(t, s, `<h5 class="first second">Hello</h5>`)
// 		})
// 	})
// }

var result *bytes.Buffer

// func BenchmarkLink(b *testing.B) {
// 	buf := new(bytes.Buffer)

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, LinkView{To: "https://example.org/", Text: "Hello"})
// 	}

// 	result = buf
// }

func BenchmarkText(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Text("hello"))
	}

	result = buf
}

func BenchmarkHeader(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Header())
	}

	result = buf
}

func BenchmarkDiv(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div())
	}

	result = buf
}

func BenchmarkDivWithClasses1(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Class("first"))
	}

	result = buf
}

func BenchmarkDivWithClasses2Together(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Class("first", "second"))
	}

	result = buf
}

func BenchmarkDivWithClasses2(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Class("first").Class("second"))
	}

	result = buf
}

func BenchmarkDivWithChildClassNames(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div(ClassName("first"), ClassName("second")))
	}

	result = buf
}

func BenchmarkButton(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Button(Text("Click me")))
	}

	result = buf
}

func BenchmarkButtonSubmit(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Button(Text("Click me")).Submit())
	}

	result = buf
}

// func BenchmarkH(b *testing.B) {
// 	buf := new(bytes.Buffer)

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, H{1, Text("Click me")})
// 	}

// 	result = buf
// }

// func BenchmarkPremadeH(b *testing.B) {
// 	buf := new(bytes.Buffer)
// 	view := H{1, Text("Click me")}

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, view)
// 	}

// 	result = buf
// }

// func BenchmarkHB(b *testing.B) {
// 	buf := new(bytes.Buffer)

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, HB{1, Text("Click me")})
// 	}

// 	result = buf
// }

// func BenchmarkHeading(b *testing.B) {
// 	buf := new(bytes.Buffer)

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, Heading(1, Text("Click me")))
// 	}

// 	result = buf
// }

func BenchmarkH1(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, H(1, Text("Click me")))
	}

	result = buf
}

// func BenchmarkPremadeHB(b *testing.B) {
// 	buf := new(bytes.Buffer)
// 	view := HB{1, Text("Click me")}

// 	for n := 0; n < b.N; n++ {
// 		buf.Reset()
// 		Render(buf, view)
// 	}

// 	result = buf
// }
