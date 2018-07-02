package structsStrings

type Struct10 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 string
	Field10                                                                string
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
		Field06: "Field06-Value",
		Field07: "Field07-Value",
		Field08: "Field08-Value",
		Field09: "Field09-Value",
		Field10: "Field10-Value",
	}
}