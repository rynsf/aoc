package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isPaperRoll(room [][]int, x, y int) bool {
	leny := len(room)
	lenx := len(room[0])

	if x < 0 || y < 0 || x >= lenx || y >= leny {
		return false
	}
	return room[y][x] == 1
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	var room [][]int
	for sc.Scan() {
		line := sc.Text()

		l := make([]int, len(line))
		for i, c := range line {
			if c == '@' {
				l[i] = 1
			}
		}

		room = append(room, l)
	}

	total := 0

	for i, line := range room {
		for j, block := range line {
			if block != 1 {
				continue
			}

			countRolls := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if y == 0 && x == 0 {
						continue
					}
					if isPaperRoll(room, j+y, i+x) {
						countRolls += 1
					}
				}
			}

			if countRolls < 4 {
				total += 1
			}

		}
	}

	fmt.Println(total)
}
