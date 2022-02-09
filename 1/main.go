package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var litters = 100
	fmt.Println("the capacity needs to be contained in bottles =", litters)
	bottles := listOfPossibleBottle(0, 30)
	get3Bottle := get3bottles(bottles)
	// fmt.Println("list of bottle =", get3Bottle)
	needOfBottle := getNeedofBottle(get3Bottle, litters)
	// fmt.Println("need of bottle =", needOfBottle)
	var m = make(map[int]int)
	for _, v1 := range get3Bottle {
		m[v1] = 0
		for _, v2 := range needOfBottle {
			if v1 == v2 {
				m[v1]++
			}
		}
	}
	for i, v := range m {
		fmt.Printf("need of bottle number %d = %d \n", i, v)
	}
}

func getNeedofBottle(threebottles []int, litters int) (output []int) {
	// sort from max
	threebottlesSorted := sortFromMax(threebottles)
	var n int = len(threebottlesSorted)
	for i := 0; i < n; i++ {
		if litters-threebottlesSorted[i] >= 0 {
			litters = litters - threebottlesSorted[i]
			output = append(output, threebottlesSorted[i])
			i--
		}
	}
	return output
}

func sortFromMax(elements []int) []int {
	n := len(elements)

	for k := 0; k < n; k++ {
		max := k
		for j := k + 1; j < n; j++ {
			if elements[j] > elements[max] {
				max = j
			}
		}
		elements[k], elements[max] = elements[max], elements[k]
	}
	return elements
}

func get3bottles(slc []int) (slcOut []int) {
	for {
		one := getRandIdx(slc)
		two := getRandIdx(slc)
		three := getRandIdx(slc)
		if one == 0 || two == 0 || three == 0 {
			continue
		}
		if one != two && two != three && three != one {
			return append(slcOut, one, two, three)
		}
	}
}

func getRandIdx(slc []int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(slc))
}

func listOfPossibleBottle(x, y int) []int {
	slc := make([]int, 0)
	for i := x; i <= y; i++ {
		if primeNumber(i) {
			slc = append(slc, i)
		}
	}
	return slc
}

func primeNumber(number int) bool {
	// cari faktor
	s := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			s = append(s, i)
		}
	}
	return len(s) == 2
}
