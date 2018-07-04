package structsMethods

type Struct17 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 int
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 int
}
func (s Struct17) ObjectPrimitiveMethod() int {
	return s.Field17
}
func (sp *Struct17) PointerPrimitiveMethod() int {
	return sp.Field17
}
func (s Struct17) ObjectParameterMethod(p Struct17) int {
	return p.Field17
}
func (sp *Struct17) PointerParameterMethod(p *Struct17) int {
	return p.Field17
}
func (s Struct17) ObjectReturnMethod(p Struct17) Struct17 {
	return p
}
func (sp *Struct17) PointerReturnMethod(p *Struct17) *Struct17 {
	return p
}

func NewSampleStruct17() *Struct17 {
	return &Struct17{
		Field01: 1,
		Field02: 2,
		Field03: 3,
		Field04: 4,
		Field05: 5,
		Field06: 6,
		Field07: 7,
		Field08: 8,
		Field09: 9,
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