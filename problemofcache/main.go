package main

import (
	"fmt"
	"myproject/cache"
)

func main() {
	// cache := cache.NewCache()
	// cache.Set("tariq", 21)
	// cache.Set("hasak", 153)
	// cache.Print()

	// fmt.Println("\nGetting value:",cache.Get("hasak"))
	// fmt.Println("Getting value:",cache.Get("boy"))

	// fmt.Println("(tariq): ",cache.Delete("tariq"))
	// fmt.Println("(Boy): ",cache.Delete("boy"))

	// cache.Print()

	cache := cache.NewCache()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)

}
