package structs

import (
	"testing"
	"github.com/PROger4ever-Golang/Go-benchmarks/structs/structsMethods"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkStruct.*$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/structs
//BenchmarkStructAllocating/01/Package-4                    	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/01/Object-4                     	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/01/Pointer-4                    	1000000000	        15.7 ns/op	       8 B/op	       1 allocs/op
//BenchmarkStructAllocating/01/New-4                        	1000000000	        15.6 ns/op	       8 B/op	       1 allocs/op
//BenchmarkStructAllocating/03/Package-4                    	10000000000	         0.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/03/Object-4                     	10000000000	         0.88 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/03/ObjectPointer-4              	500000000	        35.8 ns/op	      32 B/op	       1 allocs/op
//BenchmarkStructAllocating/03/New-4                        	500000000	        34.8 ns/op	      32 B/op	       1 allocs/op
//BenchmarkStructAllocating/05/Package-4                    	10000000000	         1.47 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/05/Object-4                     	2000000000	         7.69 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/05/ObjectPointer-4              	500000000	        34.7 ns/op	      48 B/op	       1 allocs/op
//BenchmarkStructAllocating/05/New-4                        	500000000	        39.7 ns/op	      48 B/op	       1 allocs/op
//BenchmarkStructAllocating/10/Package-4                    	5000000000	         2.95 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/10/Object-4                     	2000000000	        10.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/10/ObjectPointer-4              	500000000	        38.1 ns/op	      80 B/op	       1 allocs/op
//BenchmarkStructAllocating/10/New-4                        	500000000	        39.2 ns/op	      80 B/op	       1 allocs/op
//BenchmarkStructAllocating/17/Package-4                    	3000000000	         5.01 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/17/Object-4                     	1000000000	        19.2 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructAllocating/17/ObjectPointer-4              	300000000	        45.2 ns/op	     144 B/op	       1 allocs/op
//BenchmarkStructAllocating/17/New-4                        	300000000	        45.9 ns/op	     144 B/op	       1 allocs/op
//BenchmarkStructUsage/01/Fields/Package-4                  	10000000000	         1.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Fields/Object-4                   	10000000000	         1.49 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Fields/ObjectPointer-4            	10000000000	         1.55 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Call/Direct-4                     	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Call/Function-4                   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Call/Object-4                     	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Call/ObjectPointer-4              	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Parameter/Primitive-4              	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Parameter/Object-4                	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Parameter/ObjectPointer-4         	10000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Return/Function-4                 	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Return/Object-4                   	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/01/Return/ObjectPointer-4            	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Fields/Package-4                  	5000000000	         3.06 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Fields/Object-4                   	10000000000	         1.48 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Fields/ObjectPointer-4            	10000000000	         1.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Call/Direct-4                     	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Call/Function-4                   	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Call/Object-4                     	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Call/ObjectPointer-4              	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Parameter/Primitive-4              	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Parameter/Object-4                	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Parameter/ObjectPointer-4         	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Return/Function-4                 	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Return/Object-4                   	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/03/Return/ObjectPointer-4            	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Fields/Package-4                  	10000000000	         2.98 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Fields/Object-4                   	10000000000	         2.10 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Fields/ObjectPointer-4            	10000000000	         2.31 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Call/Direct-4                     	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Call/Function-4                   	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Call/Object-4                     	10000000000	         1.93 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Call/ObjectPointer-4              	10000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Parameter/Primitive-4              	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Parameter/Object-4                	5000000000	         3.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Parameter/ObjectPointer-4         	10000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Return/Function-4                 	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Return/Object-4                   	2000000000	         7.95 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/05/Return/ObjectPointer-4            	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Fields/Package-4                  	5000000000	         3.29 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Fields/Object-4                   	10000000000	         2.97 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Fields/ObjectPointer-4            	3000000000	         4.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Call/Direct-4                     	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Call/Function-4                   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Call/Object-4                     	10000000000	         2.71 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Call/ObjectPointer-4              	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Parameter/Primitive-4              	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Parameter/Object-4                	3000000000	         5.04 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Parameter/ObjectPointer-4         	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Return/Function-4                 	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Return/Object-4                   	1000000000	        12.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/10/Return/ObjectPointer-4            	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Fields/Package-4                  	3000000000	         5.04 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Fields/Object-4                   	3000000000	         5.04 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Fields/ObjectPointer-4            	2000000000	         6.95 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Call/Direct-4                     	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Call/Function-4                   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Call/Object-4                     	3000000000	         4.15 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Call/ObjectPointer-4              	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Parameter/Primitive-4              	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Parameter/Object-4                	2000000000	         7.41 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Parameter/ObjectPointer-4         	10000000000	         0.61 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Return/Function-4                 	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Return/Object-4                   	500000000	        28.7 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructUsage/17/Return/ObjectPointer-4            	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/01/ObjectToObject-4     	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/01/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/01/Referencing-4        	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/01/Dereferencing-4      	10000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/03/ObjectToObject-4     	10000000000	         0.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/03/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/03/Referencing-4        	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/03/Dereferencing-4      	10000000000	         0.92 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/05/ObjectToObject-4     	10000000000	         1.64 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/05/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/05/Referencing-4        	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/05/Dereferencing-4      	10000000000	         1.80 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/10/ObjectToObject-4     	10000000000	         2.31 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/10/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/10/Referencing-4        	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/10/Dereferencing-4      	10000000000	         2.67 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/17/ObjectToObject-4     	5000000000	         3.87 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/17/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/17/Referencing-4        	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructPointerDereference/17/Dereferencing-4      	5000000000	         3.89 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/01/ObjectToObject-4     	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/01/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/01/ObjectToPointer-4    	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/01/PointerToObject-4    	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/03/ObjectToObject-4     	10000000000	         0.61 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/03/PointerToPointer-4   	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/03/ObjectToPointer-4    	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/03/PointerToObject-4    	10000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/05/ObjectToObject-4     	10000000000	         1.84 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/05/PointerToPointer-4   	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/05/ObjectToPointer-4    	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/05/PointerToObject-4    	10000000000	         2.07 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/10/ObjectToObject-4     	10000000000	         2.71 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/10/PointerToPointer-4   	10000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/10/ObjectToPointer-4    	10000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/10/PointerToObject-4    	3000000000	         5.06 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/17/ObjectToObject-4     	3000000000	         4.15 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/17/PointerToPointer-4   	10000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/17/ObjectToPointer-4    	10000000000	         0.59 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStructReceiverConversion/17/PointerToObject-4    	3000000000	         4.28 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/structs	1344.904s


var (
	tmpField01, tmpField02, tmpField03, tmpField04, tmpField05, tmpField06, tmpField07, tmpField08, tmpField09 int
	tmpField10, tmpField11, tmpField12, tmpField13, tmpField14, tmpField15, tmpField16, tmpField17             int
)

var tmpObject0100, tmpObject0101 structsMethods.Struct01
var tmpObject0300, tmpObject0301 structsMethods.Struct03
var tmpObject0500, tmpObject0501 structsMethods.Struct05
var tmpObject1000, tmpObject1001 structsMethods.Struct10
var tmpObject1700, tmpObject1701 structsMethods.Struct17

var tmpPointer0100, tmpPointer0101 = structsMethods.NewSampleStruct01(), structsMethods.NewSampleStruct01()
var tmpPointer0300, tmpPointer0301 = structsMethods.NewSampleStruct03(), structsMethods.NewSampleStruct03()
var tmpPointer0500, tmpPointer0501 = structsMethods.NewSampleStruct05(), structsMethods.NewSampleStruct05()
var tmpPointer1000, tmpPointer1001 = structsMethods.NewSampleStruct10(), structsMethods.NewSampleStruct10()
var tmpPointer1700, tmpPointer1701 = structsMethods.NewSampleStruct17(), structsMethods.NewSampleStruct17()

func functionCall() int {
	return tmpField01
}
func functionParameter(p int) int {
	return p
}

func BenchmarkStructAllocating(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("Package", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = n + 1
			}
		})
		b.Run("Object", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0100 = structsMethods.Struct01{
					Field01: n + 1,
				}
			}
		})
		b.Run("Pointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0100 = &structsMethods.Struct01{
					Field01: n + 1,
				}
			}
		})
		b.Run("New", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0100 = new(structsMethods.Struct01)
				tmpPointer0100.Field01 = n + 1
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("Package", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = n + 1
				tmpField02 = n + 2
				tmpField03 = n + 3
			}
		})
		b.Run("Object", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0300 = structsMethods.Struct03{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
				}
			}
		})
		b.Run("ObjectPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0300 = &structsMethods.Struct03{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
				}
			}
		})
		b.Run("New", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0300 = new(structsMethods.Struct03)
				tmpPointer0300.Field01 = n + 1
				tmpPointer0300.Field02 = n + 2
				tmpPointer0300.Field03 = n + 3
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("Package", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = n + 1
				tmpField02 = n + 2
				tmpField03 = n + 3
				tmpField04 = n + 4
				tmpField05 = n + 5
			}
		})
		b.Run("Object", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0500 = structsMethods.Struct05{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
				}
			}
		})
		b.Run("ObjectPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0500 = &structsMethods.Struct05{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
				}
			}
		})
		b.Run("New", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0500 = new(structsMethods.Struct05)
				tmpPointer0500.Field01 = n + 1
				tmpPointer0500.Field02 = n + 2
				tmpPointer0500.Field03 = n + 3
				tmpPointer0500.Field04 = n + 4
				tmpPointer0500.Field05 = n + 5
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("Package", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = n + 1
				tmpField02 = n + 2
				tmpField03 = n + 3
				tmpField04 = n + 4
				tmpField05 = n + 5
				tmpField06 = n + 6
				tmpField07 = n + 7
				tmpField08 = n + 8
				tmpField09 = n + 9
				tmpField10 = n + 10
			}
		})
		b.Run("Object", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1000 = structsMethods.Struct10{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
					Field06: n + 6,
					Field07: n + 7,
					Field08: n + 8,
					Field09: n + 9,
					Field10: n + 10,
				}
			}
		})
		b.Run("ObjectPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1000 = &structsMethods.Struct10{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
					Field06: n + 6,
					Field07: n + 7,
					Field08: n + 8,
					Field09: n + 9,
					Field10: n + 10,
				}
			}
		})
		b.Run("New", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1000 = new(structsMethods.Struct10)
				tmpPointer1000.Field01 = n + 1
				tmpPointer1000.Field02 = n + 2
				tmpPointer1000.Field03 = n + 3
				tmpPointer1000.Field04 = n + 4
				tmpPointer1000.Field05 = n + 5
				tmpPointer1000.Field06 = n + 6
				tmpPointer1000.Field07 = n + 7
				tmpPointer1000.Field08 = n + 8
				tmpPointer1000.Field09 = n + 9
				tmpPointer1000.Field10 = n + 10
			}
		})
	})

	b.Run("17", func(b *testing.B) {
		b.Run("Package", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = n + 1
				tmpField02 = n + 2
				tmpField03 = n + 3
				tmpField04 = n + 4
				tmpField05 = n + 5
				tmpField06 = n + 6
				tmpField07 = n + 7
				tmpField08 = n + 8
				tmpField09 = n + 9
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
		b.Run("Object", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1700 = structsMethods.Struct17{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
					Field06: n + 6,
					Field07: n + 7,
					Field08: n + 8,
					Field09: n + 9,
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
		b.Run("ObjectPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1700 = &structsMethods.Struct17{
					Field01: n + 1,
					Field02: n + 2,
					Field03: n + 3,
					Field04: n + 4,
					Field05: n + 5,
					Field06: n + 6,
					Field07: n + 7,
					Field08: n + 8,
					Field09: n + 9,
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
		b.Run("New", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1700 = new(structsMethods.Struct17)
				tmpPointer1700.Field01 = n + 1
				tmpPointer1700.Field02 = n + 2
				tmpPointer1700.Field03 = n + 3
				tmpPointer1700.Field04 = n + 4
				tmpPointer1700.Field05 = n + 5
				tmpPointer1700.Field06 = n + 6
				tmpPointer1700.Field07 = n + 7
				tmpPointer1700.Field08 = n + 8
				tmpPointer1700.Field09 = n + 9
				tmpPointer1700.Field10 = n + 10
				tmpPointer1700.Field11 = n + 11
				tmpPointer1700.Field12 = n + 12
				tmpPointer1700.Field13 = n + 13
				tmpPointer1700.Field14 = n + 14
				tmpPointer1700.Field15 = n + 15
				tmpPointer1700.Field16 = n + 16
				tmpPointer1700.Field17 = n + 17
			}
		})
	})
}

func BenchmarkStructUsage(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("Fields", func(b *testing.B) {
			b.Run("Package", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = tmpField01 + 1
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0100.Field01 = tmpObject0100.Field01 + 1
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0100.Field01 = tmpPointer0100.Field01 + 1
				}
			})
		})

		b.Run("Call", func(b *testing.B) {
			b.Run("Direct", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = n + 1
				}
			})
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = tmpObject0100.ObjectPrimitiveMethod()
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = tmpPointer0100.PointerPrimitiveMethod()
				}
			})
		})

		b.Run("Parameter", func(b *testing.B) {
			b.Run("Primitive", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = functionParameter(n)
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = tmpObject0100.ObjectParameterMethod(tmpObject0100)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = tmpPointer0100.PointerParameterMethod(tmpPointer0100)
				}
			})
		})

		b.Run("Return", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField01 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0100 = tmpObject0100.ObjectReturnMethod(tmpObject0100)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0100 = tmpPointer0100.PointerReturnMethod(tmpPointer0100)
				}
			})
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("Fields", func(b *testing.B) {
			b.Run("Package", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = tmpField03 + 1
					tmpField02 = tmpField02 + 2
					tmpField03 = tmpField03 + 3
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0300.Field01 = tmpObject0300.Field01 + 1
					tmpObject0300.Field02 = tmpObject0300.Field02 + 2
					tmpObject0300.Field03 = tmpObject0300.Field03 + 3
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0300.Field01 = tmpPointer0300.Field01 + 1
					tmpPointer0300.Field02 = tmpPointer0300.Field02 + 2
					tmpPointer0300.Field03 = tmpPointer0300.Field03 + 3
				}
			})
		})

		b.Run("Call", func(b *testing.B) {
			b.Run("Direct", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = n + 1
				}
			})
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = tmpObject0300.ObjectPrimitiveMethod()
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = tmpPointer0300.PointerPrimitiveMethod()
				}
			})
		})

		b.Run("Parameter", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = functionParameter(n)
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = tmpObject0300.ObjectParameterMethod(tmpObject0300)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = tmpPointer0300.PointerParameterMethod(tmpPointer0300)
				}
			})
		})

		b.Run("Return", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField03 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0300 = tmpObject0300.ObjectReturnMethod(tmpObject0300)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0300 = tmpPointer0300.PointerReturnMethod(tmpPointer0300)
				}
			})
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("Fields", func(b *testing.B) {
			b.Run("Package", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = tmpField05 + 1
					tmpField02 = tmpField02 + 2
					tmpField03 = tmpField03 + 3
					tmpField04 = tmpField04 + 4
					tmpField05 = tmpField05 + 5
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0500.Field01 = tmpObject0500.Field01 + 1
					tmpObject0500.Field02 = tmpObject0500.Field02 + 2
					tmpObject0500.Field03 = tmpObject0500.Field03 + 3
					tmpObject0500.Field04 = tmpObject0500.Field04 + 4
					tmpObject0500.Field05 = tmpObject0500.Field05 + 5
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0500.Field01 = tmpPointer0500.Field01 + 1
					tmpPointer0500.Field02 = tmpPointer0500.Field02 + 2
					tmpPointer0500.Field03 = tmpPointer0500.Field03 + 3
					tmpPointer0500.Field04 = tmpPointer0500.Field04 + 4
					tmpPointer0500.Field05 = tmpPointer0500.Field05 + 5
				}
			})
		})

		b.Run("Call", func(b *testing.B) {
			b.Run("Direct", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = n + 1
				}
			})
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = tmpObject0500.ObjectPrimitiveMethod()
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = tmpPointer0500.PointerPrimitiveMethod()
				}
			})
		})

		b.Run("Parameter", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = functionParameter(n)
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = tmpObject0500.ObjectParameterMethod(tmpObject0500)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = tmpPointer0500.PointerParameterMethod(tmpPointer0500)
				}
			})
		})

		b.Run("Return", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField05 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject0500 = tmpObject0500.ObjectReturnMethod(tmpObject0500)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer0500 = tmpPointer0500.PointerReturnMethod(tmpPointer0500)
				}
			})
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("Fields", func(b *testing.B) {
			b.Run("Package", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = tmpField10 + 1
					tmpField02 = tmpField02 + 2
					tmpField03 = tmpField03 + 3
					tmpField04 = tmpField04 + 4
					tmpField05 = tmpField05 + 5
					tmpField06 = tmpField06 + 6
					tmpField07 = tmpField07 + 7
					tmpField08 = tmpField08 + 8
					tmpField09 = tmpField09 + 9
					tmpField10 = tmpField10 + 10
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject1000.Field01 = tmpObject1000.Field01 + 1
					tmpObject1000.Field02 = tmpObject1000.Field02 + 2
					tmpObject1000.Field03 = tmpObject1000.Field03 + 3
					tmpObject1000.Field04 = tmpObject1000.Field04 + 4
					tmpObject1000.Field05 = tmpObject1000.Field05 + 5
					tmpObject1000.Field06 = tmpObject1000.Field06 + 6
					tmpObject1000.Field07 = tmpObject1000.Field07 + 7
					tmpObject1000.Field08 = tmpObject1000.Field08 + 8
					tmpObject1000.Field09 = tmpObject1000.Field09 + 9
					tmpObject1000.Field10 = tmpObject1000.Field10 + 10
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer1000.Field01 = tmpPointer1000.Field01 + 1
					tmpPointer1000.Field02 = tmpPointer1000.Field02 + 2
					tmpPointer1000.Field03 = tmpPointer1000.Field03 + 3
					tmpPointer1000.Field04 = tmpPointer1000.Field04 + 4
					tmpPointer1000.Field05 = tmpPointer1000.Field05 + 5
					tmpPointer1000.Field06 = tmpPointer1000.Field06 + 6
					tmpPointer1000.Field07 = tmpPointer1000.Field07 + 7
					tmpPointer1000.Field08 = tmpPointer1000.Field08 + 8
					tmpPointer1000.Field09 = tmpPointer1000.Field09 + 9
					tmpPointer1000.Field10 = tmpPointer1000.Field10 + 10
				}
			})
		})

		b.Run("Call", func(b *testing.B) {
			b.Run("Direct", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = n + 1
				}
			})
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = tmpObject1000.ObjectPrimitiveMethod()
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = tmpPointer1000.PointerPrimitiveMethod()
				}
			})
		})

		b.Run("Parameter", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = functionParameter(n)
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = tmpObject1000.ObjectParameterMethod(tmpObject1000)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = tmpPointer1000.PointerParameterMethod(tmpPointer1000)
				}
			})
		})

		b.Run("Return", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField10 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject1000 = tmpObject1000.ObjectReturnMethod(tmpObject1000)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer1000 = tmpPointer1000.PointerReturnMethod(tmpPointer1000)
				}
			})
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("Fields", func(b *testing.B) {
			b.Run("Package", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = tmpField17 + 1
					tmpField02 = tmpField02 + 2
					tmpField03 = tmpField03 + 3
					tmpField04 = tmpField04 + 4
					tmpField05 = tmpField05 + 5
					tmpField06 = tmpField06 + 6
					tmpField07 = tmpField07 + 7
					tmpField08 = tmpField08 + 8
					tmpField09 = tmpField09 + 9
					tmpField10 = tmpField10 + 10
					tmpField11 = tmpField11 + 11
					tmpField12 = tmpField12 + 12
					tmpField13 = tmpField13 + 13
					tmpField14 = tmpField14 + 14
					tmpField15 = tmpField15 + 15
					tmpField16 = tmpField16 + 16
					tmpField17 = tmpField17 + 17
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject1700.Field01 = tmpObject1700.Field01 + 1
					tmpObject1700.Field02 = tmpObject1700.Field02 + 2
					tmpObject1700.Field03 = tmpObject1700.Field03 + 3
					tmpObject1700.Field04 = tmpObject1700.Field04 + 4
					tmpObject1700.Field05 = tmpObject1700.Field05 + 5
					tmpObject1700.Field06 = tmpObject1700.Field06 + 6
					tmpObject1700.Field07 = tmpObject1700.Field07 + 7
					tmpObject1700.Field08 = tmpObject1700.Field08 + 8
					tmpObject1700.Field09 = tmpObject1700.Field09 + 9
					tmpObject1700.Field10 = tmpObject1700.Field10 + 10
					tmpObject1700.Field11 = tmpObject1700.Field11 + 11
					tmpObject1700.Field12 = tmpObject1700.Field12 + 12
					tmpObject1700.Field13 = tmpObject1700.Field13 + 13
					tmpObject1700.Field14 = tmpObject1700.Field14 + 14
					tmpObject1700.Field15 = tmpObject1700.Field15 + 15
					tmpObject1700.Field16 = tmpObject1700.Field16 + 16
					tmpObject1700.Field17 = tmpObject1700.Field17 + 17
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer1700.Field01 = tmpPointer1700.Field01 + 1
					tmpPointer1700.Field02 = tmpPointer1700.Field02 + 2
					tmpPointer1700.Field03 = tmpPointer1700.Field03 + 3
					tmpPointer1700.Field04 = tmpPointer1700.Field04 + 4
					tmpPointer1700.Field05 = tmpPointer1700.Field05 + 5
					tmpPointer1700.Field06 = tmpPointer1700.Field06 + 6
					tmpPointer1700.Field07 = tmpPointer1700.Field07 + 7
					tmpPointer1700.Field08 = tmpPointer1700.Field08 + 8
					tmpPointer1700.Field09 = tmpPointer1700.Field09 + 9
					tmpPointer1700.Field10 = tmpPointer1700.Field10 + 10
					tmpPointer1700.Field11 = tmpPointer1700.Field11 + 11
					tmpPointer1700.Field12 = tmpPointer1700.Field12 + 12
					tmpPointer1700.Field13 = tmpPointer1700.Field13 + 13
					tmpPointer1700.Field14 = tmpPointer1700.Field14 + 14
					tmpPointer1700.Field15 = tmpPointer1700.Field15 + 15
					tmpPointer1700.Field16 = tmpPointer1700.Field16 + 16
					tmpPointer1700.Field17 = tmpPointer1700.Field17 + 17
				}
			})
		})

		b.Run("Call", func(b *testing.B) {
			b.Run("Direct", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = n + 1
				}
			})
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = tmpObject1700.ObjectPrimitiveMethod()
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = tmpPointer1700.PointerPrimitiveMethod()
				}
			})
		})

		b.Run("Parameter", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = functionParameter(n)
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = tmpObject1700.ObjectParameterMethod(tmpObject1700)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = tmpPointer1700.PointerParameterMethod(tmpPointer1700)
				}
			})
		})

		b.Run("Return", func(b *testing.B) {
			b.Run("Function", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpField17 = functionCall()
				}
			})
			b.Run("Object", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpObject1700 = tmpObject1700.ObjectReturnMethod(tmpObject1700)
				}
			})
			b.Run("ObjectPointer", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tmpPointer1700 = tmpPointer1700.PointerReturnMethod(tmpPointer1700)
				}
			})
		})
	})
}

