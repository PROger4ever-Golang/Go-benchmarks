package jsonsStrings

import (
	"encoding/json"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsStrings"
)

func NewSampleJsonString10() string {
	var sampleJsonBytes, _ = json.Marshal(structsStrings.NewSampleStruct10())
	return string(sampleJsonBytes)
}