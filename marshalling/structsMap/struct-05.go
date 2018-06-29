package structsMap

type Struct05 struct {
	Field1, Field2, Field3, Field4, Field5 string
}
func (es *Struct05) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
		"Field2": es.Field2,
		"Field3": es.Field3,
		"Field4": es.Field4,
		"Field5": es.Field5,
	}
}
func (es *Struct05) UnmarshalMap(theMap map[string]interface{}) *Struct05 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	es.Field4 = theMap["Field4"].(string)
	es.Field5 = theMap["Field5"].(string)
	return es
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
		Field4: "Field4-Value",
		Field5: "Field5-Value",
	}
}