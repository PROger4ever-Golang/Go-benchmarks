package jsonsStrings

import (
	"encoding/json"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsStrings"
)

func NewSampleJsonString03() string {
	var sampleJsonBytes, _ = json.Marshal(structsStrings.NewSampleStruct03())
	return string(sampleJsonBytes)
}