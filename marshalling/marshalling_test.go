package marshaling

import (
	"encoding/json"
	"reflect"
	"testing"

	fatihStructs "github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/jsonsStrings"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/mapsStrings"
	"github.com/PROger4ever-Golang/Go-benchmarks/marshalling/structsMap"
	"github.com/PROger4ever-Golang/Go-benchmarks/marshalling/structsCustom"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkMarshalling$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/marshalling
//BenchmarkMarshalling/01/UnmarshalMap-4                           	500000000	        32.2 ns/op	      16 B/op	       1 allocs/op
//BenchmarkMarshalling/01/UnmarshalReflection-4                    	100000000	       137 ns/op	      24 B/op	       2 allocs/op
//BenchmarkMarshalling/01/UnmarshalMitchellhMapstructure-4         	20000000	      1081 ns/op	     448 B/op	      12 allocs/op
//BenchmarkMarshalling/01/JsonUnmarshalMap-4                       	10000000	      1257 ns/op	     784 B/op	      12 allocs/op
//BenchmarkMarshalling/01/JsonUnmarshalStruct-4                    	20000000	       857 ns/op	     360 B/op	       5 allocs/op
//BenchmarkMarshalling/01/JsonUnmarshalCustom-4                    	20000000	       629 ns/op	     368 B/op	       6 allocs/op
//BenchmarkMarshalling/01/JsonUnmarshalStructUsingMap-4            	10000000	      1252 ns/op	     784 B/op	      12 allocs/op
//BenchmarkMarshalling/01/MarshalMap-4                             	100000000	       189 ns/op	     352 B/op	       3 allocs/op
//BenchmarkMarshalling/01/MarshalReflection-4                      	50000000	       280 ns/op	     376 B/op	       4 allocs/op
//BenchmarkMarshalling/01/MarshalFatihStructs-4                    	20000000	       669 ns/op	     560 B/op	       8 allocs/op
//BenchmarkMarshalling/01/JsonMarshalMap-4                         	20000000	      1035 ns/op	     800 B/op	      11 allocs/op
//BenchmarkMarshalling/01/JsonMarshalStruct-4                      	50000000	       342 ns/op	     208 B/op	       2 allocs/op
//BenchmarkMarshalling/01/JsonMarshalCustom-4                      	30000000	       566 ns/op	     344 B/op	       5 allocs/op
//BenchmarkMarshalling/01/JsonMarshalUsingMap-4                    	20000000	      1035 ns/op	     800 B/op	      11 allocs/op
//BenchmarkMarshalling/03/UnmarshalMap-4                           	200000000	        66.1 ns/op	      48 B/op	       1 allocs/op
//BenchmarkMarshalling/03/UnmarshalReflection-4                    	50000000	       315 ns/op	      72 B/op	       4 allocs/op
//BenchmarkMarshalling/03/UnmarshalMitchellhMapstructure-4         	 5000000	      2790 ns/op	    1504 B/op	      26 allocs/op
//BenchmarkMarshalling/03/JsonUnmarshalMap-4                       	 5000000	      2760 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkMarshalling/03/JsonUnmarshalStruct-4                    	10000000	      1848 ns/op	     472 B/op	       7 allocs/op
//BenchmarkMarshalling/03/JsonUnmarshalCustom-4                    	20000000	      1059 ns/op	     480 B/op	       8 allocs/op
//BenchmarkMarshalling/03/JsonUnmarshalStructUsingMap-4            	 5000000	      2820 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkMarshalling/03/MarshalMap-4                             	50000000	       302 ns/op	     384 B/op	       5 allocs/op
//BenchmarkMarshalling/03/MarshalReflection-4                      	30000000	       568 ns/op	     456 B/op	       8 allocs/op
//BenchmarkMarshalling/03/MarshalFatihStructs-4                    	10000000	      1658 ns/op	    1280 B/op	      16 allocs/op
//BenchmarkMarshalling/03/JsonMarshalMap-4                         	10000000	      2250 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkMarshalling/03/JsonMarshalStruct-4                      	20000000	       758 ns/op	     400 B/op	       3 allocs/op
//BenchmarkMarshalling/03/JsonMarshalCustom-4                      	20000000	      1034 ns/op	     520 B/op	       6 allocs/op
//BenchmarkMarshalling/03/JsonMarshalUsingMap-4                    	10000000	      2244 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkMarshalling/05/UnmarshalMap-4                           	100000000	       128 ns/op	      80 B/op	       1 allocs/op
//BenchmarkMarshalling/05/UnmarshalReflection-4                    	30000000	       564 ns/op	     120 B/op	       6 allocs/op
//BenchmarkMarshalling/05/UnmarshalMitchellhMapstructure-4         	 3000000	      4478 ns/op	    2816 B/op	      39 allocs/op
//BenchmarkMarshalling/05/JsonUnmarshalMap-4                       	 3000000	      4194 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkMarshalling/05/JsonUnmarshalStruct-4                    	 5000000	      2738 ns/op	     584 B/op	       9 allocs/op
//BenchmarkMarshalling/05/JsonUnmarshalCustom-4                    	10000000	      1448 ns/op	     592 B/op	      10 allocs/op
//BenchmarkMarshalling/05/JsonUnmarshalStructUsingMap-4            	 3000000	      4323 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkMarshalling/05/MarshalMap-4                             	30000000	       404 ns/op	     416 B/op	       7 allocs/op
//BenchmarkMarshalling/05/MarshalReflection-4                      	20000000	       839 ns/op	     536 B/op	      12 allocs/op
//BenchmarkMarshalling/05/MarshalFatihStructs-4                    	 5000000	      2688 ns/op	    2272 B/op	      23 allocs/op
//BenchmarkMarshalling/05/JsonMarshalMap-4                         	 5000000	      3086 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkMarshalling/05/JsonMarshalStruct-4                      	20000000	       949 ns/op	     448 B/op	       3 allocs/op
//BenchmarkMarshalling/05/JsonMarshalCustom-4                      	10000000	      1353 ns/op	     664 B/op	       6 allocs/op
//BenchmarkMarshalling/05/JsonMarshalUsingMap-4                    	 5000000	      3073 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkMarshalling/10/UnmarshalMap-4                           	100000000	       200 ns/op	     160 B/op	       1 allocs/op
//BenchmarkMarshalling/10/UnmarshalReflection-4                    	20000000	       989 ns/op	     240 B/op	      11 allocs/op
//BenchmarkMarshalling/10/UnmarshalMitchellhMapstructure-4         	 2000000	      9476 ns/op	    6287 B/op	      72 allocs/op
//BenchmarkMarshalling/10/JsonUnmarshalMap-4                       	 2000000	      8468 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkMarshalling/10/JsonUnmarshalStruct-4                    	 3000000	      5174 ns/op	     872 B/op	      14 allocs/op
//BenchmarkMarshalling/10/JsonUnmarshalCustom-4                    	 5000000	      2460 ns/op	     880 B/op	      15 allocs/op
//BenchmarkMarshalling/10/JsonUnmarshalStructUsingMap-4            	 2000000	      8744 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkMarshalling/10/MarshalMap-4                             	20000000	       905 ns/op	     790 B/op	      12 allocs/op
//BenchmarkMarshalling/10/MarshalReflection-4                      	10000000	      2051 ns/op	    1318 B/op	      23 allocs/op
//BenchmarkMarshalling/10/MarshalFatihStructs-4                    	 3000000	      5505 ns/op	    4886 B/op	      40 allocs/op
//BenchmarkMarshalling/10/JsonMarshalMap-4                         	 2000000	      6207 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkMarshalling/10/JsonMarshalStruct-4                      	10000000	      1689 ns/op	     896 B/op	       4 allocs/op
//BenchmarkMarshalling/10/JsonMarshalCustom-4                      	10000000	      2240 ns/op	    1048 B/op	       6 allocs/op
//BenchmarkMarshalling/10/JsonMarshalUsingMap-4                    	 2000000	      6242 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkMarshalling/17/UnmarshalMap-4                           	50000000	       331 ns/op	     288 B/op	       1 allocs/op
//BenchmarkMarshalling/17/UnmarshalReflection-4                    	10000000	      1757 ns/op	     424 B/op	      18 allocs/op
//BenchmarkMarshalling/17/UnmarshalMitchellhMapstructure-4         	 1000000	     17723 ns/op	   12991 B/op	     117 allocs/op
//BenchmarkMarshalling/17/JsonUnmarshalMap-4                       	 1000000	     15283 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkMarshalling/17/JsonUnmarshalStruct-4                    	 2000000	      9385 ns/op	    1304 B/op	      21 allocs/op
//BenchmarkMarshalling/17/JsonUnmarshalCustom-4                    	 3000000	      4083 ns/op	    1312 B/op	      22 allocs/op
//BenchmarkMarshalling/17/JsonUnmarshalStructUsingMap-4            	 1000000	     15859 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkMarshalling/17/MarshalMap-4                             	10000000	      1527 ns/op	    1486 B/op	      19 allocs/op
//BenchmarkMarshalling/17/MarshalReflection-4                      	 3000000	      4088 ns/op	    2835 B/op	      38 allocs/op
//BenchmarkMarshalling/17/MarshalFatihStructs-4                    	 2000000	      9093 ns/op	    6475 B/op	      62 allocs/op
//BenchmarkMarshalling/17/JsonMarshalMap-4                         	 1000000	     10883 ns/op	    5102 B/op	      64 allocs/op
//BenchmarkMarshalling/17/JsonMarshalStruct-4                      	 5000000	      2698 ns/op	    1728 B/op	       5 allocs/op
//BenchmarkMarshalling/17/JsonMarshalCustom-4                      	 5000000	      3691 ns/op	    1624 B/op	       6 allocs/op
//BenchmarkMarshalling/17/JsonMarshalUsingMap-4                    	 1000000	     11137 ns/op	    5102 B/op	      64 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/marshalling	1295.486s


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

var sampleMap01 = mapsStrings.NewSampleMap01()
var sampleMap03 = mapsStrings.NewSampleMap03()
var sampleMap05 = mapsStrings.NewSampleMap05()
var sampleMap10 = mapsStrings.NewSampleMap10()
var sampleMap17 = mapsStrings.NewSampleMap17()

var sampleStructMap01 = structsMap.NewSampleStruct01()
var sampleStructMap03 = structsMap.NewSampleStruct03()
var sampleStructMap05 = structsMap.NewSampleStruct05()
var sampleStructMap10 = structsMap.NewSampleStruct10()
var sampleStructMap17 = structsMap.NewSampleStruct17()

var sampleStructCustom01 = structsCustom.NewSampleStruct01()
var sampleStructCustom03 = structsCustom.NewSampleStruct03()
var sampleStructCustom05 = structsCustom.NewSampleStruct05()
var sampleStructCustom10 = structsCustom.NewSampleStruct10()
var sampleStructCustom17 = structsCustom.NewSampleStruct17()

var sampleJsonString01 = jsonsStrings.NewSampleJsonString01()
var sampleJsonString03 = jsonsStrings.NewSampleJsonString03()
var sampleJsonString05 = jsonsStrings.NewSampleJsonString05()
var sampleJsonString10 = jsonsStrings.NewSampleJsonString10()
var sampleJsonString17 = jsonsStrings.NewSampleJsonString17()

var tmpStructMap01 *structsMap.Struct01
var tmpStructMap03 *structsMap.Struct03
var tmpStructMap05 *structsMap.Struct05
var tmpStructMap10 *structsMap.Struct10
var tmpStructMap17 *structsMap.Struct17

var tmpStructCustom01 *structsCustom.Struct01
var tmpStructCustom03 *structsCustom.Struct03
var tmpStructCustom05 *structsCustom.Struct05
var tmpStructCustom10 *structsCustom.Struct10
var tmpStructCustom17 *structsCustom.Struct17

var tmpMap map[string]interface{}
var tmpJsonBytes []byte
var tmpJsonString string

func BenchmarkMarshalling(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				tmpStructMap01.UnmarshalMap(sampleMap01)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				UnmarshalReflection(tmpStructMap01, sampleMap01)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				mapstructure.Decode(sampleMap01, tmpStructMap01)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString01), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				json.Unmarshal([]byte(sampleJsonString01), tmpStructMap01)
			}
		})

		b.Run("JsonUnmarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom01 = &structsCustom.Struct01{}
				json.Unmarshal([]byte(sampleJsonString01), tmpStructCustom01)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap01 = &structsMap.Struct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString01), &tmpMap)
				tmpStructMap01.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap01.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStructMap01)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStructMap01).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap01.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructMap01)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructCustom01)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap01.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				tmpStructMap03.UnmarshalMap(sampleMap03)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				UnmarshalReflection(tmpStructMap03, sampleMap03)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				mapstructure.Decode(sampleMap03, tmpStructMap03)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString03), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				json.Unmarshal([]byte(sampleJsonString03), tmpStructMap03)
			}
		})

		b.Run("JsonUnmarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom03 = &structsCustom.Struct03{}
				json.Unmarshal([]byte(sampleJsonString03), tmpStructCustom03)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap03 = &structsMap.Struct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString03), &tmpMap)
				tmpStructMap03.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap03.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStructMap03)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStructMap03).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap03.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructMap03)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructCustom03)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap03.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				tmpStructMap05.UnmarshalMap(sampleMap05)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				UnmarshalReflection(tmpStructMap05, sampleMap05)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				mapstructure.Decode(sampleMap05, tmpStructMap05)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString05), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				json.Unmarshal([]byte(sampleJsonString05), tmpStructMap05)
			}
		})

		b.Run("JsonUnmarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom05 = &structsCustom.Struct05{}
				json.Unmarshal([]byte(sampleJsonString05), tmpStructCustom05)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap05 = &structsMap.Struct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString05), &tmpMap)
				tmpStructMap05.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap05.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStructMap05)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStructMap05).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap05.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructMap05)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructCustom05)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap05.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				tmpStructMap10.UnmarshalMap(sampleMap10)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				UnmarshalReflection(tmpStructMap10, sampleMap10)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				mapstructure.Decode(sampleMap10, tmpStructMap10)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString10), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				json.Unmarshal([]byte(sampleJsonString10), tmpStructMap10)
			}
		})

		b.Run("JsonUnmarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom10 = &structsCustom.Struct10{}
				json.Unmarshal([]byte(sampleJsonString10), tmpStructCustom10)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap10 = &structsMap.Struct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString10), &tmpMap)
				tmpStructMap10.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap10.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStructMap10)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStructMap10).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap10.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructMap10)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructCustom10)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap10.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("17", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				tmpStructMap17.UnmarshalMap(sampleMap17)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				UnmarshalReflection(tmpStructMap17, sampleMap17)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				mapstructure.Decode(sampleMap17, tmpStructMap17)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString17), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				json.Unmarshal([]byte(sampleJsonString17), tmpStructMap17)
			}
		})

		b.Run("JsonUnmarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom17 = &structsCustom.Struct17{}
				json.Unmarshal([]byte(sampleJsonString17), tmpStructCustom17)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructMap17 = &structsMap.Struct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString17), &tmpMap)
				tmpStructMap17.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap17.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStructMap17)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStructMap17).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap17.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructMap17)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalCustom", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStructCustom17)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStructMap17.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})
}
