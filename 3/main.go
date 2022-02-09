package main

import (
	"fmt"
)

func main() {
	fibs := FibonacciFromXY(2, 3)
	fmt.Println(findOddNumberAndEven(fibs))
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func FibonacciFromXY(x, y int) []int {
	fibs := make([]int, 0)
	for i := x + 1; i <= y+2; i++ {
		fibs = append(fibs, FibonacciRecursion(i))
	}
	return fibs
}

func findOddNumberAndEven(slc []int) string {
	odd := 0
	even := 0
	for _, v := range slc {
		if v%2 == 0 {
			odd = odd + v
		} else {
			even = even + v
		}
	}
	return fmt.Sprintf("sum odd number = %d, sum even number = %d", odd, even)
}
