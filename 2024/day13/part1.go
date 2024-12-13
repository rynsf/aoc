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

type ClawMachine struct {
	a, b, price Vector
}

// takes in clawMachine return the least amount of token it will take to win, or zero if there is not way
func leastPrice(m ClawMachine) int {
	var possiblePrice []int
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			endSpotX := (m.a.x * a) + (m.b.x * b)
			endSpotY := (m.a.y * a) + (m.b.y * b)
			if endSpotX == m.price.x && endSpotY == m.price.y {
				possiblePrice = append(possiblePrice, (3*a)+(1*b))
			}
		}
	}
	if len(possiblePrice) == 0 {
		return 0
	}
	least := slices.Min(possiblePrice)
	return least
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
		c.price.y, _ = strconv.Atoi(nums[1])
	}

	fmt.Println(clawMachines)

	totalCost := 0
	for _, m := range clawMachines {
		totalCost += leastPrice(m)
	}
	fmt.Println(totalCost)
}
