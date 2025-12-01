package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(x, y int) int {
	return (x%y + y) % y
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	ops := []int{}

	for sc.Scan() {
		sign := 1
		line := sc.Text()
		if line[0] == 'L' {
			sign = -1
		}

		turns, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		ops = append(ops, turns*sign)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	dial := 50
	passwd := 0
	for _, o := range ops {
		dial = mod(dial+o, 100)
		if dial == 0 {
			passwd += 1
		}
	}

	fmt.Println(passwd)
}
