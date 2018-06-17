package structsInt

type Struct17 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 int
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 int
}

func NewSampleStruct17() *Struct17 {
	return &Struct17{
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
		Field11: 11,
		Field12: 12,
		Field13: 13,
		Field14: 14,
		Field15: 15,
		Field16: 16,
		Field17: 17,
	}
}