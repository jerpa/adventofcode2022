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
		s := strings.Split(v, ",")
		a := strings.Split(s[0], "-")
		b := strings.Split(s[1], "-")
		if (c.GetInt(a[0]) <= c.GetInt(b[0]) && c.GetInt(b[1]) <= c.GetInt(a[1])) || (c.GetInt(b[0]) <= c.GetInt(a[0]) && c.GetInt(a[1]) <= c.GetInt(b[1])) {
			sum++
		}
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range inp {
		s := strings.Split(v, ",")
		a := strings.Split(s[0], "-")
		b := strings.Split(s[1], "-")
		if (c.GetInt(a[0]) <= c.GetInt(b[1]) && c.GetInt(b[0]) <= c.GetInt(a[1])) || (c.GetInt(b[0]) <= c.GetInt(a[1]) && c.GetInt(a[0]) <= c.GetInt(b[1])) {
			sum++
		}
	}

	return sum
}
