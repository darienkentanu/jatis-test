package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	m := make(map[string]int)
	steps := 7
	m["A"] = 0
	m["B"] = 0
	m["C"] = 0
	m["D"] = 0
	m["E"] = 0

	insertRandomObj(m)

	for i := 1; i <= steps; i++ {
		for i := range m {
			if m[i] == 1 {
				// clear map
				m[i] = 0
				// insert random object adjacent with current box
				insertRandomObj2(m, i)
			}
		}
	}
	fmt.Println(m)
	objectInMap := findObj(m)
	fmt.Println(objectInMap)
}

func insertRandomObj2(m map[string]int, huruf string) {
	rand.Seed(time.Now().UnixNano())
	angka := 0
	for {
		angka = rand.Intn(2)
		if angka == 0 {
			continue
		}
		if angka != 0 {
			break
		}
	}
	var slc = make([]string, 0)
	for i := range m {
		slc = append(slc, i)
	}
	slcSorted := sortString(slc)
	for i, v := range slcSorted {
		temp := 0
		if huruf == v {
			if angka == 1 {
				temp = temp + i + 1
				if temp > len(slc)-1 {
					temp = 0
				}
				m[string(slc[temp])] = 1
				break
			} else {
				temp--
				if temp < 0 {
					temp = len(slc) - 1
				}
				m[string(slc[temp])] = 1
				break
			}
		}
	}
}

func findObj(m map[string]int) string {
	keys := []string{}
	for k := range m {
		if m[k] != 0 {
			keys = append(keys, k)
		}
	}
	return keys[0]
}

func insertRandomObj(m map[string]int) {
	rand.Seed(time.Now().UnixNano())
	randINT := rand.Intn(len(m))
	switch randINT {
	case 0:
		m["A"] = 1
	case 1:
		m["B"] = 1
	case 2:
		m["C"] = 1
	case 3:
		m["D"] = 1
	default:
		m["E"] = 1
	}
}

func sortString(stringsSlice []string) []string {
	strings1 := strings.Join(stringsSlice, "")
	var stringsSorted = make([]int, 0)
	for _, v := range strings1 {
		stringsSorted = append(stringsSorted, int(v))
	}
	sort.Ints(stringsSorted)
	var stringsSortedString = make([]byte, len(stringsSorted))
	for i := 0; i < len(stringsSorted); i++ {
		stringsSortedString[i] = byte(stringsSorted[i])
	}
	return strings.Split(string(stringsSortedString), "")
}
