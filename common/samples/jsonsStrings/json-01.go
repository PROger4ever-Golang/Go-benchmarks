package jsonsStrings

import (
	"encoding/json"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsStrings"
)

func NewSampleJsonString01() string {
	var sampleJsonBytes, _ = json.Marshal(structsStrings.NewSampleStruct01())
	return string(sampleJsonBytes)
}