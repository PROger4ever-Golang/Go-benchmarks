package structsMap

type Struct03 struct {
	Field01, Field02, Field03 string
}
func (es *Struct03) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field01": es.Field01,
		"Field02": es.Field02,
		"Field03": es.Field03,
	}
}
func (es *Struct03) UnmarshalMap(theMap map[string]interface{}) *Struct03 {
	es.Field01 = theMap["Field01"].(string)
	es.Field02 = theMap["Field02"].(string)
	es.Field03 = theMap["Field03"].(string)
	return es
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
	}
}