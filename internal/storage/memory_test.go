package storage

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("get non-existing key", func(t *testing.T) {
		s := Memory{
			mu:   sync.RWMutex{},
			data: map[string][]byte{},
		}

		val, ok := s.Get("foo")
		assert.Nil(t, val)
		assert.False(t, ok)
	})

	t.Run("get existing key", func(t *testing.T) {
		s := Memory{
			mu:   sync.RWMutex{},
			data: map[string][]byte{"foo": []byte("bar")},
		}

		val, ok := s.Get("foo")
		assert.Equal(t, []byte("bar"), val)
		assert.True(t, ok)
	})

	t.Run("set key", func(t *testing.T) {
		s := NewMemory()
		s.Set("foo", []byte("bar"))

		val, ok := s.Get("foo")
		assert.Equal(t, []byte("bar"), val)
		assert.True(t, ok)
	})

	t.Run("delete key", func(t *testing.T) {
		s := NewMemory()
		s.Set("foo", []byte("bar"))

		s.Delete("foo")

		val, ok := s.Get("foo")
		assert.Nil(t, val)
		assert.False(t, ok)
	})
}
