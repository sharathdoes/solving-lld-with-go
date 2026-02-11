package main

import (
	"fmt"
	"lru-cache/internal/dll"
)

func main() {

	lru := dll.NewLRU(2)

	output := make([]interface{}, 0)

	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1,1], [2,2], [1], [3,3], [2], [4,4], [1], [3], [4]]

	output = append(output, nil) // LRUCache init

	// put(1,1)
	lru.PushValue("1", "1")
	output = append(output, nil)

	// put(2,2)
	lru.PushValue("2", "2")
	output = append(output, nil)

	// get(1)
	if val, ok := lru.GetValue("1"); ok {
		output = append(output, val)
	} else {
		output = append(output, -1)
	}

	// put(3,3)
	lru.PushValue("3", "3")
	output = append(output, nil)

	// get(2)
	if val, ok := lru.GetValue("2"); ok {
		output = append(output, val)
	} else {
		output = append(output, -1)
	}

	// put(4,4)
	lru.PushValue("4", "4")
	output = append(output, nil)

	// get(1)
	if val, ok := lru.GetValue("1"); ok {
		output = append(output, val)
	} else {
		output = append(output, -1)
	}

	// get(3)
	if val, ok := lru.GetValue("3"); ok {
		output = append(output, val)
	} else {
		output = append(output, -1)
	}

	// get(4)
	if val, ok := lru.GetValue("4"); ok {
		output = append(output, val)
	} else {
		output = append(output, -1)
	}

	fmt.Println(output)

}
