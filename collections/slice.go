package collections

type Slice struct {
	keys   []int
	values []int
}

func (c *Slice) findIndex(key int) int {
	for i, k := range c.keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (c *Slice) setByIndex(index int, value int) {
	c.values[index] = value
}

func (c *Slice) deleteByIndex(index int) {
	if freeSpace := cap(c.keys) - len(c.keys); freeSpace < len(c.keys)*2 { // not too much free space
		copy(c.keys[index:], c.keys[index+1:])
		copy(c.values[index:], c.values[index+1:])
		c.keys = c.keys[:len(c.keys)-1]
		c.values = c.values[:len(c.values)-1]
		return
	}

	//recreate
	pk1 := c.keys[:index]
	pk2 := c.keys[index+1:]
	c.keys = make([]int, len(c.keys)-1, (len(c.keys)-1)*3/2)
	copy(c.keys, pk1)
	copy(c.keys[index:], pk2)

	pv1 := c.values[:index]
	pv2 := c.values[index+1:]
	c.values = make([]int, len(c.values)-1, (len(c.values)-1)*3/2)
	copy(c.values, pv1)
	copy(c.values[index:], pv2)
}

func (c *Slice) add(key int, value int) {
	c.keys = append(c.keys, key)
	c.values = append(c.values, value)
}

func (c *Slice) Clear() {
	c.keys = make([]int, 0)
	c.values = make([]int, 0)
}

func (c *Slice) Get(key int) (value int, ok bool) {
	index := c.findIndex(key)
	if index == -1 {
		return 0, false
	}
	return c.values[index], true
}

func (c *Slice) Set(key int, value int) {
	index := c.findIndex(key)
	if index == -1 {
		c.add(key, value)
		return
	}
	c.setByIndex(index, value)
}

func (c *Slice) Delete(key int) {
	index := c.findIndex(key)
	if index == -1 {
		return
	}
	c.deleteByIndex(index)
}

func (c *Slice) ForEach(f ForEachFunction) {
	for i, k := range c.keys {
		f(k, c.values[i])
	}
}

func NewSlice() *Slice {
	return &Slice{
		keys:   make([]int, 0),
		values: make([]int, 0),
	}
}
