package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	strings := randomString(10)
	fmt.Println("unsorted:", strings)
	var stringsSorted = make([]int, 0)
	for _, v := range strings {
		stringsSorted = append(stringsSorted, int(v))
	}
	sort.Ints(stringsSorted)
	var stringsSortedString = make([]byte, len(stringsSorted))
	for i := 0; i < len(stringsSorted); i++ {
		stringsSortedString[i] = byte(stringsSorted[i])
	}
	fmt.Println("sorted string:", string(stringsSortedString))
	count := countOccurence(stringsSortedString)
	for i, v := range count {
		if v == 0 {
			continue
		}
		fmt.Printf("occurence of letter %s = %d\n", string(i), v)
	}
}

func countOccurence(s []byte) map[string]int {
	var count = make(map[string]int)
	for _, v := range letters {
		count[string(v)] = 0
	}
	for _, v1 := range s {
		temp := v1
		for i := range count {
			if string(i) == string(temp) {
				count[string(temp)]++
				break
			}
		}
	}
	return count
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
