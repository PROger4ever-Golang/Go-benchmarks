package collections

import (
	"github.com/PROger4ever-Golang/Go-benchmarks/common"
	"reflect"
	"strconv"
	"testing"
)

//approximate result
//$ go test -benchmem -bench=^BenchmarkCollections$ -benchtime=10s --timeout 99999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/collections
//BenchmarkCollections/Slice/1/Set-4             	3000000000	         5.33 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1/Get-4             	3000000000	         5.69 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1/Miss-4            	2000000000	         6.85 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1/ForEach-4         	2000000000	         8.03 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3/Set-4             	3000000000	         5.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3/Get-4             	3000000000	         5.83 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3/Miss-4            	2000000000	         6.49 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3/ForEach-4         	1000000000	        12.9 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5/Set-4             	2000000000	         5.93 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5/Get-4             	2000000000	         6.30 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5/Miss-4            	2000000000	         6.88 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5/ForEach-4         	1000000000	        18.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10/Set-4            	2000000000	         6.99 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10/Get-4            	2000000000	         7.13 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10/Miss-4           	2000000000	         8.48 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10/ForEach-4        	500000000	        30.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/30/Set-4            	1000000000	        12.2 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/30/Get-4            	1000000000	        12.7 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/30/Miss-4           	1000000000	        15.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/30/ForEach-4        	200000000	        80.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/30/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/50/Set-4            	1000000000	        20.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/50/Get-4            	1000000000	        21.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/50/Miss-4           	1000000000	        20.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/50/ForEach-4        	100000000	       136 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/50/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/100/Set-4           	500000000	        33.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/100/Get-4           	500000000	        34.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/100/Miss-4          	300000000	        56.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/100/ForEach-4       	50000000	       264 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/100/Delete-4        	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/300/Set-4           	200000000	        65.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/300/Get-4           	200000000	        67.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/300/Miss-4          	100000000	       117 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/300/ForEach-4       	20000000	       767 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/300/Delete-4        	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/500/Set-4           	200000000	        95.9 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/500/Get-4           	200000000	        97.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/500/Miss-4          	100000000	       176 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/500/ForEach-4       	10000000	      1265 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/500/Delete-4        	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1000/Set-4          	100000000	       171 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1000/Get-4          	100000000	       175 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1000/Miss-4         	50000000	       325 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1000/ForEach-4      	 5000000	      2523 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/1000/Delete-4       	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3000/Set-4          	30000000	       468 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3000/Get-4          	30000000	       471 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3000/Miss-4         	20000000	       924 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3000/ForEach-4      	 2000000	      7601 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/3000/Delete-4       	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5000/Set-4          	20000000	       766 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5000/Get-4          	20000000	       770 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5000/Miss-4         	10000000	      1529 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5000/ForEach-4      	 1000000	     12697 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/5000/Delete-4       	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10000/Set-4         	10000000	      1510 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10000/Get-4         	10000000	      1520 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10000/Miss-4        	 5000000	      3031 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10000/ForEach-4     	  500000	     25393 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Slice/10000/Delete-4      	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1/Set-4               	1000000000	        12.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1/Get-4               	2000000000	         7.65 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1/Miss-4              	2000000000	        10.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1/ForEach-4           	300000000	        55.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1/Delete-4            	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3/Set-4               	1000000000	        13.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3/Get-4               	2000000000	         8.05 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3/Miss-4              	2000000000	        10.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3/ForEach-4           	200000000	        77.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3/Delete-4            	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5/Set-4               	1000000000	        14.9 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5/Get-4               	2000000000	         8.55 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5/Miss-4              	2000000000	        10.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5/ForEach-4           	200000000	        94.7 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5/Delete-4            	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10/Set-4              	1000000000	        14.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10/Get-4              	1000000000	        12.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10/Miss-4             	1000000000	        14.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10/ForEach-4          	100000000	       165 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10/Delete-4           	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/30/Set-4              	1000000000	        14.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/30/Get-4              	1000000000	        12.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/30/Miss-4             	1000000000	        14.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/30/ForEach-4          	30000000	       472 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/30/Delete-4           	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/50/Set-4              	1000000000	        17.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/50/Get-4              	1000000000	        13.7 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/50/Miss-4             	1000000000	        17.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/50/ForEach-4          	20000000	       723 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/50/Delete-4           	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/100/Set-4             	1000000000	        18.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/100/Get-4             	1000000000	        14.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/100/Miss-4            	1000000000	        19.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/100/ForEach-4         	10000000	      1450 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/100/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/300/Set-4             	1000000000	        22.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/300/Get-4             	1000000000	        16.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/300/Miss-4            	1000000000	        14.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/300/ForEach-4         	 3000000	      4474 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/300/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/500/Set-4             	500000000	        24.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/500/Get-4             	1000000000	        18.2 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/500/Miss-4            	1000000000	        15.2 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/500/ForEach-4         	 2000000	      7970 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/500/Delete-4          	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1000/Set-4            	500000000	        25.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1000/Get-4            	1000000000	        22.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1000/Miss-4           	1000000000	        15.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1000/ForEach-4        	 1000000	     15703 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/1000/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3000/Set-4            	500000000	        29.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3000/Get-4            	500000000	        27.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3000/Miss-4           	1000000000	        18.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3000/ForEach-4        	  300000	     42456 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/3000/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5000/Set-4            	500000000	        30.1 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5000/Get-4            	500000000	        28.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5000/Miss-4           	1000000000	        18.3 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5000/ForEach-4        	  200000	     74823 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/5000/Delete-4         	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10000/Set-4           	500000000	        31.5 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10000/Get-4           	500000000	        33.8 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10000/Miss-4          	1000000000	        20.4 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10000/ForEach-4       	  100000	    151488 ns/op	       0 B/op	       0 allocs/op
//BenchmarkCollections/Map/10000/Delete-4        	10000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	github.com/PROger4ever-Golang/Go-benchmarks/collections	1848.499s



var collections = []Collection{
	NewSlice(),
	NewMap(),
}

func forEachFunction(key int, value int) {
	tmpInt1 = key
	tmpInt2 = value
}

var tmpInt1, tmpInt2 int

var maxIndices = []int{1, 3, 5, 10, 30, 50, 100, 300, 500, 1000, 3000, 5000, 10000}

func BenchmarkCollections(b *testing.B) {
	for _, collection := range collections {
		b.Run(reflect.TypeOf(collection).Elem().Name(), func(b *testing.B) {
			for _, maxIndex := range maxIndices {
				b.Run(strconv.Itoa(maxIndex), func(b *testing.B) {
					b.Run("Set", func(b *testing.B) {
						collection.Clear()
						ci := common.NewCyclicIterator(b.N, maxIndex)

						for ci.Next() {
							collection.Set(ci.Index, ci.Count)
						}
					})

					b.Run("Get", func(b *testing.B) {
						ci := common.NewCyclicIterator(b.N, maxIndex)

						for ci.Next() {
							tmpInt1, _ = collection.Get(ci.Index)
						}
					})

					b.Run("Miss", func(b *testing.B) {
						from := maxIndex*2 + 1
						to := from + b.N
						for n := from; n < to; n++ {
							tmpInt1, _ = collection.Get(n)
						}
					})

					b.Run("ForEach", func(b *testing.B) {
						for n := 0; n < b.N; n++ {
							collection.ForEach(forEachFunction)
						}
					})

					b.Run("Delete", func(b *testing.B) {
						for n := 0; n < maxIndex; n++ {
							collection.Delete(n)
						}
					})
				})
			}
		})
	}
}
