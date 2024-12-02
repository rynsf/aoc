package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeIndex(arr []int, i int) []int {
	r := make([]int, 0)
	r = append(r, arr[:i]...)
	r = append(r, arr[i+1:]...)
	return r
}

func isLevelSafe(level []int) (bool, int) {
	if len(level) < 2 {
		return true, -1
	}
	ascending := isAscending(level[0], level[1])
	for i := 0; i < len(level)-1; i++ {
		if !(isGradual(level[i], level[i+1]) && isAscending(level[i], level[i+1]) == ascending) {
			return false, i
		}
	}
	return true, -1
}

func isGradual(a int, b int) bool {
	v := a - b
	if v < 0 {
		v = -v
	}
	return v >= 1 && v <= 3
}

func isAscending(a int, b int) bool {
	return a-b < 0
}

func main() {
	var reports [][]int
	var safe int
	f, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		parts := strings.Fields(line)
		var level []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			level = append(level, num)
		}
		reports = append(reports, level)
	}

	for _, level := range reports {
		s, _ := isLevelSafe(level)
		if s {
			safe += 1
		} else {
			for i := 0; i < len(level); i++ {
				removed := removeIndex(level, i)
				s, _ := isLevelSafe(removed)
				if s {
					safe += 1
				}
			}
		}
	}

	fmt.Println(safe)
}
