package lru

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapacityAlwaysReplace(t *testing.T) {
	lru := NewLRUCache(1)

	criteria := []struct {
		items      map[string]any
		removedKey string
		addedKey   string
		addedValue any
	}{
		{map[string]any{"a": 1}, "", "a", 1},
		{map[string]any{"b": 3}, "a", "b", 3},
		{map[string]any{"c": true}, "b", "c", true},
		{map[string]any{"d": "a"}, "c", "d", "a"},
	}


	for _, c := range criteria {
		lru.Set(c.addedKey, c.addedValue)

		for i := range c.items {
			if _, ok := lru.items[i]; !ok {
				t.Errorf("Expected key %s , found nil", i)
				return
			}

			if lru.items[i].value != c.items[i] {
				t.Errorf("Expected value %v, found %v", c.items[i], lru.items[i].value)
			}
		}

		if _, ok := c.items[c.removedKey]; ok {
			t.Errorf("Expected key %s to be removed, found in items", c.removedKey)
		}

		if len(c.items) != len(lru.items) {
			t.Errorf("Expected items length %d, found %d", len(c.items), len(lru.items))
		}
	}
}

func TestGetCorrectValues(t *testing.T) {
    lru := NewLRUCache(4)
    lru.Set("a", 1)
    lru.Set("b", true)
    lru.Set("c", "string")
    lru.Set("d", 3.14)

    if value, _ := lru.Get("a"); value != 1 {
        t.Errorf("Expected value 1, found %v", value)
    }
    if value, _ := lru.Get("b"); value != true {
        t.Errorf("Expected value true, found %v", value)
    }
    if value, _ := lru.Get("c"); value != "string" {
        t.Errorf("Expected value string, found %v", value)
    }
    if value, _ := lru.Get("d"); value != 3.14 {
        t.Errorf("Expected value 3.14, found %v", value)
    }
}

func TestGetNotFoundError(t *testing.T) {
    lru := NewLRUCache(1)
    _, err := lru.Get("a")
    assert.True(t, errors.Is(err, ErrKeyNotFound), "Expected error not found, found %v", err)
}

func TestReplaceCorrectKey(t *testing.T) {
    lru := NewLRUCache(3)
    expected := map[string]any{"a": 1, "e": 4, "d": 3}
    lru.Set("a", 1)
    lru.Set("b", 3)
    lru.Set("c", 2)
    lru.Get("a")
    lru.Set("d", 3)
    lru.Set("e", 4)

    for i := range expected {
        if _, ok := lru.items[i]; !ok {
            t.Errorf("Expected key %s, found nil", i)
            return
        }

        if lru.items[i].value != expected[i] {
            t.Errorf("Expected value %v, found %v", expected[i], lru.items[i].value)
        }
    }

    if len(lru.items) != 3 {
        t.Errorf("Expected items length 3, found %d", len(lru.items))
    }

    if _, ok := lru.items["b"]; ok {
        t.Errorf("Expected key b to be removed, found in items")
    }

    if _, ok := lru.items["c"]; ok {
        t.Errorf("Expected key c to be removed, found in items")
    }
}

