package randomtools

import "fmt"

type cache []string

func GenerateLRUCache(input string) string {

	theCache := cache{}

	for _, letter := range input {
		theCache = append(theCache, string(letter))
	}

	fmt.Println(theCache)

	return "you are done"
}
