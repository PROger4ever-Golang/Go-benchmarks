package structsCustom

import (
	"bytes"
)

type Struct05 struct {
	Field01, Field02, Field03, Field04, Field05 string
}
const jsonSkeletonLength05 = jsonStartLength + shortFieldLength*4 + jsonEndLength
func (es *Struct05) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength05 + len(es.Field01) + len(es.Field02) + len(es.Field03) + len(es.Field04) +
		len(es.Field05))
	buf = append(buf, `{"Field01":"`...); buf = append(buf, es.Field01...) //escape field values?
	buf = append(buf, `","Field02":"`...); buf = append(buf, es.Field02...)
	buf = append(buf, `","Field03":"`...); buf = append(buf, es.Field03...)
	buf = append(buf, `","Field04":"`...); buf = append(buf, es.Field04...)
	buf = append(buf, `","Field05":"`...); buf = append(buf, es.Field05...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct05) UnmarshalJSON(data []byte) error {
	l := len(data)-2 // `"}` at the end
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field01 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field02 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field03 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field04 = string(data[s:e])
	s = e + shortFieldLength; es.Field05 = string(data[s:l])
	return nil
}

func NewSampleStruct05() *Struct05 {
	return &Struct05{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
	}
}