package interfaces

import (
	"testing"
	"reflect"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
)

//approximate result
//$ go test -bench=BenchmarkReflection -benchtime=10s --timeout 99999s ./...
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/interfaces	0.001s
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/reflection
//BenchmarkReflection/01/UnmarshalManual-4         	500000000	        32.8 ns/op
//BenchmarkReflection/01/UnmarshalReflection-4     	100000000	       132 ns/op
//BenchmarkReflection/01/UnmarshalMitchellhMapstructure-4         	20000000	      1119 ns/op
//BenchmarkReflection/01/JsonUnmarshalMap-4                       	 1000000	     15324 ns/op
//BenchmarkReflection/01/JsonUnmarshalStruct-4                    	 2000000	      9287 ns/op
//BenchmarkReflection/01/JsonUnmarshalStructUsingMap-4            	 1000000	     15368 ns/op
//BenchmarkReflection/01/MarshalManual-4                          	100000000	       210 ns/op
//BenchmarkReflection/01/MarshalReflection-4                      	50000000	       370 ns/op
//BenchmarkReflection/01/MarshalFatihStructs-4                    	20000000	       808 ns/op
//BenchmarkReflection/01/JsonMarshalMap-4                         	20000000	      1184 ns/op
//BenchmarkReflection/01/JsonMarshalStruct-4                      	30000000	       374 ns/op
//BenchmarkReflection/01/JsonMarshalUsingMap-4                    	10000000	      1264 ns/op
//BenchmarkReflection/03/UnmarshalManual-4                        	200000000	        75.4 ns/op
//BenchmarkReflection/03/UnmarshalReflection-4                    	50000000	       301 ns/op
//BenchmarkReflection/03/UnmarshalMitchellhMapstructure-4         	 5000000	      3034 ns/op
//BenchmarkReflection/03/JsonUnmarshalMap-4                       	 1000000	     15220 ns/op
//BenchmarkReflection/03/JsonUnmarshalStruct-4                    	 2000000	      9485 ns/op
//BenchmarkReflection/03/JsonUnmarshalStructUsingMap-4            	 1000000	     15328 ns/op
//BenchmarkReflection/03/MarshalManual-4                          	50000000	       382 ns/op
//BenchmarkReflection/03/MarshalReflection-4                      	20000000	       698 ns/op
//BenchmarkReflection/03/MarshalFatihStructs-4                    	10000000	      1985 ns/op
//BenchmarkReflection/03/JsonMarshalMap-4                         	 5000000	      2394 ns/op
//BenchmarkReflection/03/JsonMarshalStruct-4                      	20000000	       769 ns/op
//BenchmarkReflection/03/JsonMarshalUsingMap-4                    	10000000	      2455 ns/op
//BenchmarkReflection/05/UnmarshalManual-4                        	100000000	       139 ns/op
//BenchmarkReflection/05/UnmarshalReflection-4                    	30000000	       540 ns/op
//BenchmarkReflection/05/UnmarshalMitchellhMapstructure-4         	 3000000	      4991 ns/op
//BenchmarkReflection/05/JsonUnmarshalMap-4                       	 1000000	     15213 ns/op
//BenchmarkReflection/05/JsonUnmarshalStruct-4                    	 2000000	      9529 ns/op
//BenchmarkReflection/05/JsonUnmarshalStructUsingMap-4            	 1000000	     15431 ns/op
//BenchmarkReflection/05/MarshalManual-4                          	30000000	       507 ns/op
//BenchmarkReflection/05/MarshalReflection-4                      	20000000	       956 ns/op
//BenchmarkReflection/05/MarshalFatihStructs-4                    	 5000000	      3237 ns/op
//BenchmarkReflection/05/JsonMarshalMap-4                         	 5000000	      3266 ns/op
//BenchmarkReflection/05/JsonMarshalStruct-4                      	20000000	       960 ns/op
//BenchmarkReflection/05/JsonMarshalUsingMap-4                    	 5000000	      3300 ns/op
//BenchmarkReflection/10/UnmarshalManual-4                        	100000000	       229 ns/op
//BenchmarkReflection/10/UnmarshalReflection-4                    	20000000	       981 ns/op
//BenchmarkReflection/10/UnmarshalMitchellhMapstructure-4         	 1000000	     11045 ns/op
//BenchmarkReflection/10/JsonUnmarshalMap-4                       	 1000000	     15281 ns/op
//BenchmarkReflection/10/JsonUnmarshalStruct-4                    	 2000000	      9486 ns/op
//BenchmarkReflection/10/JsonUnmarshalStructUsingMap-4            	 1000000	     15728 ns/op
//BenchmarkReflection/10/MarshalManual-4                          	20000000	      1117 ns/op
//BenchmarkReflection/10/MarshalReflection-4                      	10000000	      2267 ns/op
//BenchmarkReflection/10/MarshalFatihStructs-4                    	 2000000	      6811 ns/op
//BenchmarkReflection/10/JsonMarshalMap-4                         	 2000000	      6419 ns/op
//BenchmarkReflection/10/JsonMarshalStruct-4                      	10000000	      1746 ns/op
//BenchmarkReflection/10/JsonMarshalUsingMap-4                    	 2000000	      6506 ns/op
//BenchmarkReflection/17/UnmarshalManual-4                        	50000000	       416 ns/op
//BenchmarkReflection/17/UnmarshalReflection-4                    	10000000	      1652 ns/op
//BenchmarkReflection/17/UnmarshalMitchellhMapstructure-4         	 1000000	     21091 ns/op
//BenchmarkReflection/17/JsonUnmarshalMap-4                       	 1000000	     15437 ns/op
//BenchmarkReflection/17/JsonUnmarshalStruct-4                    	 2000000	      9281 ns/op
//BenchmarkReflection/17/JsonUnmarshalStructUsingMap-4            	 1000000	     16038 ns/op
//BenchmarkReflection/17/MarshalManual-4                          	10000000	      1758 ns/op
//BenchmarkReflection/17/MarshalReflection-4                      	 3000000	      4556 ns/op
//BenchmarkReflection/17/MarshalFatihStructs-4                    	 1000000	     10569 ns/op
//BenchmarkReflection/17/JsonMarshalMap-4                         	 1000000	     10870 ns/op
//BenchmarkReflection/17/JsonMarshalStruct-4                      	 5000000	      3148 ns/op
//BenchmarkReflection/17/JsonMarshalUsingMap-4                    	 1000000	     11353 ns/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/reflection	1129.891s


