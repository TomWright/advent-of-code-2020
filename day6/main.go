package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
)

type Group []Answers

func (g Group) CountA() int {
	uniqueAnswers := make([]string, 0)
	for _, answers := range g {
		for _, answer := range answers {
			if !util.StringSliceContains(answer, uniqueAnswers) {
				uniqueAnswers = append(uniqueAnswers, answer)
			}
		}
	}
	return len(uniqueAnswers)
}

func (g Group) CountB() int {
	tally := make(map[string]int)
	for _, answers := range g {
		for _, answer := range answers {
			tally[answer]++
		}
	}
	required := len(g)
	total := 0
	for _, v := range tally {
		if v == required {
			total++
		}
	}
	return total
}

type Answers []string

func main() {
	groups := make([]Group, 0)
	group := Group{}
	if err := util.ParseFileLineByLine("day6/input.txt", func(line string) error {
		if line == "" {
			groups = append(groups, group)
			group = Group{}
			return nil
		}

		a := Answers{}
		for _, answer := range line {
			a = append(a, string(answer))
		}
		group = append(group, a)

		return nil
	}); err != nil {
		panic(err)
	}
	groups = append(groups, group)

	totalA := 0
	totalB := 0
	for _, g := range groups {
		totalA += g.CountA()
		totalB += g.CountB()
	}

	fmt.Println(totalA)
	fmt.Println(totalB)
}
