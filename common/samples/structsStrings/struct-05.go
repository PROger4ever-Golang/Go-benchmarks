package structsStrings

type Struct05 struct {
	Field01, Field02, Field03, Field04, Field05 string
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
	}
}