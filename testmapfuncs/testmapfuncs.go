package testmapfuncs

import (
	"fmt"
	"maps"
)

func TestMapFuncs() {
	someMap := make(map[string]int)
	someMap["hey"] = 1
	someMap["there"] = 2
	// iterate through key and value
	//OP: hey1, there2
	for key, val := range maps.All(someMap) {
		fmt.Println(key, val)
	}
	// iterate through key alone
	// OP: hey there
	for key := range maps.Keys(someMap) {
		fmt.Println(key)
	}
	// iterate through values alone
	// OP: 1 2
	for val := range maps.Values(someMap) {
		fmt.Println(val)
	}
}
