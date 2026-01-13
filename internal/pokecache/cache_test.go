package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(10 * time.Minute)
	cache.Add("test", []byte("test"))
	val, ok := cache.Get("test")
	if !ok {
		t.Errorf("expected value to be found")
	}
	if string(val) != "test" {
		t.Errorf("expected value to be 'test', got %s", string(val))
	}
}
