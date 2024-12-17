package main

import (
	"fmt"
	"os"

	"github.com/LeeMartin77/AOC2024/17/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	res := solution.ComputeSolutionOne(data)

	fmt.Printf("Solution One: %s\n", res)

	// res2 := solution.ComputeSolutionTwo(data)

	// fmt.Printf("Solution Two: %d\n", res2)
}
