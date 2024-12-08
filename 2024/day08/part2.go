package main

import (
	"bufio"
	"fmt"
	"os"
)

func outOfBound(x, y, w, h int) bool {
	if x < 0 || y < 0 || x >= w || y >= h {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	var aMap []string
	for sc.Scan() {
		aMap = append(aMap, sc.Text())
	}
	width := len(aMap[0])
	height := len(aMap)

	antinodes := make([][]int, height)
	for i := range antinodes {
		antinodes[i] = make([]int, width)
	}

	for ay, line := range aMap {
		for ax, f1 := range line {
			if f1 != '.' {
				for by, l := range aMap {
					for bx, f2 := range l {
						if f1 == f2 && ay != by && ax != bx {
							deltaX := bx - ax
							deltaY := by - ay
							for x, y := ax, ay; !outOfBound(x, y, width, height); x, y = x+deltaX, y+deltaY {
								antinodes[y][x] = 1
							}
						}
					}
				}
			}
		}
	}

	for _, line := range aMap {
		fmt.Println(line)
	}

	fmt.Println()

	for _, line := range antinodes {
		for _, n := range line {
			if n == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	result := 0
	for _, line := range antinodes {
		for _, n := range line {
			if n == 1 {
				result += n
			}
		}
	}
	fmt.Println(result)
}
