package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
)

type Area struct {
	Width  int
	Height int
	Data   [][]*Position
}

func (a Area) GetPosition(x int, y int) *Position {
	if y >= a.Height {
		return nil
	}
	if x >= a.Width {
		x = x % a.Width
	}
	return a.Data[y][x]
}

type Position struct {
	Content string
}

func (p Position) IsTree() bool {
	return p.Content == "#"
}

func (p Position) IsOpen() bool {
	return p.Content == "."
}

func slide(a Area, x, y int, stepX, stepY int) int {
	treeCount := 0
	for {
		pos := a.GetPosition(x, y)
		if pos == nil {
			break
		}
		if pos.IsTree() {
			treeCount++
		}

		x += stepX
		y += stepY
	}

	return treeCount
}

func main() {
	area := Area{
		Data: [][]*Position{},
	}

	if err := util.ParseFileLineByLine("day3/input.txt", func(line string) error {
		lineData := make([]*Position, 0)
		for _, char := range line {
			lineData = append(lineData, &Position{Content: string(char)})
		}

		area.Data = append(area.Data, lineData)

		return nil
	}); err != nil {
		panic(err)
	}

	area.Height = len(area.Data)
	area.Width = len(area.Data[0])

	total := -1

	for _, xy := range [][]int{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	} {
		res := slide(area, 0, 0, xy[0], xy[1])
		if total == -1 {
			total = res
			continue
		}
		total = total * res
	}

	fmt.Println(total)

}
