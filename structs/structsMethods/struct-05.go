package structsMethods

type Struct05 struct {
	Field01, Field02, Field03, Field04, Field05 int
}
func (s Struct05) ObjectPrimitiveMethod() int {
	return s.Field05
}
func (sp *Struct05) PointerPrimitiveMethod() int {
	return sp.Field05
}
func (s Struct05) ObjectParameterMethod(p Struct05) int {
	return p.Field05
}
func (sp *Struct05) PointerParameterMethod(p *Struct05) int {
	return p.Field05
}
func (s Struct05) ObjectReturnMethod(p Struct05) Struct05 {
	return p
}
func (sp *Struct05) PointerReturnMethod(p *Struct05) *Struct05 {
	return p
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