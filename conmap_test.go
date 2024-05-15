package conmap_test

import (
	"fmt"
	"testing"

	"github.com/Winens/conmap"
)

func TestMap(t *testing.T) {
	m := conmap.New[int, string]()

	m.Store(1, "one")
	m.Store(2, "two")
	m.Store(3, "three")

	t.Run("Load", func(t *testing.T) {
		value, ok := m.Load(1)
		if !ok {
			t.Fatalf("expected key 1 to exist")
		}
		if value != "one" {
			t.Fatalf("expected value to be 'one', got %q", value)
		}
	})

	t.Run("Store", func(t *testing.T) {
		m.Store(1, "uno")
		value, ok := m.Load(1)
		if !ok {
			t.Fatalf("expected key 1 to exist")
		}
		if value != "uno" {
			t.Fatalf("expected value to be 'uno', got %q", value)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		m.Delete(1)
		_, ok := m.Load(1)
		if ok {
			t.Fatalf("expected key 1 to be deleted")
		}
	})

	t.Run("Range", func(t *testing.T) {
		m.Range(func(key int, value string) bool {
			fmt.Printf("%d: %s\n", key, value)
			return true
		})
	})

	t.Run("Len", func(t *testing.T) {
		if m.Len() != 2 {
			t.Fatalf("expected length to be 2, got %d", m.Len())
		}
	})
}

func BenchmarkMap(b *testing.B) {
	m := conmap.New[int, string]()

	b.Run("Store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(i, "value")
		}
	})

	b.Run("Load", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Delete(i)
		}
	})

	b.Run("Range", func(b *testing.B) {
		m.Range(func(key int, value string) bool {
			return true
		})
	})

	b.Run("Len", func(b *testing.B) {
		m.Len()
	})
}
