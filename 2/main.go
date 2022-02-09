package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var needs = 400

func main() {
	// 100 gr foods
	var foods = make(map[string]int)

	priceRice := PriceRice()
	priceCorn := PriceCorn()
	pricePotato := PricePotato()

	foods["rice"] = priceRice
	foods["corn"] = priceCorn
	foods["potato"] = pricePotato

	prices := make([]int, 0)
	prices = append(prices, priceRice)
	prices = append(prices, priceCorn)
	prices = append(prices, pricePotato)

	sort.Ints(prices)
	var food string
	for i := range foods {
		if foods[i] == prices[0] {
			food = i
		}
	}
	fmt.Println("you should buy", food, "because the price is the cheapest=", prices[0])
}

func PriceRice() int {
	return randomPriceRice() * needs / 28
}

func PriceCorn() int {
	return randomPriceCorn() * needs / 21
}

func PricePotato() int {
	return randomPricePotato() * needs / 17
}

func randomPriceRice() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000)
	return 28000 + random
}

func randomPriceCorn() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000)
	return 21000 + random
}

func randomPricePotato() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000)
	return 17000 + random
}
