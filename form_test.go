package dovetail

import (
	"bytes"
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

	t.Run("Rendering Field with text input named description", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Write description", Textbox("description")))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Write description</span><input type="text" name="description"/></label>`)
		})
	})

	t.Run("Rendering Field with text input with class", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Write description", Textbox("description").DefaultValue("some value").Use(Class("some classes"))))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Write description</span><input type="text" name="description" value="some value" class="some classes"/></label>`)
		})
	})

	t.Run("Rendering Field with text input with 3 rows", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Favorite number", Textbox("description").Rows(3).DefaultValue("some value")))

		t.Run(`it renders <label> with child <span> with label text and <textarea> with name and rows attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Favorite number</span><textarea name="description" rows="3">some value</textarea></label>`)
		})
	})

	t.Run("Rendering Field with file input named image", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Add picture", FileInput("image")))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Add picture</span><input type="file" name="image"/></label>`)
		})
	})

	t.Run("Rendering Field with type of number", func(t *testing.T) {
		s := subjectAsString(FieldLabelled("Favorite number", NumberInput("fave_number")))

		t.Run(`it renders <label> with child <span> with label text and <input> with type and name attibutes`, func(t *testing.T) {
			assert.Equal(t, s, `<label><span>Favorite number</span><input type="number" name="fave_number"/></label>`)
		})
	})

	t.Run("Rendering Form with child field and submit button", func(t *testing.T) {
		s := subjectAsString(
			FormTo("/things", Multipart).With(
				Class("border"),
				FieldLabelled("Add picture",
					FileInput("image"),
					Class("block"),
				),
				ButtonOld(Text("Upload")).Submit(),
				// ButtonOld("Upload", Submit),
				// ButtonOld(Submit, Text("Upload")),
				// SubmitButton("Upload"),
			),
		)

		t.Run(`it renders <form> with labelled input and submit button`, func(t *testing.T) {
			assert.Equal(t, s, strings.Replace(`
<form method="post" action="/things" enctype="multipart/form-data" class="border">
<label class="block"><span>Add picture</span><input type="file" name="image"/></label>
<button type="submit">Upload</button>
</form>
`, "\n", "", -1))
		})
	})
}

func BenchmarkFormField(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, FieldLabelled("Add picture", FileInput("image")))
	}

	result = buf
}

func BenchmarkFormFieldWithClass(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, FieldLabelled("Add picture",
			FileInput("image"),
			Class("block"),
		))
	}

	result = buf
}

func BenchmarkFormFieldWithClass2(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, FieldLabelled("Add picture",
			FileInput("image"),
		).Class("block"))
	}

	result = buf
}
