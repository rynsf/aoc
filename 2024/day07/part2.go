package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	x    int
	nums []int
}

func pow(x, y int) int {
	result := math.Pow(float64(x), float64(y))
	return int(result)
}

func evalEq(nums []int, op int) int {
	result := nums[0]
	for _, n := range nums[1:] {
		switch op % 3 {
		case 2:
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(n))
		case 1:
			result *= n
		case 0:
			result += n
		}
		op /= 3
	}
	return result
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)

	var eq []equation
	for sc.Scan() {
		s := strings.Fields(sc.Text())
		x, err := strconv.Atoi((s[0])[:len(s[0])-1])
		if err != nil {
			panic(err)
		}
		nums := []int{}
		for _, n := range s[1:] {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		eq = append(eq, equation{x, nums})
	}

	result := 0
	for _, e := range eq {
		// 0 is add 1 is multiply
		// loop over all the permulation of operations, this can be done by having a loop that goes from 0 to number of operand ^ 2 - 1.
		// // evalute the expresion, if it is equal to x, then accumulate
		for op := 0; op < pow(3, len(e.nums)); op++ {
			x := evalEq(e.nums, op)
			if x == e.x {
				result += x
				break
			}
		}
	}
	fmt.Println(result)
}
