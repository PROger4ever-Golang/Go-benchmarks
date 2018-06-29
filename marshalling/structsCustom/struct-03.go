package structsCustom

import (
	"bytes"
)

type Struct03 struct {
	Field1, Field2, Field3 string
}
const jsonSkeletonLength03 = jsonStartLength + shortFieldLength*2 + jsonEndLength
func (es *Struct03) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength03+ len(es.Field1) + len(es.Field2) + len(es.Field3))
	buf = append(buf, `{"Field1":"`...); buf = append(buf, es.Field1...) //escape field values?
	buf = append(buf, `","Field2":"`...); buf = append(buf, es.Field2...)
	buf = append(buf, `","Field3":"`...); buf = append(buf, es.Field3...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct03) UnmarshalJSON(data []byte) error {
	l := len(data) - 2 // `"}` at the end
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field1 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field2 = string(data[s:e])
	s = e + shortFieldLength; es.Field3 = string(data[s:l])
	return nil
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
	}
}