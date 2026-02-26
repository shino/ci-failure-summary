package main

import "fmt"

func main() {
	fmt.Println("Starting CI failure summary demo...")
	fmt.Println("The quick brown fox jumps over the lazy dog.")
	fmt.Println("Initializing components...")
	fmt.Println("Processing data...")
	fmt.Println("About to encounter a fatal error!")

	recurse(0)
}

func recurse(depth int) {
	if depth >= 20 {
		panic("fatal: something went terribly wrong at depth 20")
	}
	recurse(depth + 1)
}
