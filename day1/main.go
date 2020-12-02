package main

import (
	"advent-of-code-2020/util"
	"fmt"
)

func main() {
	input, err := util.ParseFileOfInts("day1/input.txt")
	if err != nil {
		panic(err)
	}

	a, b := findSubjects(2020, input...)
	res := a * b
	fmt.Println(res)
}

func findSubjects(total int, vals ...int) (int, int) {
	for ia, a := range vals {
		for ib, b := range vals {
			if ia == ib {
				continue
			}
			if a+b == total {
				return a, b
			}
		}
	}
	return 0, 0
}
