package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Map passed as parameter
func updateMap(m map[string]int) {
	m["updated"] = 100
}

// Using errgroup for concurrent map operations
func positiveCase1() {
	var g errgroup.Group

	store := make(map[string]string)
	store["item1"] = "value1"
	store["item2"] = "value2"

	for key, val := range store {
		k, v := key, val
		g.Go(func() error {
			// ruleid: concurrent-writes-iteration
			store[k] = v + "_modified"
			return nil
		})
	}

	// Wait for all goroutines to complete
	g.Wait()
}

type DataStore struct {
	Items map[string]Item
}

type Item struct {
	Value string
}

// using sync.Map for concurrent safety
func negativeCase1() {
	var safeMap sync.Map

	go func() {
		// ok: concurrent-writes-iteration
		safeMap.Store("key", 1)
	}()

	safeMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

// Iterating over a copy for reads
func negativeCase2() {
	var g errgroup.Group
	store := make(map[string]string)
	store["item1"] = "value1"
	store["item2"] = "value2"

	// Safe copy of keys to avoid iteration during modification
	keys := make([]string, 0, len(store))
	for k := range store {
		keys = append(keys, k)
	}

	for _, key := range keys {
		k := key
		g.Go(func() error {
			// ok: concurrent-writes-iteration
			store[k] = store[k] + "_modified"
			return nil
		})
	}

	g.Wait()
}

// Map passed to goroutine directly (pattern excluded by rule)
func negativeCase3() {
	myMap := make(map[string]int)

	// ok: concurrent-writes-iteration
	go func(m map[string]int) {
		m["key"] = 1
	}(myMap)

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

// Sequential operations (no concurrency)
func negativeCase4() {
	myMap := make(map[string]int)
	myMap["key1"] = 1

	// sequential, not concurrent
	for k, v := range myMap {
		// ok: concurrent-writes-iteration
		fmt.Println(k, v)
	}

	myMap["key2"] = 2
}

func main() {
	positiveCase1()

	negativeCase1()
	negativeCase2()
	negativeCase3()
	negativeCase4()
}
