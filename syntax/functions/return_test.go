package functions

import (
	"testing"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsInt"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkReturn$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/syntax/functions
//BenchmarkReturn/01/Unnamed-4             	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/01/UnusedNamed-4         	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/01/Named-4               	10000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/03/Unnamed-4             	10000000000	         0.90 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/03/UnusedNamed-4         	10000000000	         0.90 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/03/Named-4               	10000000000	         0.90 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/05/Unnamed-4             	10000000000	         1.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/05/UnusedNamed-4         	10000000000	         1.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/05/Named-4               	10000000000	         1.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/10/Unnamed-4             	10000000000	         3.01 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/10/UnusedNamed-4         	5000000000	         3.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/10/Named-4               	5000000000	         3.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/17/Unnamed-4             	3000000000	         5.10 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/17/UnusedNamed-4         	3000000000	         5.10 ns/op	       0 B/op	       0 allocs/op
//BenchmarkReturn/17/Named-4               	1000000000	        12.6 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/syntax/functions	188.457s

var (
	tmpField01, tmpField02, tmpField03, tmpField04, tmpField05, tmpField06, tmpField07, tmpField08, tmpField09 int
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

func returnUnnamed01(n int) int {
	return n + 1
}
func returnUnusedNamed01(n int) (Field01 int) {
	return n + 1
}
func returnNamed01(n int) (Field01 int) {
	Field01 = n + 1
	return
}

func returnUnnamed03(n int) (int, int, int) {
	return n + 1, n + 2, n + 3
}
func returnUnusedNamed03(n int) (Field01, Field02, Field03 int) {
	return n + 1, n + 2, n + 3
}
func returnNamed03(n int) (Field01, Field02, Field03 int) {
	Field01 = n + 1
	Field02 = n + 2
	Field03 = n + 3
	return
}

func returnUnnamed05(n int) (int, int, int, int, int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5
}
func returnUnusedNamed05(n int) (Field01, Field02, Field03, Field04, Field05 int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5
}
func returnNamed05(n int) (Field01, Field02, Field03, Field04, Field05 int) {
	Field01 = n + 1
	Field02 = n + 2
	Field03 = n + 3
	Field04 = n + 4
	Field05 = n + 5
	return
}

func returnUnnamed10(n int) (int, int, int, int, int, int, int, int, int, int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5, n + 6, n + 7, n + 8, n + 9, n + 10
}
func returnUnusedNamed10(n int) (Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09, Field10 int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5, n + 6, n + 7, n + 8, n + 9, n + 10
}
func returnNamed10(n int) (Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09, Field10 int) {
	Field01 = n + 1
	Field02 = n + 2
	Field03 = n + 3
	Field04 = n + 4
	Field05 = n + 5
	Field06 = n + 6
	Field07 = n + 7
	Field08 = n + 8
	Field09 = n + 9
	Field10 = n + 10
	return
}

func returnUnnamed17(n int) (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5,
		n + 6, n + 7, n + 8, n + 9, n + 10,
		n + 11, n + 12, n + 13, n + 14, n + 15, n + 16, n + 17
}
func returnUnusedNamed17(n int) (Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09, Field10,
Field11, Field12, Field13, Field14, Field15, Field16, Field17 int) {
	return n + 1, n + 2, n + 3, n + 4, n + 5,
		n + 6, n + 7, n + 8, n + 9, n + 10,
		n + 11, n + 12, n + 13, n + 14, n + 15, n + 16, n + 17
}
func returnNamed17(n int) (Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09, Field10,
Field11, Field12, Field13, Field14, Field15, Field16, Field17 int) {
	Field01 = n + 1
	Field02 = n + 2
	Field03 = n + 3
	Field04 = n + 4
	Field05 = n + 5
	Field06 = n + 6
	Field07 = n + 7
	Field08 = n + 8
	Field09 = n + 9
	Field10 = n + 10
	Field11 = n + 11
	Field12 = n + 12
	Field13 = n + 13
	Field14 = n + 14
	Field15 = n + 15
	Field16 = n + 16
	Field17 = n + 17
	return
}

func BenchmarkReturn(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("Unnamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = returnUnnamed01(n)
			}
		})

		b.Run("UnusedNamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = returnUnusedNamed01(n)
			}
		})

		b.Run("Named", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01 = returnNamed01(n)
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("Unnamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03 = returnUnnamed03(n)
			}
		})

		b.Run("UnusedNamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03 = returnUnusedNamed03(n)
			}
		})

		b.Run("Named", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03 = returnNamed03(n)
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("Unnamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05 = returnUnnamed05(n)
			}
		})

		b.Run("UnusedNamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05 = returnUnusedNamed05(n)
			}
		})

		b.Run("Named", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05 = returnNamed05(n)
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("Unnamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10 = returnUnnamed10(n)
			}
		})

		b.Run("UnusedNamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10 = returnUnusedNamed10(n)
			}
		})

		b.Run("Named", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10 = returnNamed10(n)
			}
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("Unnamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10,
					tmpField11, tmpField12, tmpField13, tmpField14, tmpField15,
					tmpField16, tmpField17 = returnUnnamed17(n)
			}
		})

		b.Run("UnusedNamed", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10,
					tmpField11, tmpField12, tmpField13, tmpField14, tmpField15,
					tmpField16, tmpField17 = returnUnusedNamed17(n)
			}
		})

		b.Run("Named", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tmpField01, tmpField02, tmpField03, tmpField04, tmpField05,
					tmpField06, tmpField07, tmpField08, tmpField09, tmpField10,
					tmpField11, tmpField12, tmpField13, tmpField14, tmpField15,
					tmpField16, tmpField17 = returnNamed17(n)
			}
		})
	})
}
