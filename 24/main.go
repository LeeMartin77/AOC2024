package main

import (
	"fmt"
	"os"

	"github.com/LeeMartin77/AOC2024/24/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	res := solution.ComputeSolutionOne(data)

	fmt.Printf("Solution One: %d\n", res)

	res2 := solution.ComputeSolutionTwo(data, 4)

	fmt.Printf("Solution Two: %s\n", res2)
}