func MarshalReflection(object interface{}) map[string]interface{} {
	theMap := make(map[string]interface{})
	t := reflect.TypeOf(object).Elem()
	v := reflect.ValueOf(object).Elem()
	for i := 0; i < v.NumField(); i++ {
		theMap[t.Field(i).Name] = v.Field(i)
	}	
	return theMap
}

func UnmarshalReflection(object interface{}, theMap map[string]interface{}) {
	t := reflect.TypeOf(object).Elem()
	v := reflect.ValueOf(object).Elem()
	for i := 0; i < t.NumField(); i++ {
		if vv, ok := theMap[t.Field(i).Name]; ok {
			v.Field(i).Set(reflect.ValueOf(vv))
		}
	}
}

type ExampleStruct01 struct {
	Field1 string
}

func (es *ExampleStruct01) MarshalManual() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
	}
}

func (es *ExampleStruct01) UnmarshalManual(theMap map[string]interface{}) *ExampleStruct01 {
	es.Field1 = theMap["Field1"].(string)
	return es
}

type ExampleStruct03 struct {
	Field1, Field2, Field3 string
}

func (es *ExampleStruct03) MarshalManual() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
		"Field2": es.Field2,
		"Field3": es.Field3,
	}
}

func (es *ExampleStruct03) UnmarshalManual(theMap map[string]interface{}) *ExampleStruct03 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	return es
}

type ExampleStruct05 struct {
	Field1, Field2, Field3, Field4, Field5 string
}

func (es *ExampleStruct05) MarshalManual() map[string]interface{} {
	return map[string]interface{}{
		"Field1": es.Field1,
		"Field2": es.Field2,
		"Field3": es.Field3,
		"Field4": es.Field4,
		"Field5": es.Field5,
	}
}

