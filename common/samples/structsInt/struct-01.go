package structsInt

type Struct01 struct {
	Field01 int
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field01: 1,
	}
}