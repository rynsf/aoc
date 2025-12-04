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

// returns the number of paper lifted
func liftPaper(room [][]int) int {
	removed := 0
	toRemove := [][]int{}

	for j, line := range room {
		for i, block := range line {
			if block != 1 {
				continue
			}

			countRolls := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if y == 0 && x == 0 {
						continue
					}
					if isPaperRoll(room, i+x, j+y) {
						countRolls += 1
					}
				}
			}

			if countRolls < 4 {
				removed += 1
				toRemove = append(toRemove, []int{i, j})
			}
		}
	}

	for _, r := range toRemove {
		room[r[1]][r[0]] = 0
	}

	return removed
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

	for {
		lifted := liftPaper(room)
		total += lifted
		if lifted == 0 {
			break
		}
	}

	fmt.Println(total)
}
