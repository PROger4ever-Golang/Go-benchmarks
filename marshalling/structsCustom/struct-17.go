package structsCustom

import (
	"bytes"
)

type Struct17 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 string
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 string
}
const jsonSkeletonLength17 = jsonStartLength + shortFieldLength*8 + longFieldLength*8 + jsonEndLength
func (es *Struct17) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength17 + len(es.Field01) + len(es.Field02) + len(es.Field03) + len(es.Field04) +
		len(es.Field05) + len(es.Field06) + len(es.Field07) + len(es.Field08) + len(es.Field09) + len(es.Field10) +
		len(es.Field11) + len(es.Field12) + len(es.Field13) + len(es.Field14) + len(es.Field15) + len(es.Field16) +
		len(es.Field17))
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
	buf = append(buf, `","Field11":"`...); buf = append(buf, es.Field11...)
	buf = append(buf, `","Field12":"`...); buf = append(buf, es.Field12...)
	buf = append(buf, `","Field13":"`...); buf = append(buf, es.Field13...)
	buf = append(buf, `","Field14":"`...); buf = append(buf, es.Field14...)
	buf = append(buf, `","Field15":"`...); buf = append(buf, es.Field15...)
	buf = append(buf, `","Field16":"`...); buf = append(buf, es.Field16...)
	buf = append(buf, `","Field17":"`...); buf = append(buf, es.Field17...)
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct17) UnmarshalJSON(data []byte) error {
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
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field10 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field11 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field12 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field13 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field14 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field15 = string(data[s:e])
	s = e + longFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field16 = string(data[s:e])
	s = e + longFieldLength; es.Field17 = string(data[s:l])
	return nil
}

func NewSampleStruct17() *Struct17 {
	return &Struct17{
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
		Field11: "Field11-Value",
		Field12: "Field12-Value",
		Field13: "Field13-Value",
		Field14: "Field14-Value",
		Field15: "Field15-Value",
		Field16: "Field16-Value",
		Field17: "Field17-Value",
	}
}