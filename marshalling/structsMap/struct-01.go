package structsMap

type Struct01 struct {
	Field1 string
}
func (es *Struct01) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
	}
}
func (es *Struct01) UnmarshalMap(theMap map[string]interface{}) *Struct01 {
	es.Field1 = theMap["Field1"].(string)
	return es
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field1: "Field1-Value",
	}
}