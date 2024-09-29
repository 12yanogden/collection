package collection

import (
	"fmt"
	"reflect"
	// "strconv"
)

type Collection struct {
	data map[any]any
	order []any
}

func NewFromMap(initial map[any]any) *Collection {
	return &Collection{
		data: initial,
	}
}

func NewFromArray(initial []any) *Collection {
	collection := Collection{
		data: make(map[any]any),
		order: make([]any, 0),
	}

	for i := 0; i < len(initial); i++ {
		collection.data[i] = initial[i]
		collection.order = append(collection.order, initial[i])
	}

	return &collection
}

func (c *Collection) Get(key any) any {
	return c.data[key]
}

func (c *Collection) Set(key any, value any) {
	c.data[key] = value
}

func (c Collection) Map(fn func(any) any) Collection {
	for key, value := range c.data {
		c.data[key] = fn(value)
	}

	return c
}
 
func (c Collection) Concat(other Collection) Collection {
	c.
}

func (c Collection) String() string {
	return 
}
type stringable interface {
	String() string
}
