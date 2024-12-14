package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func endPos(r Robot) Vector {
	newPos := Vector{r.pos.x, r.pos.y}
	for i := 0; i < 100; i++ {
		newPos.x = mod((newPos.x + r.vel.x), w)
		newPos.y = mod((newPos.y + r.vel.y), h)
	}
	return newPos
}

func whichQuad(pos Vector) int {
	midx := w / 2
	midy := h / 2
	if pos.x < midx && pos.y < midy {
		return 0
	}
	if pos.x > midx && pos.y < midy {
		return 1
	}
	if pos.x < midx && pos.y > midy {
		return 2
	}
	if pos.x > midx && pos.y > midy {
		return 3
	}
	return 4
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

	diffQuads := [5]int{0, 0, 0, 0, 0}
	for _, r := range robots {
		p := endPos(r)
		diffQuads[whichQuad(p)] += 1
	}

	sum := 1
	for _, n := range diffQuads[:4] {
		sum *= n
	}
	fmt.Println(sum)
}
