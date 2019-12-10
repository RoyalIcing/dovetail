package dovetail

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestTailwind(t *testing.T) {
	t.Run("Style with padding and text styles", func(t *testing.T) {
		s := subjectAsString(Div().Tailwind(Pt8, Pb8, Text2XL, FontBold))

		t.Run(`it renders <div> with class "pt-8 pb-8 text-2xl font-bold"`, func(t *testing.T) {
			assert.Equal(t, s, `<div class="pt-8 pb-8 text-2xl font-bold"></div>`)
		})
	})

	t.Run("Style with md styles", func(t *testing.T) {
		s := subjectAsString(Div().Tailwind(Pt8, Pb8, TextBase, FontBold).Md(Pt1, Text2XL))

		t.Run(`it renders <div> with class "pt-8 pb-8 text-base font-bold md:pt-1 md:text-2xl"`, func(t *testing.T) {
			assert.Equal(t, s, `<div class="pt-8 pb-8 text-base font-bold md:pt-1 md:text-2xl"></div>`)
		})
	})
}

func BenchmarkTailwindJustDiv(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div())
	}

	result = buf
}

func BenchmarkTailwind0Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Tailwind())
	}

	result = buf
}

func BenchmarkTailwind2Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Tailwind(Text2XL, FontBold))
	}

	result = buf
}

func BenchmarkTailwind4Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Tailwind(Pt8, Pb8, Text2XL, FontBold))
	}

	result = buf
}

func BenchmarkTailwind8Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().Tailwind(Pt8, Pb8, Mb8, Text2XL, FontBold, TextBlue300, BgBlue800, RoundedFull))
	}

	result = buf
}

func BenchmarkTailwindAddClasses2Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().AddClasses(TailwindToClass(Text2XL, FontBold)))
	}

	result = buf
}

func BenchmarkTailwindAddClasses4Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().AddClasses(TailwindToClass(Pt8, Pb8, Text2XL, FontBold)))
	}

	result = buf
}

func BenchmarkTailwindAddClasses8Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().AddClasses(TailwindToClass(Pt8, Pb8, Mb8, Text2XL, FontBold, TextBlue300, BgBlue800, RoundedFull)))
	}

	result = buf
}

func BenchmarkTailwindDivWithClasses8Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, DivWithClasses(TailwindToClass(Pt8, Pb8, Mb8, Text2XL, FontBold, TextBlue300, BgBlue800, RoundedFull)))
	}

	result = buf
}

func BenchmarkTailwindChangeClasses2Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().ChangeClasses(TailwindChanger(Text2XL, FontBold)))
	}

	result = buf
}

func BenchmarkTailwindChangeClasses4Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().ChangeClasses(TailwindChanger(Pt8, Pb8, Text2XL, FontBold)))
	}

	result = buf
}

func BenchmarkTailwindChangeClasses8Classes(b *testing.B) {
	buf := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		buf.Reset()
		Render(buf, Div().ChangeClasses(TailwindChanger(Pt8, Pb8, Mb8, Text2XL, FontBold, TextBlue300, BgBlue800, RoundedFull)))
	}

	result = buf
}
