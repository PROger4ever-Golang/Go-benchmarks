package bitwise_operators

import (
	"testing"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkBitwiseOperators$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/syntax/bitwise-operators
//BenchmarkBitwiseOperators/DirectComparison-4         	10000000000	         1.41 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/AND-4                      	10000000000	         1.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/OR-4                       	10000000000	         1.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/XOR-4                      	10000000000	         2.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/AND_NOT-4                  	10000000000	         2.66 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/LeftShift-4                	3000000000	         4.13 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBitwiseOperators/RightShift-4               	3000000000	         4.10 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/syntax/bitwise-operators	118.271s

type Signed int

var (
	Signed0 Signed = 0
	Signed1 Signed = 255
)

type Unsigned uint

var (
	Unsigned0 Unsigned = 0
	Unsigned1 Unsigned = 255
)

var (
	tmpBool00, tmpBool01, tmpBool02, tmpBool03 bool
)

func BenchmarkBitwiseOperators(b *testing.B) {
	b.Run("DirectComparison", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Signed0 == Signed0
			tmpBool01 = Signed0 == Signed1
			tmpBool02 = Signed1 == Signed0
			tmpBool03 = Signed1 == Signed1
		}
	})

	b.Run("AND", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Signed0&Signed0 == Signed0
			tmpBool01 = Signed0&Signed1 == Signed0
			tmpBool02 = Signed1&Signed0 == Signed0
			tmpBool03 = Signed1&Signed1 == Signed1
		}
	})

	b.Run("OR", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Signed0|Signed0 == Signed0
			tmpBool01 = Signed0|Signed1 == Signed1
			tmpBool02 = Signed1|Signed0 == Signed1
			tmpBool03 = Signed1|Signed1 == Signed1
		}
	})

	b.Run("XOR", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Signed0^Signed0 == Signed0
			tmpBool01 = Signed0^Signed1 == Signed1
			tmpBool02 = Signed1^Signed0 == Signed1
			tmpBool03 = Signed1^Signed1 == Signed1
		}
	})

	b.Run("AND_NOT", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Signed0&^Signed0 == Signed0
			tmpBool01 = Signed0&^Signed1 == Signed0
			tmpBool02 = Signed1&^Signed0 == Signed1
			tmpBool03 = Signed1&^Signed1 == Signed0
		}
	})

	b.Run("LeftShift", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Unsigned0<<Unsigned0 == Unsigned0
			tmpBool01 = Unsigned0<<Unsigned1 == Unsigned0
			tmpBool02 = Unsigned1<<Unsigned0 == Unsigned1
			tmpBool03 = Unsigned1<<Unsigned1 == Unsigned0
		}
	})

	b.Run("RightShift", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpBool00 = Unsigned0>>Unsigned0 == Unsigned0
			tmpBool01 = Unsigned0>>Unsigned1 == Unsigned0
			tmpBool02 = Unsigned1>>Unsigned0 == Unsigned1
			tmpBool03 = Unsigned1>>Unsigned1 == Unsigned0
		}
	})
}
