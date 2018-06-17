package structsStrings

type Struct05 struct {
	Field1, Field2, Field3, Field4, Field5 string
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
		Field4: "Field4-Value",
		Field5: "Field5-Value",
	}
}