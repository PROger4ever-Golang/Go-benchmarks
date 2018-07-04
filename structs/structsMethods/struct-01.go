package structsMethods

type Struct01 struct {
	Field01 int
}
func (s Struct01) ObjectPrimitiveMethod() int {
	return s.Field01
}
func (sp *Struct01) PointerPrimitiveMethod() int {
	return sp.Field01
}
func (s Struct01) ObjectParameterMethod(p Struct01) int {
	return p.Field01
}
func (sp *Struct01) PointerParameterMethod(p *Struct01) int {
	return p.Field01
}
func (s Struct01) ObjectReturnMethod(p Struct01) Struct01 {
	return p
}
func (sp *Struct01) PointerReturnMethod(p *Struct01) *Struct01 {
	return p
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field01: 1,
	}
}