func (es *ExampleStruct05) UnmarshalManual(theMap map[string]interface{}) *ExampleStruct05 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	es.Field4 = theMap["Field4"].(string)
	es.Field5 = theMap["Field5"].(string)
	return es
}

type ExampleStruct10 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10                                                                string
}

func (es *ExampleStruct10) MarshalManual() map[string]interface{} {
	return map[string]interface{}{
		"Field1":  es.Field1,
		"Field2":  es.Field2,
		"Field3":  es.Field3,
		"Field4":  es.Field4,
		"Field5":  es.Field5,
		"Field6":  es.Field6,
		"Field7":  es.Field7,
		"Field8":  es.Field8,
		"Field9":  es.Field9,
		"Field10": es.Field10,
	}
}

func (es *ExampleStruct10) UnmarshalManual(theMap map[string]interface{}) *ExampleStruct10 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	es.Field4 = theMap["Field4"].(string)
	es.Field5 = theMap["Field5"].(string)
	es.Field6 = theMap["Field6"].(string)
	es.Field7 = theMap["Field7"].(string)
	es.Field8 = theMap["Field8"].(string)
	es.Field9 = theMap["Field9"].(string)
	es.Field10 = theMap["Field10"].(string)
	return es
}

type ExampleStruct17 struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9 string
	Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17 string
}

func (es *ExampleStruct17) MarshalManual() map[string]interface{} {
	return map[string]interface{}{
		"Field1":  es.Field1,
		"Field2":  es.Field2,
		"Field3":  es.Field3,
		"Field4":  es.Field4,
		"Field5":  es.Field5,
		"Field6":  es.Field6,
		"Field7":  es.Field7,
		"Field8":  es.Field8,
		"Field9":  es.Field9,
		"Field10": es.Field10,
		"Field11": es.Field11,
		"Field12": es.Field12,
		"Field13": es.Field13,
		"Field14": es.Field14,
		"Field15": es.Field15,
		"Field16": es.Field16,
		"Field17": es.Field17,
	}
}

func (es *ExampleStruct17) UnmarshalManual(theMap map[string]interface{}) *ExampleStruct17 {
	es.Field1 = theMap["Field1"].(string)
	es.Field2 = theMap["Field2"].(string)
	es.Field3 = theMap["Field3"].(string)
	es.Field4 = theMap["Field4"].(string)
	es.Field5 = theMap["Field5"].(string)
	es.Field6 = theMap["Field6"].(string)
	es.Field7 = theMap["Field7"].(string)
	es.Field8 = theMap["Field8"].(string)
	es.Field9 = theMap["Field9"].(string)
	es.Field10 = theMap["Field10"].(string)
	es.Field11 = theMap["Field11"].(string)
	es.Field12 = theMap["Field12"].(string)
	es.Field13 = theMap["Field13"].(string)
	es.Field14 = theMap["Field14"].(string)
	es.Field15 = theMap["Field15"].(string)
	es.Field16 = theMap["Field16"].(string)
	es.Field17 = theMap["Field17"].(string)
	return es
}


var exampleMap01 = map[string]interface{}{
	"Field1":  "Field1-Value",
}
var exampleMap03 = map[string]interface{}{
	"Field1":  "Field1-Value",
	"Field2":  "Field2-Value",
	"Field3":  "Field3-Value",
}
var exampleMap05 = map[string]interface{}{
	"Field1":  "Field1-Value",
	"Field2":  "Field2-Value",
	"Field3":  "Field3-Value",
	"Field4":  "Field4-Value",
	"Field5":  "Field5-Value",
}
var exampleMap10 = map[string]interface{}{
	"Field1":  "Field1-Value",
	"Field2":  "Field2-Value",
	"Field3":  "Field3-Value",
	"Field4":  "Field4-Value",
	"Field5":  "Field5-Value",
	"Field6":  "Field6-Value",
	"Field7":  "Field7-Value",
	"Field8":  "Field8-Value",
	"Field9":  "Field9-Value",
	"Field10": "Field10-Value",
}
var exampleMap17 = map[string]interface{}{
	"Field1":  "Field1-Value",
	"Field2":  "Field2-Value",
	"Field3":  "Field3-Value",
	"Field4":  "Field4-Value",
	"Field5":  "Field5-Value",
	"Field6":  "Field6-Value",
	"Field7":  "Field7-Value",
	"Field8":  "Field8-Value",
	"Field9":  "Field9-Value",
	"Field10": "Field10-Value",
	"Field11": "Field11-Value",
	"Field12": "Field12-Value",
	"Field13": "Field13-Value",
	"Field14": "Field14-Value",
	"Field15": "Field15-Value",
	"Field16": "Field16-Value",
	"Field17": "Field17-Value",
}
var exampleJsonBytes, _ = json.Marshal(exampleMap17)
var exampleJsonString = string(exampleJsonBytes)


