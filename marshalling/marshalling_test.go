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

//TODO: add benchmarks for external json-libraries and marshal-codegenerators like github.com/pquerna/ffjson and github.com/mailru/easyjson

//approximate result
//$ go test -benchmem -bench=^BenchmarkMarshalling.*$ -benchtime=10s --timeout 99999s ./...
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/interfaces	0.002s
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/marshalling
//BenchmarkMarshallingToMaps/01/UnmarshalMap-4                           	500000000	        31.8 ns/op	      16 B/op	       1 allocs/op
//BenchmarkMarshallingToMaps/01/UnmarshalReflection-4                    	100000000	       134 ns/op	      24 B/op	       2 allocs/op
//BenchmarkMarshallingToMaps/01/UnmarshalMitchellhMapstructure-4         	20000000	      1072 ns/op	     448 B/op	      12 allocs/op
//BenchmarkMarshallingToMaps/01/MarshalMap-4                             	100000000	       187 ns/op	     352 B/op	       3 allocs/op
//BenchmarkMarshallingToMaps/01/MarshalReflection-4                      	50000000	       282 ns/op	     376 B/op	       4 allocs/op
//BenchmarkMarshallingToMaps/01/MarshalFatihStructs-4                    	20000000	       667 ns/op	     560 B/op	       8 allocs/op
//BenchmarkMarshallingToMaps/03/UnmarshalMap-4                           	200000000	        66.1 ns/op	      48 B/op	       1 allocs/op
//BenchmarkMarshallingToMaps/03/UnmarshalReflection-4                    	50000000	       317 ns/op	      72 B/op	       4 allocs/op
//BenchmarkMarshallingToMaps/03/UnmarshalMitchellhMapstructure-4         	 5000000	      2771 ns/op	    1504 B/op	      26 allocs/op
//BenchmarkMarshallingToMaps/03/MarshalMap-4                             	50000000	       295 ns/op	     384 B/op	       5 allocs/op
//BenchmarkMarshallingToMaps/03/MarshalReflection-4                      	30000000	       563 ns/op	     456 B/op	       8 allocs/op
//BenchmarkMarshallingToMaps/03/MarshalFatihStructs-4                    	10000000	      1679 ns/op	    1280 B/op	      16 allocs/op
//BenchmarkMarshallingToMaps/05/UnmarshalMap-4                           	100000000	       127 ns/op	      80 B/op	       1 allocs/op
//BenchmarkMarshallingToMaps/05/UnmarshalReflection-4                    	30000000	       564 ns/op	     120 B/op	       6 allocs/op
//BenchmarkMarshallingToMaps/05/UnmarshalMitchellhMapstructure-4         	 3000000	      4480 ns/op	    2816 B/op	      39 allocs/op
//BenchmarkMarshallingToMaps/05/MarshalMap-4                             	30000000	       404 ns/op	     416 B/op	       7 allocs/op
//BenchmarkMarshallingToMaps/05/MarshalReflection-4                      	20000000	       856 ns/op	     536 B/op	      12 allocs/op
//BenchmarkMarshallingToMaps/05/MarshalFatihStructs-4                    	 5000000	      2714 ns/op	    2272 B/op	      23 allocs/op
//BenchmarkMarshallingToMaps/10/UnmarshalMap-4                           	100000000	       201 ns/op	     160 B/op	       1 allocs/op
//BenchmarkMarshallingToMaps/10/UnmarshalReflection-4                    	20000000	      1008 ns/op	     240 B/op	      11 allocs/op
//BenchmarkMarshallingToMaps/10/UnmarshalMitchellhMapstructure-4         	 2000000	      9497 ns/op	    6287 B/op	      72 allocs/op
//BenchmarkMarshallingToMaps/10/MarshalMap-4                             	20000000	       878 ns/op	     790 B/op	      12 allocs/op
//BenchmarkMarshallingToMaps/10/MarshalReflection-4                      	10000000	      2044 ns/op	    1318 B/op	      23 allocs/op
//BenchmarkMarshallingToMaps/10/MarshalFatihStructs-4                    	 3000000	      5412 ns/op	    4886 B/op	      40 allocs/op
//BenchmarkMarshallingToMaps/17/UnmarshalMap-4                           	50000000	       336 ns/op	     288 B/op	       1 allocs/op
//BenchmarkMarshallingToMaps/17/UnmarshalReflection-4                    	10000000	      1755 ns/op	     424 B/op	      18 allocs/op
//BenchmarkMarshallingToMaps/17/UnmarshalMitchellhMapstructure-4         	 1000000	     17584 ns/op	   12991 B/op	     117 allocs/op
//BenchmarkMarshallingToMaps/17/MarshalMap-4                             	10000000	      1533 ns/op	    1486 B/op	      19 allocs/op
//BenchmarkMarshallingToMaps/17/MarshalReflection-4                      	 5000000	      4008 ns/op	    2835 B/op	      38 allocs/op
//BenchmarkMarshallingToMaps/17/MarshalFatihStructs-4                    	 2000000	      9212 ns/op	    6475 B/op	      62 allocs/op
//BenchmarkMarshallingToJson/01/JsonUnmarshalMap-4                       	10000000	      1253 ns/op	     784 B/op	      12 allocs/op
//BenchmarkMarshallingToJson/01/JsonUnmarshalStruct-4                    	20000000	       861 ns/op	     360 B/op	       5 allocs/op
//BenchmarkMarshallingToJson/01/JsonUnmarshalCustom-4                    	20000000	       631 ns/op	     368 B/op	       6 allocs/op
//BenchmarkMarshallingToJson/01/JsonUnmarshalManual-4                    	300000000	        55.5 ns/op	      32 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/01/JsonUnmarshalStructUsingMap-4            	10000000	      1255 ns/op	     784 B/op	      12 allocs/op
//BenchmarkMarshallingToJson/01/JsonMarshalMap-4                         	20000000	      1067 ns/op	     800 B/op	      11 allocs/op
//BenchmarkMarshallingToJson/01/JsonMarshalStruct-4                      	50000000	       345 ns/op	     208 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/01/JsonMarshalCustom-4                      	30000000	       560 ns/op	     344 B/op	       5 allocs/op
//BenchmarkMarshallingToJson/01/JsonMarshalManual-4                      	300000000	        56.8 ns/op	      64 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/01/JsonMarshalUsingMap-4                    	20000000	      1028 ns/op	     800 B/op	      11 allocs/op
//BenchmarkMarshallingToJson/03/JsonUnmarshalMap-4                       	 5000000	      2753 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkMarshallingToJson/03/JsonUnmarshalStruct-4                    	10000000	      1827 ns/op	     472 B/op	       7 allocs/op
//BenchmarkMarshallingToJson/03/JsonUnmarshalCustom-4                    	20000000	      1065 ns/op	     480 B/op	       8 allocs/op
//BenchmarkMarshallingToJson/03/JsonUnmarshalManual-4                    	100000000	       185 ns/op	     176 B/op	       5 allocs/op
//BenchmarkMarshallingToJson/03/JsonUnmarshalStructUsingMap-4            	 5000000	      2799 ns/op	    1072 B/op	      24 allocs/op
//BenchmarkMarshallingToJson/03/JsonMarshalMap-4                         	10000000	      2232 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkMarshallingToJson/03/JsonMarshalStruct-4                      	20000000	       762 ns/op	     400 B/op	       3 allocs/op
//BenchmarkMarshallingToJson/03/JsonMarshalCustom-4                      	20000000	      1039 ns/op	     520 B/op	       6 allocs/op
//BenchmarkMarshallingToJson/03/JsonMarshalManual-4                      	200000000	        92.4 ns/op	     160 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/03/JsonMarshalUsingMap-4                    	10000000	      2268 ns/op	    1312 B/op	      20 allocs/op
//BenchmarkMarshallingToJson/05/JsonUnmarshalMap-4                       	 3000000	      4284 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkMarshallingToJson/05/JsonUnmarshalStruct-4                    	 5000000	      2765 ns/op	     584 B/op	       9 allocs/op
//BenchmarkMarshallingToJson/05/JsonUnmarshalCustom-4                    	10000000	      1470 ns/op	     592 B/op	      10 allocs/op
//BenchmarkMarshallingToJson/05/JsonUnmarshalManual-4                    	50000000	       267 ns/op	     288 B/op	       7 allocs/op
//BenchmarkMarshallingToJson/05/JsonUnmarshalStructUsingMap-4            	 3000000	      4353 ns/op	    1360 B/op	      36 allocs/op
//BenchmarkMarshallingToJson/05/JsonMarshalMap-4                         	 5000000	      3136 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkMarshallingToJson/05/JsonMarshalStruct-4                      	20000000	       950 ns/op	     448 B/op	       3 allocs/op
//BenchmarkMarshallingToJson/05/JsonMarshalCustom-4                      	10000000	      1386 ns/op	     664 B/op	       6 allocs/op
//BenchmarkMarshallingToJson/05/JsonMarshalManual-4                      	100000000	       117 ns/op	     256 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/05/JsonMarshalUsingMap-4                    	 5000000	      3142 ns/op	    1584 B/op	      26 allocs/op
//BenchmarkMarshallingToJson/10/JsonUnmarshalMap-4                       	 2000000	      8571 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkMarshallingToJson/10/JsonUnmarshalStruct-4                    	 3000000	      5185 ns/op	     872 B/op	      14 allocs/op
//BenchmarkMarshallingToJson/10/JsonUnmarshalCustom-4                    	 5000000	      2487 ns/op	     880 B/op	      15 allocs/op
//BenchmarkMarshallingToJson/10/JsonUnmarshalManual-4                    	30000000	       467 ns/op	     576 B/op	      12 allocs/op
//BenchmarkMarshallingToJson/10/JsonUnmarshalStructUsingMap-4            	 2000000	      8774 ns/op	    2670 B/op	      67 allocs/op
//BenchmarkMarshallingToJson/10/JsonMarshalMap-4                         	 2000000	      6345 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkMarshallingToJson/10/JsonMarshalStruct-4                      	10000000	      1655 ns/op	     896 B/op	       4 allocs/op
//BenchmarkMarshallingToJson/10/JsonMarshalCustom-4                      	10000000	      2229 ns/op	    1048 B/op	       6 allocs/op
//BenchmarkMarshallingToJson/10/JsonMarshalManual-4                      	50000000	       202 ns/op	     512 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/10/JsonMarshalUsingMap-4                    	 2000000	      6262 ns/op	    2886 B/op	      42 allocs/op
//BenchmarkMarshallingToJson/17/JsonUnmarshalMap-4                       	 1000000	     15451 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkMarshallingToJson/17/JsonUnmarshalStruct-4                    	 2000000	      9333 ns/op	    1304 B/op	      21 allocs/op
//BenchmarkMarshallingToJson/17/JsonUnmarshalCustom-4                    	 3000000	      4104 ns/op	    1312 B/op	      22 allocs/op
//BenchmarkMarshallingToJson/17/JsonUnmarshalManual-4                    	20000000	       794 ns/op	    1008 B/op	      19 allocs/op
//BenchmarkMarshallingToJson/17/JsonUnmarshalStructUsingMap-4            	 1000000	     15854 ns/op	    4955 B/op	     110 allocs/op
//BenchmarkMarshallingToJson/17/JsonMarshalMap-4                         	 1000000	     11117 ns/op	    5102 B/op	      64 allocs/op
//BenchmarkMarshallingToJson/17/JsonMarshalStruct-4                      	 5000000	      2715 ns/op	    1728 B/op	       5 allocs/op
//BenchmarkMarshallingToJson/17/JsonMarshalCustom-4                      	 5000000	      3708 ns/op	    1624 B/op	       6 allocs/op
//BenchmarkMarshallingToJson/17/JsonMarshalManual-4                      	50000000	       315 ns/op	     896 B/op	       2 allocs/op
//BenchmarkMarshallingToJson/17/JsonMarshalUsingMap-4                    	 1000000	     10993 ns/op	    5102 B/op	      64 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/marshalling	1481.594s


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

