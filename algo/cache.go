package algo

import (
	"fmt"
	"math/rand"
)

type FuncType func(int) int

func Cache() {
	cachedFibonacci := cacheRun(fibonacci)
	iterations := rand.Intn(20) + 1

	for i := 1; i <= iterations; i++ {
		cachedFibonacci(rand.Intn(i) + 1)
	}
}

func cacheRun(fn FuncType) FuncType {
	cache := make(map[int]int)

	return func(n int) int {
		if val, isExist := cache[n]; isExist {
			fmt.Println("Get from cache: ", val)
			return val
		}
		cache[n] = fn(n)
		fmt.Println("Calculated: ", cache[n])
		return cache[n]
	}
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
