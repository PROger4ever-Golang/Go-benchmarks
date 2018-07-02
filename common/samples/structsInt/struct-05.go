package structsInt

type Struct05 struct {
	Field01, Field02, Field03, Field04, Field05 int
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field01: 1,
		Field02: 2,
		Field03: 3,
		Field04: 4,
		Field05: 5,
	}
}