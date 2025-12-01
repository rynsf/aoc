package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type operations struct {
	sign int
	num  int
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	ops := []operations{}

	for sc.Scan() {
		sign := 1
		line := sc.Text()
		fmt.Println(line)

		if line[0] == 'L' {
			sign = -1
		}

		turns, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		ops = append(ops, operations{sign, turns})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	dial := 50
	passwd := 0
	for _, o := range ops {
		for i := 0; i < o.num; i++ {
			dial += o.sign

			switch dial {
			case 100:
				dial = 0
			case -1:
				dial = 99
			}

			if dial == 0 {
				passwd += 1
			}
		}
	}

	fmt.Println(passwd)
}
