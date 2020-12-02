package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
	"regexp"
	"strconv"
)

type Policy struct {
	Min  int
	Max  int
	Char rune
}

var (
	lineRegex = regexp.MustCompile(`^([0-9]+)-([0-9]+) ([a-z]+): (.*)$`)
)

func (p Policy) CheckA(password string) bool {
	num := 0
	for _, x := range password {
		if x == p.Char {
			num++
		}
	}
	return num >= p.Min && num <= p.Max
}

func (p Policy) CheckB(password string) bool {
	pw := []rune(password)
	foundMin := len(pw) >= p.Min && pw[p.Min-1] == p.Char
	foundMax := len(pw) >= p.Max && pw[p.Max-1] == p.Char
	return util.Xor(foundMin, foundMax)
}

type Item struct {
	Password string
	Policy   Policy
}

func main() {
	items := make([]Item, 0)
	if err := util.ParseFileLineByLine("day2/input.txt", func(line string) error {
		lineMatch := lineRegex.FindStringSubmatch(line)

		min, err := strconv.Atoi(lineMatch[1])
		if err != nil {
			return err
		}
		max, err := strconv.Atoi(lineMatch[2])
		if err != nil {
			return err
		}
		p := Policy{
			Min:  min,
			Max:  max,
			Char: []rune(lineMatch[3])[0],
		}
		password := lineMatch[4]

		items = append(items, Item{
			Password: password,
			Policy:   p,
		})
		return nil
	}); err != nil {
		panic(err)
	}

	validA := 0
	validB := 0

	for _, i := range items {
		if i.Policy.CheckA(i.Password) {
			validA++
		}
		if i.Policy.CheckB(i.Password) {
			validB++
		}
	}

	fmt.Println(validA)
	fmt.Println(validB)
}
