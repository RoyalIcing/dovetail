package main

import (
	"bytes"
)

func subjectAsString(view HTMLView) string {
	b := new(bytes.Buffer)
	Render(b, view)
	return b.String()
}
