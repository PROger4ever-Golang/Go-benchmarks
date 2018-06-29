package structsCustom

import (
	"bytes"
)

type Struct10 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10                                                                string
}
const jsonSkeletonLength10 = jsonStartLength + shortFieldLength*8 + longFieldLength*1 + jsonEndLength
func (es *Struct10) MarshalJSON() ([]byte, error) {
	var buf = make([]byte, 0, jsonSkeletonLength10 + len(es.Field1) + len(es.Field2) + len(es.Field3) + len(es.Field4) +
		len(es.Field5) + len(es.Field6) + len(es.Field7) + len(es.Field8) + len(es.Field9) + len(es.Field10))
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
	buf = append(buf, `"}`...)
	return buf, nil
}
func (es *Struct10) UnmarshalJSON(data []byte) error {
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
	s = e + longFieldLength; es.Field10 = string(data[s:l])
	return nil
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
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
	}
}