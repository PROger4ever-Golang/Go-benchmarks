package structsInt

type Struct03 struct {
	Field1, Field2, Field3 int
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field1: 1,
		Field2: 2,
		Field3: 3,
	}
}