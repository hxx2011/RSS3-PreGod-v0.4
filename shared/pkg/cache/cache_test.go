package cache_test

import (
	"strconv"
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/stretchr/testify/assert"
)

type Object struct {
	Name  string
	Score int64
}

func TestGetSet(t *testing.T) {
	t.Parallel()

	err := cache.Setup()
	assert.Nil(t, err)

	key := "TestGetSet"
	valueSet := &Object{Name: key}

	err = cache.Set(key, valueSet, 0)
	assert.Nil(t, err)

	valueGet := &Object{}
	err = cache.Get(key, valueGet)
	assert.Nil(t, err)
	assert.Equal(t, valueSet.Name, valueGet.Name)
}

func TestExists(t *testing.T) {
	t.Parallel()

	err := cache.Setup()
	assert.Nil(t, err)

	key := "TestGetSet"

	e, err := cache.Exists(key)
	assert.Nil(t, err)
	assert.True(t, e)

	ne, err := cache.Exists("key_not_exist")
	assert.Nil(t, err)
	assert.False(t, ne)
}

func TestZAdd(t *testing.T) {
	t.Parallel()

	err := cache.Setup()
	assert.Nil(t, err)

	key := "TestZAdd"

	n := 0
	for n < 3 {
		err = cache.ZAdd(key, &Object{Name: key + strconv.Itoa(n), Score: int64(n)}, float64(n))
		assert.Nil(t, err)

		n++
	}
}

func TestZRevRange(t *testing.T) {
	t.Parallel()

	err := cache.Setup()
	assert.Nil(t, err)

	key := "TestZAdd"

	len1 := 1
	len2 := 2

	result, err := cache.ZRevRange(key, "0", "1", 0, int64(len1))
	assert.Nil(t, err)
	assert.True(t, len(result) == len1)

	result, err = cache.ZRevRange(key, "0", "2", 0, int64(len2))
	assert.Nil(t, err)
	assert.True(t, len(result) == len2)

	for i, s := range result {
		resMap, ok := s.(map[string]interface{})
		assert.True(t, ok)
		assert.True(t, resMap["Name"] == "TestZAdd"+strconv.Itoa(len(result)-i))
	}
}

func TestZRevRangeWithScore(t *testing.T) {
	t.Parallel()

	err := cache.Setup()
	assert.Nil(t, err)

	key := "TestZAdd"

	len1 := 1
	len2 := 2

	result, err := cache.ZRevRangeWithScore(key, "0", "1", 0, int64(len1))
	assert.Nil(t, err)
	assert.True(t, len(result) == len1)

	result, err = cache.ZRevRangeWithScore(key, "1", "2", 0, int64(len2))
	assert.Nil(t, err)
	assert.True(t, len(result) == len2)

	for i, s := range result {
		resMap, ok := s.(map[string]interface{})
		assert.True(t, ok)
		assert.True(t, resMap["Name"] == "TestZAdd"+strconv.Itoa(len(result)-i))
	}
}
