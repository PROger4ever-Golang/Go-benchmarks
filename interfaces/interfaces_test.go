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

type ExampleInterface interface {
	Number() int
	SetNumber(number int)
}

type ExampleStruct struct {
	number int
}

func (es *ExampleStruct) Number() int {
	return es.number
}
func (es *ExampleStruct) SetNumber(number int) {
	es.number = number
}

var exampleObject = &ExampleStruct{}
var interfaceObject ExampleInterface
var exampleNumber int

func BenchmarkBoxing(b *testing.B) {
	b.Run("Clear", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			exampleObject.SetNumber(n)
			exampleNumber = exampleObject.Number()
		}
	})

	b.Run("Boxing", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			interfaceObject = exampleObject
			interfaceObject.SetNumber(n)
			exampleObject = interfaceObject.(*ExampleStruct)
			exampleNumber = exampleObject.Number()
		}
	})
}
