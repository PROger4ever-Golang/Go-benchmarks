package collections

type ForEachFunction func(key int, value int)

type Collection interface {
	Clear()
	Get(key int) (value int, ok bool)
	Set(key int, value int)
	Delete(key int)
	ForEach(f ForEachFunction)
}
