package structsStrings

type Struct01 struct {
	Field1 string
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field1: "Field1-Value",
	}
}