package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Vector struct {
	x, y int
}

type ClawMachine struct {
	a, b, price Vector
}

func price(m ClawMachine) int {
	x := float64(m.a.x)
	y := float64(m.b.x)
	u := float64(m.a.y)
	v := float64(m.b.y)
	n := float64(m.price.x)
	o := float64(m.price.y)

	a := (y*o - v*n) / (y*u - x*v)
	b := (n - x*a) / y

	if (a != math.Trunc(a)) || (b != math.Trunc(b)) {
		return 0
	}
	return int((3 * a) + b)
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`(\d+)`)
	sc := bufio.NewScanner(f)

	var clawMachines []ClawMachine
	var c ClawMachine
	for sc.Scan() {
		if sc.Text() == "" {
			clawMachines = append(clawMachines, c)
			continue
		}
		nums := re.FindAllString(sc.Text(), -1)
		c.a.x, _ = strconv.Atoi(nums[0])
		c.a.y, _ = strconv.Atoi(nums[1])
		sc.Scan()
		nums = re.FindAllString(sc.Text(), -1)
		c.b.x, _ = strconv.Atoi(nums[0])
		c.b.y, _ = strconv.Atoi(nums[1])
		sc.Scan()
		nums = re.FindAllString(sc.Text(), -1)
		c.price.x, _ = strconv.Atoi(nums[0])
		c.price.x += 10000000000000
		c.price.y, _ = strconv.Atoi(nums[1])
		c.price.y += 10000000000000
	}

	totalCost := 0
	for _, m := range clawMachines {
		totalCost += price(m)
	}
	fmt.Println(totalCost)
}
