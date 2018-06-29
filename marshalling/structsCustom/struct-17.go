package structsCustom

import (
	"bytes"
)

type Struct17 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 string
}
const jsonSkeletonLength17 = jsonStartLength + shortFieldLength*8 + longFieldLength*8 + jsonEndLength
func (es *Struct17) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength17 + len(es.Field1) + len(es.Field2) + len(es.Field3) + len(es.Field4) +
		len(es.Field5) + len(es.Field6) + len(es.Field7) + len(es.Field8) + len(es.Field9) + len(es.Field10) +
		len(es.Field11) + len(es.Field12) + len(es.Field13) + len(es.Field14) + len(es.Field15) + len(es.Field16) +
		len(es.Field17))
	buf = append(buf, `{"Field1":"`...); buf = append(buf, es.Field1...) //escape field values?
	buf = append(buf, `","Field2":"`...); buf = append(buf, es.Field2...)
	buf = append(buf, `","Field3":"`...); buf = append(buf, es.Field3...)
	buf = append(buf, `","Field4":"`...); buf = append(buf, es.Field4...)
	buf = append(buf, `","Field5":"`...); buf = append(buf, es.Field5...)
	buf = append(buf, `","Field6":"`...); buf = append(buf, es.Field6...)
	buf = append(buf, `","Field7":"`...); buf = append(buf, es.Field7...)
	buf = append(buf, `","Field8":"`...); buf = append(buf, es.Field8...)
	buf = append(buf, `","Field9":"`...); buf = append(buf, es.Field9...)
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
	e := bytes.Index(data[jsonStartLength:l], quoteBytes) + jsonStartLength; es.Field1 = string(data[jsonStartLength:e])
	s := e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field2 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field3 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field4 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field5 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field6 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field7 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field8 = string(data[s:e])
	s = e + shortFieldLength; e = bytes.Index(data[s:l], quoteBytes) + s; es.Field9 = string(data[s:e])
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
		Field1: "Field1-Value",
		Field2: "Field2-Value",
		Field3: "Field3-Value",
		Field4: "Field4-Value",
		Field5: "Field5-Value",
		Field6: "Field6-Value",
		Field7: "Field7-Value",
		Field8: "Field8-Value",
		Field9: "Field9-Value",
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