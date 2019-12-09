package dovetail

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestForm(t *testing.T) {
	t.Run("Rendering empty Form", func(t *testing.T) {
		s := subjectAsString(FormTo("/things"))

		t.Run(`it renders <form> with method = post and action = /things`, func(t *testing.T) {
			assert.Equal(t, s, `<form method="post" action="/things"></form>`)
		})
	})

	t.Run("Rendering empty Form with Multipart", func(t *testing.T) {
		s := subjectAsString(FormTo("/things", Multipart))

		t.Run(`it renders <form> with method = post and action = /things and enctype = multipart/form-data`, func(t *testing.T) {
			assert.Equal(t, s, `<form method="post" action="/things" enctype="multipart/form-data"></form>`)
		})
	})

	t.Run("Rendering empty Form with Multipart", func(t *testing.T) {
		s := subjectAsString(FormTo("/things").Multipart())

		t.Run(`it renders <form> with method = post and action = /things and enctype = multipart/form-data`, func(t *testing.T) {
			assert.Equal(t, s, `<form method="post" action="/things" enctype="multipart/form-data"></form>`)
		})
	})

	t.Run("Rendering Field named image with type of file", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Add picture", InputNamed("image", InputType("file"))))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Add picture</span><input type="file" name="image"/></label>`)
		})
	})

	t.Run("Rendering Field named image with type of file", func(t *testing.T) {
		s := subjectAsString(
			FieldNamed("image").
				Label(Text("Add picture")).
				Type("file"),
		)

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Add picture</span><input type="file" name="image"/></label>`)
		})
	})

	t.Run("Rendering Field with type of number", func(t *testing.T) {
		s := subjectAsString(FieldNamed("fave_number").Type("number").Label(Text("Favorite number")))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Favorite number</span><input type="number" name="fave_number"/></label>`)
		})
	})

	t.Run("Rendering Form with child field and submit button", func(t *testing.T) {
		s := subjectAsString(
			FormTo("/things", Multipart).With(
				FieldLabelled("Add picture", InputNamed("image", FileInput)),
				Button(Text("Upload")).Submit(),
				// Button("Upload", Submit),
				// Button(Submit, Text("Upload")),
				// SubmitButton("Upload"),
			),
		)

		t.Run(`it renders <form> with labelled input and submit button`, func(t *testing.T) {
			assert.Equal(t, s, strings.Replace(`
<form method="post" action="/things" enctype="multipart/form-data">
<label><span>Add picture</span><input type="file" name="image"/></label>
<button type="submit">Upload</button>
</form>
`, "\n", "", -1))
		})
	})
}
