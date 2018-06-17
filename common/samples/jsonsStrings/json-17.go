package jsonsStrings

import (
	"encoding/json"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsStrings"
)

func NewSampleJsonString17() string {
	var sampleJsonBytes, _ = json.Marshal(structsStrings.NewSampleStruct17())
	return string(sampleJsonBytes)
}