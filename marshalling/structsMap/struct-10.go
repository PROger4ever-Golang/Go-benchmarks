package structsMap

type Struct10 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 string
	Field10                                                                string
}
func (es *Struct10) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field01":  es.Field01,
		"Field02":  es.Field02,
		"Field03":  es.Field03,
		"Field04":  es.Field04,
		"Field05":  es.Field05,
		"Field06":  es.Field06,
		"Field07":  es.Field07,
		"Field08":  es.Field08,
		"Field09":  es.Field09,
		"Field10": es.Field10,
	}
}
func (es *Struct10) UnmarshalMap(theMap map[string]interface{}) *Struct10 {
	es.Field01 = theMap["Field01"].(string)
	es.Field02 = theMap["Field02"].(string)
	es.Field03 = theMap["Field03"].(string)
	es.Field04 = theMap["Field04"].(string)
	es.Field05 = theMap["Field05"].(string)
	es.Field06 = theMap["Field06"].(string)
	es.Field07 = theMap["Field07"].(string)
	es.Field08 = theMap["Field08"].(string)
	es.Field09 = theMap["Field09"].(string)
	es.Field10 = theMap["Field10"].(string)
	return es
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
		Field06: "Field06-Value",
		Field07: "Field07-Value",
		Field08: "Field08-Value",
		Field09: "Field09-Value",
		Field10: "Field10-Value",
	}
}