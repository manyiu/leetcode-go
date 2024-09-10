package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Fatal("get 1 failed")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Fatal("get 2 failed")
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Fatal("get 1 failed")
	}
	if cache.Get(3) != 3 {
		t.Fatal("get 3 failed")
	}
	if cache.Get(4) != 4 {
		t.Fatal("get 4 failed")
	}

	cache2 := Constructor(2)
	cache2.Put(1, 0)
	cache2.Put(2, 2)
	if cache2.Get(1) != 0 {
		t.Fatal("get 1 failed")
	}
	cache2.Put(3, 3)
	if cache2.Get(2) != -1 {
		t.Fatal("get 2 failed")
	}
	cache2.Put(4, 4)
	if cache2.Get(1) != -1 {
		t.Fatal("get 1 failed")
	}
	if cache2.Get(3) != 3 {
		t.Fatal("get 3 failed")
	}
	if cache2.Get(4) != 4 {
		t.Fatal("get 4 failed")
	}
}
