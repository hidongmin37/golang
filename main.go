package main

import (
	"fmt"
	"maps"
)

func main() {

	// Maps
	// Key value store { a: 1, b: 2 }
	// [1,2,3,4,5,6] O(n)
	m := make(map[string]int)
	
	a, exists := m["a"]
	m["a"] = 1
	m["b"] = 2

	//if _, ok := m["a"]; ok {
	//
	//}
	equal := maps.Equal(m, map[string]int{"a": 1, "b": 2})
	fmt.Println(equal)
	delete(m, "a")
	// clear(m) whole map

	//

	fmt.Println(a, exists)

}
