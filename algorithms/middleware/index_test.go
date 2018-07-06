package middleware

import (
	"testing"
	"github.com/PROger4ever-Golang/Go-benchmarks/common/samples/structsInt"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkMiddleware$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/algorithms/middleware
//BenchmarkMiddleware/01/Direct-4            	10000000000	         1.57 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/01/Functions-4         	10000000000	         1.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/01/Middleware-4        	10000000000	         2.85 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/03/Direct-4            	3000000000	         4.76 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/03/Functions-4         	3000000000	         4.80 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/03/Middleware-4        	2000000000	         9.18 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/05/Direct-4            	2000000000	         8.09 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/05/Functions-4         	2000000000	         8.04 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/05/Middleware-4        	1000000000	        14.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/10/Direct-4            	1000000000	        16.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/10/Functions-4         	1000000000	        16.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/10/Middleware-4        	500000000	        30.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/17/Direct-4            	500000000	        27.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/17/Functions-4         	500000000	        27.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMiddleware/17/Middleware-4        	300000000	        55.5 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/algorithms/middleware	268.780s


type middleware func(mw []middleware, context *structsInt.Struct01)

func mwFunction(mws []middleware, context *structsInt.Struct01) {
	context.Field01++

	if len(mws) == 0 {
		return
	}
	mws[0](mws[1:], context)
}

func Function(context *structsInt.Struct01) {
	context.Field01++
}

var mws01 []middleware

var mws03 = []middleware{mwFunction, mwFunction}

var mws05 = []middleware{mwFunction, mwFunction, mwFunction, mwFunction}

var mws10 = []middleware{mwFunction, mwFunction, mwFunction, mwFunction,
	mwFunction, mwFunction, mwFunction, mwFunction, mwFunction}

var mws17 = []middleware{mwFunction, mwFunction, mwFunction, mwFunction,
	mwFunction, mwFunction, mwFunction, mwFunction, mwFunction,
	mwFunction, mwFunction, mwFunction, mwFunction, mwFunction, mwFunction, mwFunction}

var context = &structsInt.Struct01{}

func BenchmarkMiddleware(b *testing.B) {
	b.Run("01", func(b *testing.B) {
		b.Run("Direct", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				context.Field01++ //01
			}
		})

		b.Run("Functions", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				Function(context) //01
			}
		})

		b.Run("Middleware", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				mwFunction(mws01, context)
			}
		})
	})

	b.Run("03", func(b *testing.B) {
		b.Run("Direct", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				context.Field01++ //01
				context.Field01++ //02
				context.Field01++ //03
			}
		})

		b.Run("Functions", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				Function(context) //01
				Function(context) //02
				Function(context) //03
			}
		})

		b.Run("Middleware", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				mwFunction(mws03, context)
			}
		})
	})

	b.Run("05", func(b *testing.B) {
		b.Run("Direct", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				context.Field01++ //01
				context.Field01++ //02
				context.Field01++ //03
				context.Field01++ //04
				context.Field01++ //05
			}
		})

		b.Run("Functions", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				Function(context) //01
				Function(context) //02
				Function(context) //03
				Function(context) //04
				Function(context) //05
			}
		})

		b.Run("Middleware", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				mwFunction(mws05, context)
			}
		})
	})

	b.Run("10", func(b *testing.B) {
		b.Run("Direct", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				context.Field01++ //01
				context.Field01++ //02
				context.Field01++ //03
				context.Field01++ //04
				context.Field01++ //05
				context.Field01++ //06
				context.Field01++ //07
				context.Field01++ //08
				context.Field01++ //09
				context.Field01++ //10
			}
		})

		b.Run("Functions", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				Function(context) //01
				Function(context) //02
				Function(context) //03
				Function(context) //04
				Function(context) //05
				Function(context) //06
				Function(context) //07
				Function(context) //08
				Function(context) //09
				Function(context) //10
			}
		})

		b.Run("Middleware", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				mwFunction(mws10, context)
			}
		})
	})
	
	b.Run("17", func(b *testing.B) {
		b.Run("Direct", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				context.Field01++ //01
				context.Field01++ //02
				context.Field01++ //03
				context.Field01++ //04
				context.Field01++ //05
				context.Field01++ //06
				context.Field01++ //07
				context.Field01++ //08
				context.Field01++ //09
				context.Field01++ //10
				context.Field01++ //11
				context.Field01++ //12
				context.Field01++ //13
				context.Field01++ //14
				context.Field01++ //15
				context.Field01++ //16
				context.Field01++ //17
			}
		})

		b.Run("Functions", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				Function(context) //01
				Function(context) //02
				Function(context) //03
				Function(context) //04
				Function(context) //05
				Function(context) //06
				Function(context) //07
				Function(context) //08
				Function(context) //09
				Function(context) //10
				Function(context) //11
				Function(context) //12
				Function(context) //13
				Function(context) //14
				Function(context) //15
				Function(context) //16
				Function(context) //17
			}
		})

		b.Run("Middleware", func(b *testing.B) {
			context.Field01 = 0
			for n := 0; n < b.N; n++ {
				mwFunction(mws17, context)
			}
		})
	})
}
