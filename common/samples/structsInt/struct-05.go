package structsInt

type Struct05 struct {
	Field1, Field2, Field3, Field4, Field5 int
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field1: 1,
		Field2: 2,
		Field3: 3,
		Field4: 4,
		Field5: 5,
	}
}