var tmpMap map[string]interface{}
var tmpJsonBytes []byte
var tmpJsonString string


var tmpStruct01 *ExampleStruct01
var tmpStruct03 *ExampleStruct03
var tmpStruct05 *ExampleStruct05
var tmpStruct10 *ExampleStruct10
var tmpStruct17 *ExampleStruct17

func BenchmarkReflection(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("UnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				tmpStruct01.UnmarshalManual(exampleMap01)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				UnmarshalReflection(tmpStruct01, exampleMap01)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				mapstructure.Decode(exampleMap01, tmpStruct01)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				json.Unmarshal([]byte(exampleJsonString), tmpStruct01)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &ExampleStruct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
				tmpStruct01.UnmarshalManual(tmpMap)
			}
		})

		b.Run("MarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct01.MarshalManual()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(tmpStruct01)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = structs.New(*tmpStruct01).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct01.MarshalManual()
				exampleJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(exampleJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(tmpStruct01)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct01.MarshalManual()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("UnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				tmpStruct03.UnmarshalManual(exampleMap03)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				UnmarshalReflection(tmpStruct03, exampleMap03)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				mapstructure.Decode(exampleMap03, tmpStruct03)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				json.Unmarshal([]byte(exampleJsonString), tmpStruct03)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &ExampleStruct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
				tmpStruct03.UnmarshalManual(tmpMap)
			}
		})

		b.Run("MarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct03.MarshalManual()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(tmpStruct03)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = structs.New(*tmpStruct03).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct03.MarshalManual()
				exampleJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(exampleJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(tmpStruct03)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct03.MarshalManual()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("UnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				tmpStruct05.UnmarshalManual(exampleMap05)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				UnmarshalReflection(tmpStruct05, exampleMap05)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				mapstructure.Decode(exampleMap05, tmpStruct05)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				json.Unmarshal([]byte(exampleJsonString), tmpStruct05)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &ExampleStruct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
				tmpStruct05.UnmarshalManual(tmpMap)
			}
		})

		b.Run("MarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct05.MarshalManual()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(tmpStruct05)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = structs.New(*tmpStruct05).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct05.MarshalManual()
				exampleJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(exampleJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(tmpStruct05)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct05.MarshalManual()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("UnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				tmpStruct10.UnmarshalManual(exampleMap10)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				UnmarshalReflection(tmpStruct10, exampleMap10)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				mapstructure.Decode(exampleMap10, tmpStruct10)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				json.Unmarshal([]byte(exampleJsonString), tmpStruct10)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &ExampleStruct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
				tmpStruct10.UnmarshalManual(tmpMap)
			}
		})

		b.Run("MarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct10.MarshalManual()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(tmpStruct10)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = structs.New(*tmpStruct10).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct10.MarshalManual()
				exampleJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(exampleJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(tmpStruct10)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct10.MarshalManual()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("17", func(b *testing.B) {
		b.Run("UnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				tmpStruct17.UnmarshalManual(exampleMap17)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				UnmarshalReflection(tmpStruct17, exampleMap17)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				mapstructure.Decode(exampleMap17, tmpStruct17)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				json.Unmarshal([]byte(exampleJsonString), tmpStruct17)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &ExampleStruct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(exampleJsonString), &tmpMap)
				tmpStruct17.UnmarshalManual(tmpMap)
			}
		})

		b.Run("MarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct17.MarshalManual()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(tmpStruct17)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = structs.New(*tmpStruct17).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct17.MarshalManual()
				exampleJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(exampleJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(tmpStruct17)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = tmpStruct17.MarshalManual()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})
}
