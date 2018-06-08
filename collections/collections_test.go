package collections

import (
	"github.com/PROger4ever-Golang/Go-benchmarks/common"
	"reflect"
	"strconv"
	"testing"
)

//approximate result
//$ go test -bench=BenchmarkCollections -benchtime=10s --timeout 999999s ./...
//goos: linux
//goarch: amd64
//pkg: github.com/PROger4ever-Golang/Go-benchmarks/collections
//BenchmarkCollections/Slice/1/Set-4   	3000000000	         5.41 ns/op
//BenchmarkCollections/Slice/1/Get-4   	3000000000	         5.79 ns/op
//BenchmarkCollections/Slice/1/Miss-4  	2000000000	         7.01 ns/op
//BenchmarkCollections/Slice/1/ForEach-4         	2000000000	         8.16 ns/op
//BenchmarkCollections/Slice/1/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/3/Set-4             	3000000000	         5.62 ns/op
//BenchmarkCollections/Slice/3/Get-4             	3000000000	         5.91 ns/op
//BenchmarkCollections/Slice/3/Miss-4            	2000000000	         6.61 ns/op
//BenchmarkCollections/Slice/3/ForEach-4         	1000000000	        13.1 ns/op
//BenchmarkCollections/Slice/3/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/5/Set-4             	2000000000	         6.00 ns/op
//BenchmarkCollections/Slice/5/Get-4             	2000000000	         6.37 ns/op
//BenchmarkCollections/Slice/5/Miss-4            	2000000000	         7.01 ns/op
//BenchmarkCollections/Slice/5/ForEach-4         	1000000000	        18.1 ns/op
//BenchmarkCollections/Slice/5/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/10/Set-4            	2000000000	         7.05 ns/op
//BenchmarkCollections/Slice/10/Get-4            	2000000000	         7.20 ns/op
//BenchmarkCollections/Slice/10/Miss-4           	2000000000	         8.57 ns/op
//BenchmarkCollections/Slice/10/ForEach-4        	500000000	        30.8 ns/op
//BenchmarkCollections/Slice/10/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/30/Set-4            	1000000000	        12.6 ns/op
//BenchmarkCollections/Slice/30/Get-4            	1000000000	        12.7 ns/op
//BenchmarkCollections/Slice/30/Miss-4           	1000000000	        15.1 ns/op
//BenchmarkCollections/Slice/30/ForEach-4        	200000000	        81.4 ns/op
//BenchmarkCollections/Slice/30/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/50/Set-4            	1000000000	        20.2 ns/op
//BenchmarkCollections/Slice/50/Get-4            	1000000000	        21.7 ns/op
//BenchmarkCollections/Slice/50/Miss-4           	1000000000	        21.2 ns/op
//BenchmarkCollections/Slice/50/ForEach-4        	100000000	       138 ns/op
//BenchmarkCollections/Slice/50/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/100/Set-4           	500000000	        33.8 ns/op
//BenchmarkCollections/Slice/100/Get-4           	500000000	        34.9 ns/op
//BenchmarkCollections/Slice/100/Miss-4          	300000000	        56.0 ns/op
//BenchmarkCollections/Slice/100/ForEach-4       	50000000	       264 ns/op
//BenchmarkCollections/Slice/100/Delete-4        	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/300/Set-4           	200000000	        66.4 ns/op
//BenchmarkCollections/Slice/300/Get-4           	200000000	        67.4 ns/op
//BenchmarkCollections/Slice/300/Miss-4          	100000000	       118 ns/op
//BenchmarkCollections/Slice/300/ForEach-4       	20000000	       780 ns/op
//BenchmarkCollections/Slice/300/Delete-4        	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/500/Set-4           	200000000	        96.8 ns/op
//BenchmarkCollections/Slice/500/Get-4           	200000000	        99.3 ns/op
//BenchmarkCollections/Slice/500/Miss-4          	100000000	       178 ns/op
//BenchmarkCollections/Slice/500/ForEach-4       	10000000	      1277 ns/op
//BenchmarkCollections/Slice/500/Delete-4        	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/1000/Set-4          	100000000	       172 ns/op
//BenchmarkCollections/Slice/1000/Get-4          	100000000	       174 ns/op
//BenchmarkCollections/Slice/1000/Miss-4         	50000000	       327 ns/op
//BenchmarkCollections/Slice/1000/ForEach-4      	 5000000	      2541 ns/op
//BenchmarkCollections/Slice/1000/Delete-4       	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/3000/Set-4          	30000000	       471 ns/op
//BenchmarkCollections/Slice/3000/Get-4          	30000000	       472 ns/op
//BenchmarkCollections/Slice/3000/Miss-4         	20000000	       927 ns/op
//BenchmarkCollections/Slice/3000/ForEach-4      	 2000000	      7689 ns/op
//BenchmarkCollections/Slice/3000/Delete-4       	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/5000/Set-4          	20000000	       775 ns/op
//BenchmarkCollections/Slice/5000/Get-4          	20000000	       775 ns/op
//BenchmarkCollections/Slice/5000/Miss-4         	10000000	      1546 ns/op
//BenchmarkCollections/Slice/5000/ForEach-4      	 1000000	     12854 ns/op
//BenchmarkCollections/Slice/5000/Delete-4       	10000000000	         0.00 ns/op
//BenchmarkCollections/Slice/10000/Set-4         	10000000	      1537 ns/op
//BenchmarkCollections/Slice/10000/Get-4         	10000000	      1539 ns/op
//BenchmarkCollections/Slice/10000/Miss-4        	 5000000	      3196 ns/op
//BenchmarkCollections/Slice/10000/ForEach-4     	  500000	     26435 ns/op
//BenchmarkCollections/Slice/10000/Delete-4      	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/1/Set-4               	1000000000	        12.4 ns/op
//BenchmarkCollections/Map/1/Get-4               	2000000000	         7.72 ns/op
//BenchmarkCollections/Map/1/Miss-4              	2000000000	        10.7 ns/op
//BenchmarkCollections/Map/1/ForEach-4           	300000000	        55.5 ns/op
//BenchmarkCollections/Map/1/Delete-4            	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/3/Set-4               	1000000000	        13.7 ns/op
//BenchmarkCollections/Map/3/Get-4               	2000000000	         8.13 ns/op
//BenchmarkCollections/Map/3/Miss-4              	2000000000	        10.7 ns/op
//BenchmarkCollections/Map/3/ForEach-4           	200000000	        77.3 ns/op
//BenchmarkCollections/Map/3/Delete-4            	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/5/Set-4               	1000000000	        15.1 ns/op
//BenchmarkCollections/Map/5/Get-4               	2000000000	         8.62 ns/op
//BenchmarkCollections/Map/5/Miss-4              	2000000000	        10.7 ns/op
//BenchmarkCollections/Map/5/ForEach-4           	200000000	        96.2 ns/op
//BenchmarkCollections/Map/5/Delete-4            	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/10/Set-4              	1000000000	        15.3 ns/op
//BenchmarkCollections/Map/10/Get-4              	1000000000	        13.3 ns/op
//BenchmarkCollections/Map/10/Miss-4             	1000000000	        14.7 ns/op
//BenchmarkCollections/Map/10/ForEach-4          	100000000	       172 ns/op
//BenchmarkCollections/Map/10/Delete-4           	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/30/Set-4              	1000000000	        14.1 ns/op
//BenchmarkCollections/Map/30/Get-4              	1000000000	        12.7 ns/op
//BenchmarkCollections/Map/30/Miss-4             	1000000000	        14.7 ns/op
//BenchmarkCollections/Map/30/ForEach-4          	30000000	       464 ns/op
//BenchmarkCollections/Map/30/Delete-4           	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/50/Set-4              	1000000000	        15.9 ns/op
//BenchmarkCollections/Map/50/Get-4              	1000000000	        13.6 ns/op
//BenchmarkCollections/Map/50/Miss-4             	1000000000	        19.6 ns/op
//BenchmarkCollections/Map/50/ForEach-4          	20000000	       730 ns/op
//BenchmarkCollections/Map/50/Delete-4           	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/100/Set-4             	1000000000	        20.1 ns/op
//BenchmarkCollections/Map/100/Get-4             	1000000000	        14.3 ns/op
//BenchmarkCollections/Map/100/Miss-4            	1000000000	        19.6 ns/op
//BenchmarkCollections/Map/100/ForEach-4         	10000000	      1447 ns/op
//BenchmarkCollections/Map/100/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/300/Set-4             	1000000000	        22.6 ns/op
//BenchmarkCollections/Map/300/Get-4             	1000000000	        15.5 ns/op
//BenchmarkCollections/Map/300/Miss-4            	1000000000	        15.6 ns/op
//BenchmarkCollections/Map/300/ForEach-4         	 3000000	      4523 ns/op
//BenchmarkCollections/Map/300/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/500/Set-4             	500000000	        23.7 ns/op
//BenchmarkCollections/Map/500/Get-4             	1000000000	        18.5 ns/op
//BenchmarkCollections/Map/500/Miss-4            	1000000000	        15.1 ns/op
//BenchmarkCollections/Map/500/ForEach-4         	 2000000	      7913 ns/op
//BenchmarkCollections/Map/500/Delete-4          	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/1000/Set-4            	500000000	        25.8 ns/op
//BenchmarkCollections/Map/1000/Get-4            	1000000000	        22.9 ns/op
//BenchmarkCollections/Map/1000/Miss-4           	1000000000	        15.7 ns/op
//BenchmarkCollections/Map/1000/ForEach-4        	 1000000	     15869 ns/op
//BenchmarkCollections/Map/1000/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/3000/Set-4            	500000000	        29.9 ns/op
//BenchmarkCollections/Map/3000/Get-4            	500000000	        28.1 ns/op
//BenchmarkCollections/Map/3000/Miss-4           	1000000000	        19.1 ns/op
//BenchmarkCollections/Map/3000/ForEach-4        	  300000	     42986 ns/op
//BenchmarkCollections/Map/3000/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/5000/Set-4            	500000000	        29.2 ns/op
//BenchmarkCollections/Map/5000/Get-4            	500000000	        27.8 ns/op
//BenchmarkCollections/Map/5000/Miss-4           	1000000000	        18.5 ns/op
//BenchmarkCollections/Map/5000/ForEach-4        	  200000	     75077 ns/op
//BenchmarkCollections/Map/5000/Delete-4         	10000000000	         0.00 ns/op
//BenchmarkCollections/Map/10000/Set-4           	500000000	        31.7 ns/op
//BenchmarkCollections/Map/10000/Get-4           	500000000	        32.0 ns/op
//BenchmarkCollections/Map/10000/Miss-4          	1000000000	        21.5 ns/op
//BenchmarkCollections/Map/10000/ForEach-4       	  100000	    156068 ns/op
//BenchmarkCollections/Map/10000/Delete-4        	10000000000	         0.00 ns/op
//PASS

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
