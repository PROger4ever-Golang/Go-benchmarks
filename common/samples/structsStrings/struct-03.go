package structsStrings

type Struct03 struct {
	Field1, Field2, Field3 string
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
	}
}