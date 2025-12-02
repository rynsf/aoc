package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func checkEqual(s string, step int) bool {
	if len(s)%step != 0 {
		return false
	}

	substr := s[:step]
	for i := 0; i+step <= len(s); i += step {
		if substr != s[i:i+step] {
			return false
		}
	}

	return true
}

func checkStepEqual(s string) bool {
	for step := 1; step < len(s); step++ {
		if checkEqual(s, step) {
			return true
		}
	}
	return false
}

func main() {
	bfile, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	line := strings.TrimSpace(string(bfile))
	strRanges := strings.Split(line, ",")

	ranges := [][]string{}

	for _, r := range strRanges {
		strNums := strings.Split(r, "-")
		ranges = append(ranges, strNums)
	}

	sum := 0

	for _, r := range ranges {
		low := atoi(r[0])
		high := atoi(r[1])

		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)

			if checkStepEqual(s) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}
