package cache

import "fmt"

type cache struct {
	elements map[string]interface{}
}

func NewCache() *cache {
	elements := make(map[string]interface{})
	return &cache{
		elements: elements,
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.elements[key] = value
}

func (c *cache) Get(key string) interface{} {
	return c.elements[key]
}

func (c *cache) Delete(key string) string {
	if c.elements[key] != nil {
		delete(c.elements, key)
		return "Elementy jnjvec"
	}
	return "Aydpisi key chka"
}

func (c *cache) Print() {
	for key, value := range c.elements {
		fmt.Printf("{%v : %v}  ",key,value)
	}
}
