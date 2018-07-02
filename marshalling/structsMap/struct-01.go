package structsMap

type Struct01 struct {
	Field01 string
}
func (es *Struct01) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field01": es.Field01,
	}
}
func (es *Struct01) UnmarshalMap(theMap map[string]interface{}) *Struct01 {
	es.Field01 = theMap["Field01"].(string)
	return es
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field01: "Field01-Value",
	}
}