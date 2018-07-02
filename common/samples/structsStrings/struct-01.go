package structsStrings

type Struct01 struct {
	Field01 string
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field01: "Field01-Value",
	}
}