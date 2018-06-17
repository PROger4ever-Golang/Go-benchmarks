package structsMap

type Struct03 struct {
	Field1, Field2, Field3 string
}
func (es *Struct03) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
		"Field2": es.Field2,
		"Field3": es.Field3,
	}
}
func (es *Struct03) UnmarshalMap(theMap map[string]interface{}) *Struct03 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	return es
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
	}
}