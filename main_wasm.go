package main

import (
	"bytes"

	"syscall/js"
)

func main() {
	view := Header(
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
	)

	b := new(bytes.Buffer)
	Render(b, view)
	rawHTML := b.String()

	js.Global().Call("updateBody", rawHTML)
}
