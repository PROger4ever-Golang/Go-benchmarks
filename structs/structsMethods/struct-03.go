package structsMethods

type Struct03 struct {
	Field01, Field02, Field03 int
}
func (s Struct03) ObjectPrimitiveMethod() int {
	return s.Field03
}
func (sp *Struct03) PointerPrimitiveMethod() int {
	return sp.Field03
}
func (s Struct03) ObjectParameterMethod(p Struct03) int {
	return p.Field03
}
func (sp *Struct03) PointerParameterMethod(p *Struct03) int {
	return p.Field03
}
func (s Struct03) ObjectReturnMethod(p Struct03) Struct03 {
	return p
}
func (sp *Struct03) PointerReturnMethod(p *Struct03) *Struct03 {
	return p
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field01: 1,
		Field02: 2,
		Field03: 3,
	}
}