package structsMap

type Struct05 struct {
	Field01, Field02, Field03, Field04, Field05 string
}
func (es *Struct05) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field01": es.Field01,
		"Field02": es.Field02,
		"Field03": es.Field03,
		"Field04": es.Field04,
		"Field05": es.Field05,
	}
}
func (es *Struct05) UnmarshalMap(theMap map[string]interface{}) *Struct05 {
	es.Field01 = theMap["Field01"].(string)
	es.Field02 = theMap["Field02"].(string)
	es.Field03 = theMap["Field03"].(string)
	es.Field04 = theMap["Field04"].(string)
	es.Field05 = theMap["Field05"].(string)
	return es
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
	}
}