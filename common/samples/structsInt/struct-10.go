package structsInt

type Struct10 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 int
	Field10                                                                int
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field1: 1,
		Field2: 2,
		Field3: 3,
		Field4: 4,
		Field5: 5,
		Field6: 6,
		Field7: 7,
		Field8: 8,
		Field9: 9,
		Field10: 10,
	}
}