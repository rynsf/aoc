package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLevelSafe(level []int) bool {
	if len(level) < 2 {
		return true
	}
	ascending := isAscending(level[0], level[1])
	for i := 0; i < len(level)-1; i++ {
		if !(isGradual(level[i], level[i+1]) && isAscending(level[i], level[i+1]) == ascending) {
			return false
		}
	}
	return true
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
	f, err := os.Open("./input")
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
		if isLevelSafe(level) {
			safe += 1
		}
	}

	fmt.Println(safe)
}
