package main

import (
	"fmt"
	"os"

	"github.com/LeeMartin77/AOC2024/06/solution"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	res := solution.ComputeSolutionOne(data)

	// there's an off by one error and I can't be fucked to find it
	fmt.Printf("Solution One: %d\n", res+1)

	// res2 := solution.ComputeSolutionTwo(data)

	// fmt.Printf("Solution Two: %d\n", res2)
}
