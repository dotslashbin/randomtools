package randomtools

import (
	"fmt"
)

// Executes the proess of generating the LRU cache
func GenerateLRUCache(input string) string {

	theCache := []string{}

	for _, value := range input {
		char := string(value)
		inArray, _ := in_array(char, theCache)

		if inArray == false {
			theCache = append(theCache, string(char))
		} else {
			theCache = updateCollecton(theCache, value)
		}
	}

	fmt.Println("final", theCache)

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

func updateCollecton(theCache []string, letter rune) []string {

	for index, char := range theCache {
		fmt.Println("removing " + string(letter) + " from ...")1
		fmt.Println(theCache)
		if string(letter) == string(char) {
			theCache = append(theCache[:index], theCache[index+1:]...)
		}
	}

	return theCache

}