func BenchmarkMarshallingToMaps(b *testing.B) {
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
	})
}


func BenchmarkMarshallingToJson(b *testing.B) {
	b.Run("01", func(b *testing.B) {
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

		b.Run("JsonUnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom01 = &structsCustom.Struct01{}
				tmpStructCustom01.UnmarshalJSON([]byte(sampleJsonString01))
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

		b.Run("JsonMarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = sampleStructCustom01.MarshalJSON()
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

		b.Run("JsonUnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom03 = &structsCustom.Struct03{}
				tmpStructCustom03.UnmarshalJSON([]byte(sampleJsonString03))
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

		b.Run("JsonMarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = sampleStructCustom03.MarshalJSON()
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

		b.Run("JsonUnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom05 = &structsCustom.Struct05{}
				tmpStructCustom05.UnmarshalJSON([]byte(sampleJsonString05))
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

		b.Run("JsonMarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = sampleStructCustom05.MarshalJSON()
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

		b.Run("JsonUnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom10 = &structsCustom.Struct10{}
				tmpStructCustom10.UnmarshalJSON([]byte(sampleJsonString10))
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

		b.Run("JsonMarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = sampleStructCustom10.MarshalJSON()
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

		b.Run("JsonUnmarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStructCustom17 = &structsCustom.Struct17{}
				tmpStructCustom17.UnmarshalJSON([]byte(sampleJsonString17))
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

		b.Run("JsonMarshalManual", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpJsonBytes, _ = sampleStructCustom17.MarshalJSON()
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
