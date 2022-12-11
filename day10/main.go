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
	x := 1
	cycle := 1
	next := 20
	for _, v := range inp {
		if next == cycle {
			sum += cycle * x
			next += 40
		}
		s := strings.Split(v, " ")
		if s[0] == "noop" {
			cycle++
		} else {
			cycle++
			if next == cycle {
				sum += cycle * x
				next += 40

			}
			cycle++
			x += c.GetInt(s[1])
		}
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	//sum := 0
	x := 2
	cycle := 1
	next := 40
	res := ""
	for _, v := range inp {
		if cycle == next {
			x += 40
			next += 40
		}
		if cycle >= x-1 && cycle <= x+1 {
			res += "#"
		} else {
			res += "."
		}
		s := strings.Split(v, " ")
		if s[0] == "noop" {
			cycle++
		} else {
			cycle++
			if cycle >= x-1 && cycle <= x+1 {
				res += "#"
			} else {
				res += "."
			}
			if cycle == next {
				x += 40
				next += 40
			}
			cycle++
			x += c.GetInt(s[1])
		}
	}
	for i := 0; i < len(res); i += 40 {
		c.Print(res[i : i+40])
	}

	return 0
}
