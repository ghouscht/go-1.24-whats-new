package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Blob []byte

func (b Blob) String() string {
	return fmt.Sprintf("Blob(%d KB)", len(b)/1024)
}

// newBlob returns a new Blob of the given size in KB.
func newBlob(size int) *Blob {
	b := make([]byte, size*1024)
	for i := range size {
		b[i] = byte(i) % 255
	}
	return (*Blob)(&b)
}

type Cache struct {
	cache map[string]weak.Pointer[Blob]
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]weak.Pointer[Blob]),
	}
}

// Set stores a value in the cache.
func (c *Cache) Set(key string, value *Blob) {
	c.cache[key] = weak.Make(value) // Store weak reference

	// Pre Go 1.24, we would have to use runtime.SetFinalizer which is notoriously hard to use,
	// see also: https://pkg.go.dev/runtime@master#SetFinalizer
	// ⬇️ Is where the magic 🪄 happens
	runtime.AddCleanup(value, func(key string) {
		// Please keep in mind, that the cleanup function is not guaranteed to run immediately
		// after an object is garbage collected.
		delete(c.cache, key)
	}, key)
}

// Get retrieves a value, returning nil if it's been garbage collected.
func (c *Cache) Get(key string) *Blob {
	if ptr, ok := c.cache[key]; ok {
		return ptr.Value()
	}
	return nil
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	return len(c.cache)
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.cache))
	for key := range c.cache {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	cache := NewCache()

	myBlob := newBlob(1000)

	// Store a value
	cache.Set("myBlob", myBlob)

	// Retrieve before GC
	fmt.Println("Before GC:", cache.Get("myBlob")) // Should print: Blob(1000 KB)

	myBlob = nil // Remove the strong reference
	// myBlob no longer has a strong reference, it should be garbage collected
	runtime.GC()

	// Try to retrieve after GC
	fmt.Println("After GC:", cache.Get("myBlob")) // Likely nil

	// ⬇️ do you rember this? With finalizers we can make this work as one would expect
	fmt.Println("Cache size:", cache.Len())
	fmt.Println("Cache keys:", cache.Keys())
}
