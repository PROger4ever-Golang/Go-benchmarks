package marshaling

import (
	"encoding/json"
	"reflect"
	"testing"

	fatihStructs "github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/jsonsStrings"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/mapsStrings"
	"github.com/PROger4ever-Golang/Go-benchmarks/reflection/structsMap"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkReflection$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/reflection
//BenchmarkReflection/01/UnmarshalMap-4                           	500000000	        34.3 ns/op	      16 B/op	       1 allocs/op
//BenchmarkReflection/01/UnmarshalReflection-4                    	100000000	       135 ns/op	      24 B/op	       2 allocs/op
//BenchmarkReflection/01/UnmarshalMitchellhMapstructure-4         	10000000	      1147 ns/op	     448 B/op	      12 allocs/op
//BenchmarkReflection/01/JsonUnmarshalMap-4                       	10000000	      1390 ns/op	     784 B/op	      12 allocs/op
//BenchmarkReflection/01/JsonUnmarshalStruct-4                    	20000000	       913 ns/op	     360 B/op	       5 allocs/op
//BenchmarkReflection/01/JsonUnmarshalStructUsingMap-4            	10000000	      1416 ns/op	     784 B/op	      12 allocs/op
//BenchmarkReflection/01/MarshalMap-4                             	100000000	       224 ns/op	     352 B/op	       3 allocs/op
//BenchmarkReflection/01/MarshalReflection-4                      	50000000	       418 ns/op	     376 B/op	       4 allocs/op
//BenchmarkReflection/01/MarshalFatihStructs-4                    	20000000	       793 ns/op	     560 B/op	       8 allocs/op
//BenchmarkReflection/01/JsonMarshalMap-4                         	10000000	      1304 ns/op	     800 B/op	      11 allocs/op
//BenchmarkReflection/01/JsonMarshalStruct-4                      	50000000	       387 ns/op	     208 B/op	       2 allocs/op
//BenchmarkReflection/01/JsonMarshalUsingMap-4                    	20000000	      1322 ns/op	     800 B/op	      11 allocs/op
//BenchmarkReflection/03/UnmarshalMap-4                           	200000000	        78.4 ns/op	      48 B/op	       1 allocs/op
//BenchmarkReflection/03/UnmarshalReflection-4                    	50000000	       306 ns/op	      72 B/op	       4 allocs/op
//BenchmarkReflection/03/UnmarshalMitchellhMapstructure-4         	 5000000	      3115 ns/op	    1504 B/op	      26 allocs/op
//BenchmarkReflection/03/JsonUnmarshalMap-4                       	 5000000	      2856 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkReflection/03/JsonUnmarshalStruct-4                    	10000000	      1921 ns/op	     472 B/op	       7 allocs/op
//BenchmarkReflection/03/JsonUnmarshalStructUsingMap-4            	 5000000	      2933 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkReflection/03/MarshalMap-4                             	50000000	       384 ns/op	     384 B/op	       5 allocs/op
//BenchmarkReflection/03/MarshalReflection-4                      	20000000	       771 ns/op	     456 B/op	       8 allocs/op
//BenchmarkReflection/03/MarshalFatihStructs-4                    	10000000	      2071 ns/op	    1280 B/op	      16 allocs/op
//BenchmarkReflection/03/JsonMarshalMap-4                         	 5000000	      2614 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkReflection/03/JsonMarshalStruct-4                      	20000000	       816 ns/op	     400 B/op	       3 allocs/op
//BenchmarkReflection/03/JsonMarshalUsingMap-4                    	 5000000	      2405 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkReflection/05/UnmarshalMap-4                           	100000000	       145 ns/op	      80 B/op	       1 allocs/op
//BenchmarkReflection/05/UnmarshalReflection-4                    	30000000	       578 ns/op	     120 B/op	       6 allocs/op
//BenchmarkReflection/05/UnmarshalMitchellhMapstructure-4         	 3000000	      5306 ns/op	    2816 B/op	      39 allocs/op
//BenchmarkReflection/05/JsonUnmarshalMap-4                       	 3000000	      4178 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkReflection/05/JsonUnmarshalStruct-4                    	 5000000	      2689 ns/op	     584 B/op	       9 allocs/op
//BenchmarkReflection/05/JsonUnmarshalStructUsingMap-4            	 3000000	      4215 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkReflection/05/MarshalMap-4                             	30000000	       574 ns/op	     416 B/op	       7 allocs/op
//BenchmarkReflection/05/MarshalReflection-4                      	20000000	       984 ns/op	     536 B/op	      12 allocs/op
//BenchmarkReflection/05/MarshalFatihStructs-4                    	 5000000	      3210 ns/op	    2272 B/op	      23 allocs/op
//BenchmarkReflection/05/JsonMarshalMap-4                         	 5000000	      3255 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkReflection/05/JsonMarshalStruct-4                      	20000000	      1001 ns/op	     448 B/op	       3 allocs/op
//BenchmarkReflection/05/JsonMarshalUsingMap-4                    	 5000000	      3294 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkReflection/10/UnmarshalMap-4                           	100000000	       241 ns/op	     160 B/op	       1 allocs/op
//BenchmarkReflection/10/UnmarshalReflection-4                    	20000000	       994 ns/op	     240 B/op	      11 allocs/op
//BenchmarkReflection/10/UnmarshalMitchellhMapstructure-4         	 1000000	     10983 ns/op	    6287 B/op	      72 allocs/op
//BenchmarkReflection/10/JsonUnmarshalMap-4                       	 2000000	      8293 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkReflection/10/JsonUnmarshalStruct-4                    	 3000000	      5104 ns/op	     872 B/op	      14 allocs/op
//BenchmarkReflection/10/JsonUnmarshalStructUsingMap-4            	 2000000	      8588 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkReflection/10/MarshalMap-4                             	20000000	      1200 ns/op	     790 B/op	      12 allocs/op
//BenchmarkReflection/10/MarshalReflection-4                      	10000000	      2340 ns/op	    1318 B/op	      23 allocs/op
//BenchmarkReflection/10/MarshalFatihStructs-4                    	 2000000	      8273 ns/op	    4886 B/op	      40 allocs/op
//BenchmarkReflection/10/JsonMarshalMap-4                         	 2000000	      6796 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkReflection/10/JsonMarshalStruct-4                      	10000000	      1832 ns/op	     896 B/op	       4 allocs/op
//BenchmarkReflection/10/JsonMarshalUsingMap-4                    	 2000000	      6656 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkReflection/17/UnmarshalMap-4                           	50000000	       444 ns/op	     288 B/op	       1 allocs/op
//BenchmarkReflection/17/UnmarshalReflection-4                    	10000000	      1706 ns/op	     424 B/op	      18 allocs/op
//BenchmarkReflection/17/UnmarshalMitchellhMapstructure-4         	 1000000	     22162 ns/op	   12991 B/op	     117 allocs/op
//BenchmarkReflection/17/JsonUnmarshalMap-4                       	 1000000	     15166 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkReflection/17/JsonUnmarshalStruct-4                    	 2000000	      9142 ns/op	    1304 B/op	      21 allocs/op
//BenchmarkReflection/17/JsonUnmarshalStructUsingMap-4            	 1000000	     15660 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkReflection/17/MarshalMap-4                             	10000000	      1903 ns/op	    1486 B/op	      19 allocs/op
//BenchmarkReflection/17/MarshalReflection-4                      	 3000000	      4653 ns/op	    2835 B/op	      38 allocs/op
//BenchmarkReflection/17/MarshalFatihStructs-4                    	 1000000	     11013 ns/op	    6475 B/op	      62 allocs/op
//BenchmarkReflection/17/JsonMarshalMap-4                         	 1000000	     11636 ns/op	    5102 B/op	      64 allocs/op
//BenchmarkReflection/17/JsonMarshalStruct-4                      	 5000000	      3078 ns/op	    1728 B/op	       5 allocs/op
//BenchmarkReflection/17/JsonMarshalUsingMap-4                    	 1000000	     11740 ns/op	    5102 B/op	      64 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/reflection	1139.801s


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

var sampleStruct01 = structsMap.NewSampleStruct01()
var sampleStruct03 = structsMap.NewSampleStruct03()
var sampleStruct05 = structsMap.NewSampleStruct05()
var sampleStruct10 = structsMap.NewSampleStruct10()
var sampleStruct17 = structsMap.NewSampleStruct17()

var sampleJsonString01 = jsonsStrings.NewSampleJsonString01()
var sampleJsonString03 = jsonsStrings.NewSampleJsonString03()
var sampleJsonString05 = jsonsStrings.NewSampleJsonString05()
var sampleJsonString10 = jsonsStrings.NewSampleJsonString10()
var sampleJsonString17 = jsonsStrings.NewSampleJsonString17()

var tmpStruct01 *structsMap.Struct01
var tmpStruct03 *structsMap.Struct03
var tmpStruct05 *structsMap.Struct05
var tmpStruct10 *structsMap.Struct10
var tmpStruct17 *structsMap.Struct17

var tmpMap map[string]interface{}
var tmpJsonBytes []byte
var tmpJsonString string

func BenchmarkReflection(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				tmpStruct01.UnmarshalMap(sampleMap01)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				UnmarshalReflection(tmpStruct01, sampleMap01)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				mapstructure.Decode(sampleMap01, tmpStruct01)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString01), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				json.Unmarshal([]byte(sampleJsonString01), tmpStruct01)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = &structsMap.Struct01{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString01), &tmpMap)
				tmpStruct01.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct01.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStruct01)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStruct01).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct01.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStruct01)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct01.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				tmpStruct03.UnmarshalMap(sampleMap03)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				UnmarshalReflection(tmpStruct03, sampleMap03)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				mapstructure.Decode(sampleMap03, tmpStruct03)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString03), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				json.Unmarshal([]byte(sampleJsonString03), tmpStruct03)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = &structsMap.Struct03{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString03), &tmpMap)
				tmpStruct03.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct03.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStruct03)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStruct03).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct03.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStruct03)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct03.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				tmpStruct05.UnmarshalMap(sampleMap05)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				UnmarshalReflection(tmpStruct05, sampleMap05)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				mapstructure.Decode(sampleMap05, tmpStruct05)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString05), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				json.Unmarshal([]byte(sampleJsonString05), tmpStruct05)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = &structsMap.Struct05{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString05), &tmpMap)
				tmpStruct05.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct05.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStruct05)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStruct05).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct05.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStruct05)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct05.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				tmpStruct10.UnmarshalMap(sampleMap10)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				UnmarshalReflection(tmpStruct10, sampleMap10)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				mapstructure.Decode(sampleMap10, tmpStruct10)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString10), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				json.Unmarshal([]byte(sampleJsonString10), tmpStruct10)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = &structsMap.Struct10{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString10), &tmpMap)
				tmpStruct10.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct10.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStruct10)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStruct10).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct10.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStruct10)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct10.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})

	b.Run("17", func(b *testing.B) {
		b.Run("UnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				tmpStruct17.UnmarshalMap(sampleMap17)
			}
		})

		b.Run("UnmarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				UnmarshalReflection(tmpStruct17, sampleMap17)
			}
		})

		b.Run("UnmarshalMitchellhMapstructure", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				mapstructure.Decode(sampleMap17, tmpStruct17)
			}
		})

		b.Run("JsonUnmarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString17), &tmpMap)
			}
		})

		b.Run("JsonUnmarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				json.Unmarshal([]byte(sampleJsonString17), tmpStruct17)
			}
		})

		b.Run("JsonUnmarshalStructUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = &structsMap.Struct17{}
				tmpMap = make(map[string]interface{})
				json.Unmarshal([]byte(sampleJsonString17), &tmpMap)
				tmpStruct17.UnmarshalMap(tmpMap)
			}
		})

		b.Run("MarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct17.MarshalMap()
			}
		})

		b.Run("MarshalReflection", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = MarshalReflection(sampleStruct17)
			}
		})

		b.Run("MarshalFatihStructs", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = fatihStructs.New(*sampleStruct17).Map()
			}
		})

		b.Run("JsonMarshalMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct17.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalStruct", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = json.Marshal(sampleStruct17)
				tmpJsonString = string(tmpJsonBytes)
			}
		})

		b.Run("JsonMarshalUsingMap", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpMap = sampleStruct17.MarshalMap()
				tmpJsonBytes, _ = json.Marshal(tmpMap)
				tmpJsonString = string(tmpJsonBytes)
			}
		})
	})
}
