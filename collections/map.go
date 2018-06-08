package collections

type Map struct {
	theMap map[int]int
}

func (c *Map) Clear() {
	c.theMap = make(map[int]int)
}

func (c *Map) Get(key int) (value int, ok bool) {
	value, ok = c.theMap[key]
	return
}

func (c *Map) Set(key int, value int) {
	c.theMap[key] = value
}

func (c *Map) Delete(key int) {
	delete(c.theMap, key)
}

func (c *Map) ForEach(f ForEachFunction) {
	for k, v := range c.theMap {
		f(k, v)
	}
}

func NewMap() *Map {
	return &Map{
		theMap: make(map[int]int),
	}
}
