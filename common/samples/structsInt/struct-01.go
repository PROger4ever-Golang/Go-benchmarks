package structsInt

type Struct01 struct {
	Field1 int
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field1: 1,
	}
}