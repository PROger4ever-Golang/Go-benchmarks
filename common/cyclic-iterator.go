package common

type CyclicIterator struct {
	Count int
	Index int

	MaxCount int
	MinIndex int
	MaxIndex int
}

func (ci *CyclicIterator) Clear() {
	ci.Index = ci.MinIndex - 1
	ci.Count = -1
}

func (ci *CyclicIterator) Next() bool {
	ci.Count++
	if ci.Index >= ci.MaxIndex {
		ci.Index = ci.MinIndex
	} else {
		ci.Index++
	}

	if ci.Count >= ci.MaxCount {
		return false
	}
	return true
}

func NewCyclicIterator(maxCount, maxIndex int) *CyclicIterator {
	return &CyclicIterator{
		MaxCount: maxCount,
		MaxIndex: maxIndex,
		Index:    -1,
		Count:    -1,
	}
}


//func TestCyclicIterator(t *testing.T) {
//	ci := NewCyclicIterator(21, 5)
//
//	for ci.Next() {
//		fmt.Printf("%v, %v\n", ci.Index, ci.Count)
//	}
//}