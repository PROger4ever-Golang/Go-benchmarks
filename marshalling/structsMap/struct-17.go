package structsMap

type Struct17 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 string
}
func (es *Struct17) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field1":  es.Field1,
		"Field2":  es.Field2,
		"Field3":  es.Field3,
		"Field4":  es.Field4,
		"Field5":  es.Field5,
		"Field6":  es.Field6,
		"Field7":  es.Field7,
		"Field8":  es.Field8,
		"Field9":  es.Field9,
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
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	es.Field4 = theMap["Field4"].(string)
	es.Field5 = theMap["Field5"].(string)
	es.Field6 = theMap["Field6"].(string)
	es.Field7 = theMap["Field7"].(string)
	es.Field8 = theMap["Field8"].(string)
	es.Field9 = theMap["Field9"].(string)
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
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
		Field4: "Field4-Value",
		Field5: "Field5-Value",
		Field6: "Field6-Value",
		Field7: "Field7-Value",
		Field8: "Field8-Value",
		Field9: "Field9-Value",
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