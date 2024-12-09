package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{1, 0},   // horizontal
	{-1, 0},  // horizontal reverse
	{0, 1},   // vertical
	{0, -1},  // vetical reverse
	{1, 1},   // bottom right
	{-1, -1}, // top left
	{1, -1},  // top right
	{-1, 1},  // bottom left
}

func OutOfBound(x, y, w, h int) bool {
	if x < 0 || y < 0 || x >= w || y >= h {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	count := 0
	flag := true
	for y, l := range lines {
		for x, c := range l {
			if c == 'X' {
				for _, d := range directions {
					flag = true
					i, j := x, y
					for _, s := range "MAS" {
						i = i + d[0]
						j = j + d[1]
						if OutOfBound(i, j, len(lines[0]), len(lines)) {
							flag = false
							break
						}
						if string(lines[j][i]) != string(s) {
							flag = false
							break
						}
					}
					if flag {
						count += 1
					}
				}
			}
		}
	}
	fmt.Println(count)
}
