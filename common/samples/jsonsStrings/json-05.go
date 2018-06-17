package jsonsStrings

import (
	"encoding/json"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsStrings"
)

func NewSampleJsonString05() string {
	var sampleJsonBytes, _ = json.Marshal(structsStrings.NewSampleStruct05())
	return string(sampleJsonBytes)
}