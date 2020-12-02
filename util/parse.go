package util

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFileOfInts(in string) ([]int, error) {
	res := make([]int, 0)
	err := ParseFileLineByLine(in, func(line string) error {
		i, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		res = append(res, i)
		return nil
	})
	return res, err
}

func ParseFileLineByLine(in string, fn func(line string) error) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if err := fn(s.Text()); err != nil {
			return err
		}
	}
	return s.Err()
}

func Xor(a, b bool) bool {
	return (a || b) && (!a || !b)
}
