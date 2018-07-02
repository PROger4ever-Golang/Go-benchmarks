package structsInt

type Struct03 struct {
	Field01, Field02, Field03 int
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field01: 1,
		Field02: 2,
		Field03: 3,
	}
}