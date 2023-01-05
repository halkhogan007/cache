package cache

import (
	"errors"
)

type Cache struct {
	memory map[string]interface{}
}

func New() *Cache {
	return &Cache{
		memory: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("Invalid key")
	}
	c.memory[key] = value
	return nil
}
func (c *Cache) Get(key string) (interface{}, error) {
	if c.memory[key] == nil {
		return nil, errors.New("unknown key for get.")
	}
	return c.memory[key], nil
}
func (c *Cache) Delete(key string) error {
	if c.memory[key] == nil {
		return errors.New("unknown key for delete.")
	}
	delete(c.memory, key)
	return nil
}
