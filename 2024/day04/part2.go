package main

import (
	"bufio"
	"fmt"
	"os"
)

func MScheck(a, b rune) bool {
	if (a == 'M' && b == 'S') || (a == 'S' && b == 'M') {
		return true
	}
	return false
}
func isXmas(br, tl, bl, tr rune) bool {
	if MScheck(br, tl) && MScheck(bl, tr) {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines [][]rune
	for sc.Scan() {
		lines = append(lines, []rune(sc.Text()))
	}

	count := 0
	for y, l := range lines {
		for x, c := range l {
			if c == 'A' {
				if x > 0 && y > 0 && x < len(lines[0])-1 && y < len(lines)-1 {
					if isXmas(lines[y+1][x+1], lines[y-1][x-1], lines[y-1][x+1], lines[y+1][x-1]) {
						count += 1
					}
				}
			}
		}
	}
	fmt.Println(count)
}
