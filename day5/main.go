package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
	"sort"
)

// FBFBBFFRLR
// First 7 define the row of seat
// Last 3 define the column of seat

func midpoint(a, b int) int {
	return (a + b) / 2
}

func NewPlane(minRow, maxRow, minCol, maxCol int) *Plane {
	return &Plane{
		minRow: minRow,
		maxRow: maxRow,
		minCol: minCol,
		maxCol: maxCol,
	}
}

func (s *Plane) FindSeat(input string) (int, int) {
	rowRange := &Range{
		Min: s.minRow,
		Max: s.maxRow,
	}
	colRange := &Range{
		Min: s.minCol,
		Max: s.maxCol,
	}
	for i, val := range input {
		value := StringToRangeDirection(string(val))
		if i < 7 {
			rowRange.Apply(value)
			continue
		}
		colRange.Apply(value)
	}
	return rowRange.Min, colRange.Min
}

func (s *Plane) FindSeatID(input string) int {
	row, col := s.FindSeat(input)
	return (row * 8) + col
}

type Plane struct {
	minCol int
	maxCol int
	minRow int
	maxRow int
}

type RangeDirection int
const (
	RangeDirectionLower RangeDirection = iota
	RangeDirectionHigher
)

func StringToRangeDirection(value string) RangeDirection {
	switch value {
	case "F", "L":
		return RangeDirectionLower
	case "B", "R":
		return RangeDirectionHigher
	default:
		panic("unknown range direction value")
	}
}

type Range struct {
	Min int
	Max int
}

func (rr *Range) Apply(value RangeDirection) {
	switch value {
	case RangeDirectionLower:
		rr.Max = midpoint(rr.Min, rr.Max) - 1
	case RangeDirectionHigher:
		rr.Min = midpoint(rr.Min, rr.Max) + 1
	default:
		panic("unknown range direction")
	}
}

const (
	minCol = 0
	maxCol = 7
	minRow = 0
	maxRow = 127
)

func main() {
	plane := NewPlane(minRow, maxRow, minCol, maxCol)
	ids := make([]int, 0)
	if err := util.ParseFileLineByLine("day5/input.txt", func(line string) error {
		id := plane.FindSeatID(line)
		ids = append(ids, id)

		return nil
	}); err != nil {
		panic(err)
	}

	sort.Ints(ids)

	var prev, id int
	for i, x := range ids {
		if i != 0 && x == (prev + 2) {
			id = prev + 1
			continue
		}
		prev = x
	}

	fmt.Println(ids[len(ids) - 1])
	fmt.Println(id)
}
