package allocation

import (
	"testing"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsInt"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkVariables$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/allocation
//BenchmarkVariables/01/Primitives-4         	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/01/StructObject-4       	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/01/StructPointer-4      	1000000000	        14.3 ns/op	       8 B/op	       1 allocs/op
//BenchmarkVariables/01/StructNew-4          	1000000000	        14.4 ns/op	       8 B/op	       1 allocs/op
//BenchmarkVariables/03/Primitives-4         	10000000000	         0.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/03/StructObject-4       	10000000000	         0.90 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/03/StructPointer-4      	1000000000	        23.2 ns/op	      32 B/op	       1 allocs/op
//BenchmarkVariables/03/StructNew-4          	1000000000	        22.9 ns/op	      32 B/op	       1 allocs/op
//BenchmarkVariables/05/Primitives-4         	10000000000	         1.49 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/05/StructObject-4       	2000000000	         7.75 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/05/StructPointer-4      	500000000	        27.0 ns/op	      48 B/op	       1 allocs/op
//BenchmarkVariables/05/StructNew-4          	500000000	        27.7 ns/op	      48 B/op	       1 allocs/op
//BenchmarkVariables/10/Primitives-4         	10000000000	         2.98 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/10/StructObject-4       	2000000000	        10.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/10/StructPointer-4      	500000000	        34.9 ns/op	      80 B/op	       1 allocs/op
//BenchmarkVariables/10/StructNew-4          	500000000	        37.7 ns/op	      80 B/op	       1 allocs/op
//BenchmarkVariables/17/Primitives-4         	3000000000	         5.07 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/17/StructObject-4       	1000000000	        13.7 ns/op	       0 B/op	       0 allocs/op
//BenchmarkVariables/17/StructPointer-4      	300000000	        48.4 ns/op	     144 B/op	       1 allocs/op
//BenchmarkVariables/17/StructNew-4          	300000000	        49.8 ns/op	     144 B/op	       1 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/allocation	338.856s

var (
	tmpField1, tmpField2, tmpField3, tmpField4, tmpField5, tmpField6, tmpField7, tmpField8, tmpField9 int
	tmpField10, tmpField11, tmpField12, tmpField13, tmpField14, tmpField15, tmpField16, tmpField17    int
)

var tmpStruct01 structsInt.Struct01
var tmpStruct03 structsInt.Struct03
var tmpStruct05 structsInt.Struct05
var tmpStruct10 structsInt.Struct10
var tmpStruct17 structsInt.Struct17

var tmpPointer01 *structsInt.Struct01
var tmpPointer03 *structsInt.Struct03
var tmpPointer05 *structsInt.Struct05
var tmpPointer10 *structsInt.Struct10
var tmpPointer17 *structsInt.Struct17

func BenchmarkVariables(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("Primitives", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField1 = n + 1
			}
		})
		b.Run("StructObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct01 = structsInt.Struct01{
					Field1: n + 1,
				}
			}
		})
		b.Run("StructPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer01 = &structsInt.Struct01{
					Field1: n + 1,
				}
			}
		})
		b.Run("StructNew", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer01 = new(structsInt.Struct01)
				tmpPointer01.Field1 = n + 1
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("Primitives", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField1 = n + 1
				tmpField2 = n + 2
				tmpField3 = n + 3
			}
		})
		b.Run("StructObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct03 = structsInt.Struct03{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
				}
			}
		})
		b.Run("StructPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer03 = &structsInt.Struct03{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
				}
			}
		})
		b.Run("StructNew", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer03 = new(structsInt.Struct03)
				tmpPointer03.Field1 = n + 1
				tmpPointer03.Field2 = n + 2
				tmpPointer03.Field3 = n + 3
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("Primitives", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField1 = n + 1
				tmpField2 = n + 2
				tmpField3 = n + 3
				tmpField4 = n + 4
				tmpField5 = n + 5
			}
		})
		b.Run("StructObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct05 = structsInt.Struct05{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
				}
			}
		})
		b.Run("StructPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer05 = &structsInt.Struct05{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
				}
			}
		})
		b.Run("StructNew", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer05 = new(structsInt.Struct05)
				tmpPointer05.Field1 = n + 1
				tmpPointer05.Field2 = n + 2
				tmpPointer05.Field3 = n + 3
				tmpPointer05.Field4 = n + 4
				tmpPointer05.Field5 = n + 5
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("Primitives", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField1 = n + 1
				tmpField2 = n + 2
				tmpField3 = n + 3
				tmpField4 = n + 4
				tmpField5 = n + 5
				tmpField6 = n + 6
				tmpField7 = n + 7
				tmpField8 = n + 8
				tmpField9 = n + 9
				tmpField10 = n + 10
			}
		})
		b.Run("StructObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct10 = structsInt.Struct10{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
					Field6: n + 6,
					Field7: n + 7,
					Field8: n + 8,
					Field9: n + 9,
					Field10: n + 10,
				}
			}
		})
		b.Run("StructPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer10 = &structsInt.Struct10{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
					Field6: n + 6,
					Field7: n + 7,
					Field8: n + 8,
					Field9: n + 9,
					Field10: n + 10,
				}
			}
		})
		b.Run("StructNew", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer10 = new(structsInt.Struct10)
				tmpPointer10.Field1 = n + 1
				tmpPointer10.Field2 = n + 2
				tmpPointer10.Field3 = n + 3
				tmpPointer10.Field4 = n + 4
				tmpPointer10.Field5 = n + 5
				tmpPointer10.Field6 = n + 6
				tmpPointer10.Field7 = n + 7
				tmpPointer10.Field8 = n + 8
				tmpPointer10.Field9 = n + 9
				tmpPointer10.Field10 = n + 10
			}
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("Primitives", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField1 = n + 1
				tmpField2 = n + 2
				tmpField3 = n + 3
				tmpField4 = n + 4
				tmpField5 = n + 5
				tmpField6 = n + 6
				tmpField7 = n + 7
				tmpField8 = n + 8
				tmpField9 = n + 9
				tmpField10 = n + 10
				tmpField11 = n + 11
				tmpField12 = n + 12
				tmpField13 = n + 13
				tmpField14 = n + 14
				tmpField15 = n + 15
				tmpField16 = n + 16
				tmpField17 = n + 17
			}
		})
		b.Run("StructObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpStruct17 = structsInt.Struct17{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
					Field6: n + 6,
					Field7: n + 7,
					Field8: n + 8,
					Field9: n + 9,
					Field10: n + 10,
					Field11: n + 11,
					Field12: n + 12,
					Field13: n + 13,
					Field14: n + 14,
					Field15: n + 15,
					Field16: n + 16,
					Field17: n + 17,
				}
			}
		})
		b.Run("StructPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer17 = &structsInt.Struct17{
					Field1: n + 1,
					Field2: n + 2,
					Field3: n + 3,
					Field4: n + 4,
					Field5: n + 5,
					Field6: n + 6,
					Field7: n + 7,
					Field8: n + 8,
					Field9: n + 9,
					Field10: n + 10,
					Field11: n + 11,
					Field12: n + 12,
					Field13: n + 13,
					Field14: n + 14,
					Field15: n + 15,
					Field16: n + 16,
					Field17: n + 17,
				}
			}
		})
		b.Run("StructNew", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer17 = new(structsInt.Struct17)
				tmpPointer17.Field1 = n + 1
				tmpPointer17.Field2 = n + 2
				tmpPointer17.Field3 = n + 3
				tmpPointer17.Field4 = n + 4
				tmpPointer17.Field5 = n + 5
				tmpPointer17.Field6 = n + 6
				tmpPointer17.Field7 = n + 7
				tmpPointer17.Field8 = n + 8
				tmpPointer17.Field9 = n + 9
				tmpPointer17.Field10 = n + 10
				tmpPointer17.Field11 = n + 11
				tmpPointer17.Field12 = n + 12
				tmpPointer17.Field13 = n + 13
				tmpPointer17.Field14 = n + 14
				tmpPointer17.Field15 = n + 15
				tmpPointer17.Field16 = n + 16
				tmpPointer17.Field17 = n + 17
			}
		})
	})
}
