package structsCustom

import (
	"bytes"
)

type Struct10 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 string
	Field10                                                                string
}
const jsonSkeletonLength10 = jsonStartLength + shortFieldLength*8 + longFieldLength*1 + jsonEndLength
func (es *Struct10) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength10 + len(es.Field01) + len(es.Field02) + len(es.Field03) + len(es.Field04) +
		len(es.Field05) + len(es.Field06) + len(es.Field07) + len(es.Field08) + len(es.Field09) + len(es.Field10))
	buf = append(buf, `{"Field01":"`...); buf = append(buf, es.Field01...) //escape field values?
	buf = append(buf, `","Field02":"`...); buf = append(buf, es.Field02...)
	buf = append(buf, `","Field03":"`...); buf = append(buf, es.Field03...)
	buf = append(buf, `","Field04":"`...); buf = append(buf, es.Field04...)
	buf = append(buf, `","Field05":"`...); buf = append(buf, es.Field05...)
	buf = append(buf, `","Field06":"`...); buf = append(buf, es.Field06...)
	buf = append(buf, `","Field07":"`...); buf = append(buf, es.Field07...)
	buf = append(buf, `","Field08":"`...); buf = append(buf, es.Field08...)
	buf = append(buf, `","Field09":"`...); buf = append(buf, es.Field09...)
	buf = append(buf, `","Field10":"`...); buf = append(buf, es.Field10...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct10) UnmarshalJSON(data []byte) error {
	l := len(data)-2 // `"}` at the end
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field01 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field02 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field03 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field04 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field05 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field06 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field07 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field08 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field09 = string(data[s:e])
	s = e + longFieldLength; es.Field10 = string(data[s:l])
	return nil
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field01: "Field01-Value",
		Field02: "Field02-Value",
		Field03: "Field03-Value",
		Field04: "Field04-Value",
		Field05: "Field05-Value",
		Field06: "Field06-Value",
		Field07: "Field07-Value",
		Field08: "Field08-Value",
		Field09: "Field09-Value",
		Field10: "Field10-Value",
	}
}