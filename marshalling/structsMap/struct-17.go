package structsMap

type Struct17 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 string
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 string
}
func (es *Struct17) MarshalMap() map[string]interface{} {
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
		"Field11": es.Field11,
		"Field12": es.Field12,
		"Field13": es.Field13,
		"Field14": es.Field14,
		"Field15": es.Field15,
		"Field16": es.Field16,
		"Field17": es.Field17,
	}
}
func (es *Struct17) UnmarshalMap(theMap map[string]interface{}) *Struct17 {
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
	es.Field11 = theMap["Field11"].(string)
	es.Field12 = theMap["Field12"].(string)
	es.Field13 = theMap["Field13"].(string)
	es.Field14 = theMap["Field14"].(string)
	es.Field15 = theMap["Field15"].(string)
	es.Field16 = theMap["Field16"].(string)
	es.Field17 = theMap["Field17"].(string)
	return es
}

func NewSampleStruct17() *Struct17 {
	return &Struct17{
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
		Field11: "Field11-Value",
		Field12: "Field12-Value",
		Field13: "Field13-Value",
		Field14: "Field14-Value",
		Field15: "Field15-Value",
		Field16: "Field16-Value",
		Field17: "Field17-Value",
	}
}