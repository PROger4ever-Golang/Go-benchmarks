package interfaces

import (
	"testing"
)

//approximate result
//$ go test -bench=. ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/interfaces
//BenchmarkBoxing/Clear-4         	2000000000	         0.85 ns/op
//BenchmarkBoxing/Boxing-4        	500000000	         3.82 ns/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/interfaces	4.094s

type SampleInterface interface {
	Number() int
	SetNumber(number int)
}

type SampleStruct struct {
	number int
}

func (es *SampleStruct) Number() int {
	return es.number
}
func (es *SampleStruct) SetNumber(number int) {
	es.number = number
}

var tmpObject = &SampleStruct{}
var interfaceObject SampleInterface
var tmpNumber int

func BenchmarkBoxing(b *testing.B) {
	b.Run("Clear", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tmpObject.SetNumber(n)
			tmpNumber = tmpObject.Number()
		}
	})

	b.Run("Boxing", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			interfaceObject = tmpObject
			interfaceObject.SetNumber(n)
			tmpObject = interfaceObject.(*SampleStruct)
			tmpNumber = tmpObject.Number()
		}
	})
}
