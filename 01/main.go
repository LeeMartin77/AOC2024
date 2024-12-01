package main

import (
	"fmt"
	"os"

	"github.com/LeeMartin77/AOC2024/01/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	res := solution.ComputeSolutionOne(data)

	fmt.Printf("Solution One: %d\n", res)
}
