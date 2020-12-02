package util

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFileOfInts(in string) ([]int, error) {
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	res := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}
