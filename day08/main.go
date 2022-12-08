package main

import (
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
	sum += len(inp) * 2
	sum += (len(inp[0]) - 2) * 2
	for y := 1; y < len(inp)-1; y++ {
		for x := 1; x < len(inp[y])-1; x++ {
			visible := true
			for d := x + 1; d < len(inp[y]); d++ {
				if c.GetInt(string(inp[y][d])) >= c.GetInt(string(inp[y][x])) {
					visible = false
					break
				}
			}
			if visible {
				sum++
				continue
			}
			visible = true
			for d := x - 1; d >= 0; d-- {
				if c.GetInt(string(inp[y][d])) >= c.GetInt(string(inp[y][x])) {
					visible = false
					break
				}
			}
			if visible {
				sum++
				continue
			}
			visible = true
			for d := y + 1; d < len(inp); d++ {
				if c.GetInt(string(inp[d][x])) >= c.GetInt(string(inp[y][x])) {
					visible = false
					break
				}
			}
			if visible {
				sum++
				continue
			}
			visible = true
			for d := y - 1; d >= 0; d-- {
				if c.GetInt(string(inp[d][x])) >= c.GetInt(string(inp[y][x])) {
					visible = false
					break
				}
			}
			if visible {
				sum++
				continue
			}
		}
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for y := 1; y < len(inp)-1; y++ {
		for x := 1; x < len(inp[y])-1; x++ {
			r := 1
			v := 1
			match := false
			for d := x + 1; d < len(inp[y]); d++ {
				if c.GetInt(string(inp[y][d])) >= c.GetInt(string(inp[y][x])) {
					r *= v
					match = true
					break
				}
				v++
			}
			if !match {
				r *= (v - 1)
			}
			v = 1
			match = false
			for d := x - 1; d >= 0; d-- {
				if c.GetInt(string(inp[y][d])) >= c.GetInt(string(inp[y][x])) {
					r *= v
					match = true
					break
				}
				v++
			}
			if !match {
				r *= (v - 1)
			}
			v = 1
			match = false
			for d := y + 1; d < len(inp); d++ {
				if c.GetInt(string(inp[d][x])) >= c.GetInt(string(inp[y][x])) {
					r *= v
					match = true
					break
				}
				v++
			}
			if !match {
				r *= (v - 1)
			}
			v = 1
			match = false
			for d := y - 1; d >= 0; d-- {
				if c.GetInt(string(inp[d][x])) >= c.GetInt(string(inp[y][x])) {
					r *= v
					match = true
					break
				}
				v++
			}
			if !match {
				r *= (v - 1)
			}
			if r > sum {
				sum = r
			}
		}
	}

	return sum
}
