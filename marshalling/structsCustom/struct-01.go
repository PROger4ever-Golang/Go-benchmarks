package structsCustom

type Struct01 struct {
	Field01 string
}
const jsonSkeletonLength01 = jsonStartLength + jsonEndLength
func (es *Struct01) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength01 + len(es.Field01))
	buf = append(buf, `{"Field01":"`...); buf = append(buf, es.Field01...) //escape field values?
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct01) UnmarshalJSON(data []byte) error {
	es.Field01 = string(data[jsonStartLength : len(data)-2]) //`"}` at the end
	return nil
}

func NewSampleStruct01() *Struct01 {
	return &Struct01{
		Field01: "Field01-Value",
	}
}
