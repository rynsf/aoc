package main

import (
	"bufio"
	"fmt"
	"os"
)

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
							// calculate coordinate of antinodes
							// deltaX := (bx - ax) * 2
							// deltaY := (by - ay) * 2
							// antiX := ax + deltaX
							// antiY := ay + deltaY
							// b + (b - y)
							antiX := bx + bx - ax
							antiY := by + by - ay
							// check if it is outofboard,
							if !(antiX < 0 || antiY < 0 || antiX >= width || antiY >= height) {
								antinodes[antiY][antiX] = 1
							}
							// frequencies are same,
							// draw a anti nodes,
							// // calculate the location where to add anti node, by doubling the distance between point a and b,
							// // then set the bit at the same coordinate in antinode array
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
