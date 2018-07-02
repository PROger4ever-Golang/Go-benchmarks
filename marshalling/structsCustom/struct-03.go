package structsCustom

import (
	"bytes"
)

type Struct03 struct {
	Field01, Field02, Field03 string
}
const jsonSkeletonLength03 = jsonStartLength + shortFieldLength*2 + jsonEndLength
func (es *Struct03) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength03+ len(es.Field01) + len(es.Field02) + len(es.Field03))
	buf = append(buf, `{"Field01":"`...); buf = append(buf, es.Field01...) //escape field values?
	buf = append(buf, `","Field02":"`...); buf = append(buf, es.Field02...)
	buf = append(buf, `","Field03":"`...); buf = append(buf, es.Field03...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct03) UnmarshalJSON(data []byte) error {
	l := len(data) - 2 // `"}` at the end
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field01 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field02 = string(data[s:e])
	s = e + shortFieldLength; es.Field03 = string(data[s:l])
	return nil
}

func NewSampleStruct03() *Struct03 {
	return &Struct03{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
	}
}