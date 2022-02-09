package main

import "fmt"

func main() {
	fmt.Println(threediceProbability(18))
}

func threediceProbability(x int) float64 {
	if x < 3 {
		fmt.Println("x must be greater than 2")
		return 0
	}
	if x > 18 {
		fmt.Println("x must be lower than 19")
		return 0
	}
	dice := []int{1, 2, 3, 4, 5, 6}
	probability := float64(1) / float64((len(dice))*3)
	return float64(probability)
}
