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
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
