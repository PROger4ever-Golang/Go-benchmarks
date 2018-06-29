package structsCustom

type Struct01 struct {
	Field1 string
}
const jsonSkeletonLength01 = jsonStartLength + jsonEndLength
func (es *Struct01) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength01 + len(es.Field1))
	buf = append(buf, `{"Field1":"`...); buf = append(buf, es.Field1...) //escape field values?
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct01) UnmarshalJSON(data []byte) error {
	es.Field1 = string(data[jsonStartLength : len(data)-2]) //`"}` at the end
	return nil
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field1: "Field1-Value",
	}
}
