package main

import (
	"bufio"
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

func parseOperands(strArr [][]string) [][]int {
	nums := [][]int{}

	for _, l := range strArr {
		arr := []int{}
		for _, n := range l {
			arr = append(arr, atoi(n))
		}
		nums = append(nums, arr)
	}

	return nums
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	strNums := [][]string{}

	for sc.Scan() {
		line := sc.Text()

		strNums = append(strNums, strings.Fields(line))

	}

	operands := parseOperands(strNums[:len(strNums)-1])
	operator := strNums[len(strNums)-1]

	total := 0

	for i, o := range operator {
		switch o {
		case "+":
			acc := 0
			for _, nums := range operands {
				acc += nums[i]
			}
			total += acc
		case "*":
			acc := 1
			for _, nums := range operands {
				acc *= nums[i]
			}
			total += acc
		}
	}

	fmt.Println(total)
}
