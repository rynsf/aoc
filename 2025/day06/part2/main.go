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

func parseOperands(strArr []string) [][]int {
	operands := [][]int{}
	op := []int{}
	for i := 0; i < len(strArr[0]); i++ {
		num := ""
		for n := 0; n < len(strArr); n++ {
			if strArr[n][i] != ' ' {
				num = num + string(strArr[n][i])
			}
		}
		if num == "" {
			operands = append(operands, op)
			op = []int{}
		} else {
			op = append(op, atoi(num))
		}
	}
	if len(op) != 0 {
		operands = append(operands, op)
	}
	return operands
}

func addArr(arr []int) int {
	acc := 0
	for _, n := range arr {
		acc += n
	}
	return acc
}

func mulArr(arr []int) int {
	acc := 1
	for _, n := range arr {
		acc *= n
	}
	return acc
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	strNums := []string{}

	for sc.Scan() {
		line := sc.Text()
		strNums = append(strNums, line)
	}

	operands := parseOperands(strNums[:len(strNums)-1])
	operator := strings.Fields(strNums[len(strNums)-1])

	total := 0

	for i := 0; i < len(operator); i++ {
		switch operator[i] {
		case "+":
			total += addArr(operands[i])
		case "*":
			total += mulArr(operands[i])
		}
	}

	fmt.Println(total)
}
