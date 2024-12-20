package main

import (
	"fmt"
	"os"

	"github.com/LeeMartin77/AOC2024/20/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	res := solution.ComputeSolutionOne(data, 100)

	fmt.Printf("Solution One: %d\n", res)

	res2 := solution.ComputeSolutionTwo(data, 100)

	fmt.Printf("Solution Two: %d\n", res2)
}
