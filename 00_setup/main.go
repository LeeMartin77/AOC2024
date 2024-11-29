package main

import (
	"os"

	"github.com/LeeMartin77/AOC2024/00_setup/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	solution.Some_dummy_code(data)
}
