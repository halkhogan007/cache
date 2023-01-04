package cache

import (
	"errors"
)

type cache struct {
	memory map[string]interface{}
}

func New() *cache {
	return &cache{
		memory: make(map[string]interface{}),
	}
}
func (c cache) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("Invalid key")
	}
	c.memory[key] = value
	return nil
}
func (c cache) Get(key string) (interface{}, error) {
	if c.memory[key] == 0 {
		return errors.New("unknown key for get."), nil
	}
	return c.memory[key], nil
}
func (c cache) Delete(key string) error {
	if c.memory[key] == 0 {
		return errors.New("unknown key for delete.")
	}
	delete(c.memory, key)
	return nil
}
