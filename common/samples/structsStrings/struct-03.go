package structsStrings

type Struct03 struct {
	Field01, Field02, Field03 string
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
	}
}