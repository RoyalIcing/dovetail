package main

// TailwindClassName is a subset of strings allowed as Tailwind class names
type TailwindClassName string

const (
	// Pt1 padding top of 1
	Pt1 TailwindClassName = "pt-1"
	// Pt2 padding top of 2
	Pt2 TailwindClassName = "pt-2"
	// Pt4 padding top of 2
	Pt4 TailwindClassName = "pt-4"
	// Pt8 padding top of 8
	Pt8 TailwindClassName = "pt-8"
	// Pb1 padding bottom of 1
	Pb1 TailwindClassName = "pb-1"
	// Pb2 padding bottom of 2
	Pb2 TailwindClassName = "pb-2"
	// Pb4 padding bottom of 4
	Pb4 TailwindClassName = "pb-4"
	// Pb8 padding bottom of 8
	Pb8 TailwindClassName = "pb-8"
	// Px3 padding left and right of 3
	Px3 TailwindClassName = "px-3"
	// Py1 padding top and bottom of 1
	Py1 TailwindClassName = "py-1"

	// Mb8 margin bottom of 8
	Mb8 TailwindClassName = "mb-8"

	// TextXS text of small size
	TextXS TailwindClassName = "text-xs"
	// TextSM text of small size
	TextSM TailwindClassName = "text-sm"
	// TextBase text of size 1rem
	TextBase TailwindClassName = "text-base"
	// TextXL text of XL size
	TextXL TailwindClassName = "text-xl"
	// Text2XL text of 2XL size
	Text2XL TailwindClassName = "text-2xl"

	// TextBlue300 blue text light 300
	TextBlue300 TailwindClassName = "text-blue-300"

	// BgBlue700 blue background dark 700
	BgBlue700 TailwindClassName = "bg-blue-700"
	// BgBlue800 blue background dark 800
	BgBlue800 TailwindClassName = "bg-blue-800"

	// FontBold bold font weight
	FontBold TailwindClassName = "font-bold"

	// RoundedFull rounded corners in a pill shape
	RoundedFull TailwindClassName = "rounded-full"
)

func Tailwind(additions ...TailwindClassName) ClassNames {
	classNames := make(ClassNames, 0, len(additions))
	for _, addition := range additions {
		classNames = append(classNames, string(addition))
	}
	return classNames
}

func (classNames ClassNames) Tailwind(additions ...TailwindClassName) ClassNames {
	for _, addition := range additions {
		classNames = append(classNames, string(addition))
	}
	return classNames
}

func (classNames ClassNames) Md(additions ...TailwindClassName) ClassNames {
	for _, addition := range additions {
		classNames = append(classNames, "md:"+string(addition))
	}
	return classNames
}

type ClassNamesChanger func(classNames ClassNames) ClassNames

func TailwindChanger(additions ...TailwindClassName) ClassNamesChanger {
	return func(classNames ClassNames) ClassNames {
		for _, addition := range additions {
			classNames = append(classNames, string(addition))
		}
		return classNames
	}
}
func (changer ClassNamesChanger) Tailwind(additions ...TailwindClassName) ClassNamesChanger {
	return func(classNames ClassNames) ClassNames {
		classNames = changer(classNames)

		for _, addition := range additions {
			classNames = append(classNames, string(addition))
		}
		return classNames
	}
}

func (changer ClassNamesChanger) Md(additions ...TailwindClassName) ClassNamesChanger {
	return func(classNames ClassNames) ClassNames {
		classNames = changer(classNames)

		for _, addition := range additions {
			classNames = append(classNames, "md:"+string(addition))
		}
		return classNames
	}
}

func Hover(baseName TailwindClassName) TailwindClassName {
	return TailwindClassName("hover:" + string(baseName))
}

func (basic HTMLElementView) Tailwind(additions ...TailwindClassName) HTMLElementView {
	basic.classNames = basic.classNames.Tailwind(additions...)
	return basic

	// classNameStrings := make([]string, 0, len(additions))
	// for _, addition := range additions {
	// 	classNameStrings = append(classNameStrings, string(addition))
	// }
	// return basic.Class(classNameStrings...)

}

func (basic HTMLElementView) Md(classNames ...TailwindClassName) HTMLElementView {
	classNameStrings := make([]string, 0, len(classNames))
	for _, className := range classNames {
		classNameStrings = append(classNameStrings, "md:"+string(className))
	}
	return basic.Class(classNameStrings...)
	// mutable := &basic
	// mutable.Class(classNameStrings...)
	// return *mutable
}