func BenchmarkStructPointerDereference(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0101 = tmpObject0100
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0101 = tmpPointer0100
			}
		})

		b.Run("Referencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0101 = &tmpObject0100
			}
		})

		b.Run("Dereferencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0101 = *tmpPointer0100
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0301 = tmpObject0300
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0301 = tmpPointer0300
			}
		})

		b.Run("Referencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0301 = &tmpObject0300
			}
		})

		b.Run("Dereferencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0301 = *tmpPointer0300
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0501 = tmpObject0500
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0501 = tmpPointer0500
			}
		})

		b.Run("Referencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer0501 = &tmpObject0500
			}
		})

		b.Run("Dereferencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject0501 = *tmpPointer0500
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1001 = tmpObject1000
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1001 = tmpPointer1000
			}
		})

		b.Run("Referencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1001 = &tmpObject1000
			}
		})

		b.Run("Dereferencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1001 = *tmpPointer1000
			}
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1701 = tmpObject1700
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1701 = tmpPointer1700
			}
		})

		b.Run("Referencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpPointer1701 = &tmpObject1700
			}
		})

		b.Run("Dereferencing", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpObject1701 = *tmpPointer1700
			}
		})
	})
}

func BenchmarkStructReceiverConversion(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = tmpObject0100.ObjectPrimitiveMethod()
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = tmpPointer0100.PointerPrimitiveMethod()
			}
		})

		b.Run("ObjectToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = tmpObject0100.PointerPrimitiveMethod()
			}
		})

		b.Run("PointerToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = tmpPointer0100.ObjectPrimitiveMethod()
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField03 = tmpObject0300.ObjectPrimitiveMethod()
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField03 = tmpPointer0300.PointerPrimitiveMethod()
			}
		})

		b.Run("ObjectToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField03 = tmpObject0300.PointerPrimitiveMethod()
			}
		})

		b.Run("PointerToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField03 = tmpPointer0300.ObjectPrimitiveMethod()
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField05 = tmpObject0500.ObjectPrimitiveMethod()
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField05 = tmpPointer0500.PointerPrimitiveMethod()
			}
		})

		b.Run("ObjectToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField05 = tmpObject0500.PointerPrimitiveMethod()
			}
		})

		b.Run("PointerToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField05 = tmpPointer0500.ObjectPrimitiveMethod()
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField10 = tmpObject1000.ObjectPrimitiveMethod()
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField10 = tmpPointer1000.PointerPrimitiveMethod()
			}
		})

		b.Run("ObjectToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField10 = tmpObject1000.PointerPrimitiveMethod()
			}
		})

		b.Run("PointerToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField10 = tmpPointer1000.ObjectPrimitiveMethod()
			}
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("ObjectToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField17 = tmpObject1700.ObjectPrimitiveMethod()
			}
		})

		b.Run("PointerToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField17 = tmpPointer1700.PointerPrimitiveMethod()
			}
		})

		b.Run("ObjectToPointer", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField17 = tmpObject1700.PointerPrimitiveMethod()
			}
		})

		b.Run("PointerToObject", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField17 = tmpPointer1700.ObjectPrimitiveMethod()
			}
		})
	})
}
