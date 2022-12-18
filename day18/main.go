package main

import (
	"fmt"
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2022/helpers"
)

func main() {
	start := time.Now()
	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	//c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	inp := c.ReadInputFile()
	cubes := map[string]bool{}
	sum := 0
	for _, v := range inp {
		cubes[v] = true
	}
	for k := range cubes {
		c := c.GetInts(strings.Split(k, ","))
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0]-1, c[1], c[2])]; !ok {
			sum++
		}
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0]+1, c[1], c[2])]; !ok {
			sum++
		}
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0], c[1]-1, c[2])]; !ok {
			sum++
		}
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0], c[1]+1, c[2])]; !ok {
			sum++
		}
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0], c[1], c[2]-1)]; !ok {
			sum++
		}
		if _, ok := cubes[fmt.Sprintf("%d,%d,%d", c[0], c[1], c[2]+1)]; !ok {
			sum++
		}
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
