package structsStrings

type Struct10 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10                                                                string
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
		Field4: "Field4-Value",
		Field5: "Field5-Value",
		Field6: "Field6-Value",
		Field7: "Field7-Value",
		Field8: "Field8-Value",
		Field9: "Field9-Value",
		Field10: "Field10-Value",
	}
}