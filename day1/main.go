package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
)

func main() {
	input, err := util.ParseFileOfInts("day1/input.txt")
	if err != nil {
		panic(err)
	}

	{
		a, b := findSubjects1(2020, input...)
		res := a * b
		fmt.Println(res)
	}
	{
		a, b, c := findSubjects2(2020, input...)
		res := a * b * c
		fmt.Println(res)
	}
}

func findSubjects1(total int, vals ...int) (int, int) {
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

func findSubjects2(total int, vals ...int) (int, int, int) {
	for ia, a := range vals {
		for ib, b := range vals {
			for ic, c := range vals {
				if ia == ib || ia == ic || ib == ic {
					continue
				}
				if a+b+c == total {
					return a, b, c
				}
			}
		}
	}
	return 0, 0, 0
}
