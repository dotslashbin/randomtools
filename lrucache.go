package randomtools

import (
	"fmt"
)

type cache_container struct {
	values []string
}

// GenerateLRUCache generates the output of the LRU cache
func GenerateLRUCache(input string) string {

	cache := cache_container{}
	cachedVals := &cache.values

	for _, char := range input {
		valueToCheck := string(char)
		inArray, _ := in_array(valueToCheck, *cachedVals)

		if inArray == false {
			*cachedVals = append(*cachedVals, valueToCheck)
		} else {

			updateCollecton(&cache, valueToCheck)
		}
	}

	fmt.Println("Final: ")
	fmt.Println(cache)

	return "you are done"
}

// in_array Checks if a paritcular element is found in a collection
func in_array(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}

// updateCollection updates the cache container collection trough reference. This is where the main logic is
func updateCollecton(cache *cache_container, letter string) {

	for index, char := range cache.values {
		if string(char) == letter {
			// Removing the cached letter
			cache.values = append(cache.values[:index], cache.values[index+1:]...)

			// adding to the end
			cache.values = append(cache.values, letter)
		}
	}
}
