package structsCustom

import (
	"bytes"
)

type Struct05 struct {
	Field1, Field2, Field3, Field4, Field5 string
}
const jsonSkeletonLength05 = jsonStartLength + shortFieldLength*4 + jsonEndLength
func (es *Struct05) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength05 + len(es.Field1) + len(es.Field2) + len(es.Field3) + len(es.Field4) +
		len(es.Field5))
	buf = append(buf, `{"Field1":"`...); buf = append(buf, es.Field1...) //escape field values?
	buf = append(buf, `","Field2":"`...); buf = append(buf, es.Field2...)
	buf = append(buf, `","Field3":"`...); buf = append(buf, es.Field3...)
	buf = append(buf, `","Field4":"`...); buf = append(buf, es.Field4...)
	buf = append(buf, `","Field5":"`...); buf = append(buf, es.Field5...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct05) UnmarshalJSON(data []byte) error {
	l := len(data)-2 // `"}` at the end
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field1 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field2 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field3 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field4 = string(data[s:e])
	s = e + shortFieldLength; es.Field5 = string(data[s:l])
	return nil
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