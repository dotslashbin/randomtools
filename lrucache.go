package randomtools

import (
	"fmt"
)

type cache_container struct {
	values []string
}

// Executes the proess of generating the LRU cache
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

func in_array(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}

func updateCollecton(cache *cache_container, letter string) {

	for index, char := range cache.values {
		if string(char) == letter {
			// Removing the cached letter 
			fmt.Println("Before : ", cache)
			cache.values = append(cache.values[:index], cache.values[index+1:]...)

			// adding to the end
			cache.values = append(cache.values, letter)
			fmt.Println("after: ", cache)
		}
	}
}
