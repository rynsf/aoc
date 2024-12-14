package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Vector struct {
	x, y int
}

type Robot struct {
	pos Vector
	vel Vector
}

const w = 101
const h = 103

func mod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func nextPos(r Robot) Vector {
	newPos := Vector{r.pos.x, r.pos.y}
	newPos.x = mod((newPos.x + r.vel.x), w)
	newPos.y = mod((newPos.y + r.vel.y), h)
	return newPos
}

func updateRobots(robots []Robot) {
	for i := range robots {
		robots[i].pos = nextPos(robots[i])
	}
}

func drawFrame(robots []Robot) [][]int {
	frame := makeIntSlice(w, h)
	for i := range robots {
		frame[robots[i].pos.y][robots[i].pos.x] += 1
	}
	return frame
}

func renderFrame(frame [][]int) {
	for _, l := range frame {
		for _, c := range l {
			if c == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func outOfBound(x, y, w, h int) bool {
	if x < 0 || y < 0 || x >= w || y >= h {
		return true
	}
	return false
}

func watershed(x, y int, frame, visited [][]int) int {
	if outOfBound(x, y, len(frame[0]), len(frame)) {
		return 0
	}
	if frame[y][x] == 0 {
		return 0
	}
	if visited[y][x] == 1 {
		return 0
	}
	visited[y][x] = 1
	size := 1
	size += watershed(x-1, y-1, frame, visited)
	size += watershed(x, y-1, frame, visited)
	size += watershed(x+1, y-1, frame, visited)

	size += watershed(x-1, y, frame, visited)
	size += watershed(x+1, y, frame, visited)

	size += watershed(x-1, y+1, frame, visited)
	size += watershed(x, y+1, frame, visited)
	size += watershed(x+1, y+1, frame, visited)

	return size
}

func largestConnectedRegion(frame [][]int) int {
	var connectedSize []int
	visited := makeIntSlice(w, h)
	for y, line := range frame {
		for x, char := range line {
			if char > 0 && visited[y][x] == 0 {
				connectedSize = append(connectedSize, watershed(x, y, frame, visited))
			}
		}
	}

	return slices.Max(connectedSize)
}

func makeIntSlice(x, y int) [][]int {
	s := make([][]int, y)
	for i := range s {
		s[i] = make([]int, x)
	}
	return s
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`-?\d+`)
	sc := bufio.NewScanner(f)

	var robots []Robot
	for sc.Scan() {
		nums := re.FindAllString(sc.Text(), -1)
		num := []int{}
		for _, n := range nums {
			i, _ := strconv.Atoi(n)
			num = append(num, i)
		}
		robots = append(robots, Robot{Vector{num[0], num[1]}, Vector{num[2], num[3]}})
	}

	frame := makeIntSlice(w, h)
	for i := 0; i < 1000000; i++ {
		frame = drawFrame(robots)
		c := largestConnectedRegion(frame)
		if c > 50 {
			renderFrame(frame)
			fmt.Println("Seconds: ", i)
			return
		}
		updateRobots(robots)
	}
}
