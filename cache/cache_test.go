package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCache(t *testing.T) {
	stringCache := NewSimpleCache[string, string]()
	stringCache.Set("key", "value")

	strVal, err := stringCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", strVal)

	intCache := NewSimpleCache[int, int]()
	intCache.Set(1, 1)
	intVal, err := intCache.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, intVal)

	intVal, err = intCache.Get(2)
	assert.Error(t, err)
	assert.Equal(t, "key not found", err.Error())
	assert.Equal(t, 0, intVal)

	floatCache := NewSimpleCache[string, float64]()
	floatCache.Set("key", 1.0)
	floatVal, err := floatCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, 1.0, floatVal)
	floatCache.Delete("key")
	floatVal, err = floatCache.Get("key")
	assert.Error(t, err)
	assert.Equal(t, "key not found", err.Error())
	assert.Equal(t, 0.0, floatVal)

	type myStruct struct {
		ID   int
		Name string
	}

	structCache := NewSimpleCache[string, myStruct]()
	structCache.Set("key", myStruct{ID: 1, Name: "value"})
	structVal, err := structCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, myStruct{ID: 1, Name: "value"}, structVal)
	structCache.Delete("key")
	structVal, err = structCache.Get("key")
	assert.Error(t, err)
	assert.Equal(t, "key not found", err.Error())
	assert.Equal(t, myStruct{}, structVal)
}
