package structsMap

type Struct10 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10                                                                string
}
func (es *Struct10) MarshalMap() map[string]interface{} {
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
	}
}
func (es *Struct10) UnmarshalMap(theMap map[string]interface{}) *Struct10 {
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
	return es
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
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
	}
}