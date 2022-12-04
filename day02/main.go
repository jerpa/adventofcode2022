package main

import (
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2022/helpers"
)

func main() {
	start := time.Now()
	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range inp {
		a := strings.Split(v, " ")
		sum += calcPoint(a[0], a[1])
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range inp {
		a := strings.Split(v, " ")
		sum += choose(a[0], a[1])
	}

	return sum
}
func choose(opp, res string) int {
	switch res {
	case "X":
		if opp == "A" {
			return calcPoint(opp, "Z")
		}
		if opp == "B" {
			return calcPoint(opp, "X")
		}
		if opp == "C" {
			return calcPoint(opp, "Y")
		}
	case "Y":
		if opp == "A" {
			return calcPoint(opp, "X")
		}
		if opp == "B" {
			return calcPoint(opp, "Y")
		}
		if opp == "C" {
			return calcPoint(opp, "Z")
		}
	case "Z":
		if opp == "A" {
			return calcPoint(opp, "Y")
		}
		if opp == "B" {
			return calcPoint(opp, "Z")
		}
		if opp == "C" {
			return calcPoint(opp, "X")
		}
	}
	return 0
}

func calcPoint(opp, me string) int {
	sum := 0
	switch me {
	case "X":
		sum += 1
	case "Y":
		sum += 2
	case "Z":
		sum += 3
	}
	if opp == "A" && me == "Y" {
		return sum + 6
	}
	if opp == "B" && me == "Z" {
		return sum + 6
	}
	if opp == "C" && me == "X" {
		return sum + 6
	}
	if (opp == "A" && me == "X") || (opp == "B" && me == "Y") || (opp == "C" && me == "Z") {
		return sum + 3
	}
	return sum
}
