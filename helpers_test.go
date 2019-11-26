package main

import (
	"bytes"
)

func subjectAsString(view View) string {
	b := new(bytes.Buffer)
	Render(b, view)
	return b.String()
}
