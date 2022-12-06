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
func overlap(str1, str2 string) rune {
	for _, v := range str1 {
		for _, b := range str2 {
			if v == b {
				return v
			}
		}
	}
	return 0
}
func part1() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range inp {
		x := overlap(v[:len(v)/2], v[len(v)/2:])
		if 'A' <= x && x <= 'Z' {
			sum += int(x-'A') + 27
		} else {
			sum += int(x-'a') + 1
		}
	}

	return sum
}
func overlap3(str1, str2, str3 string) rune {
	for _, v := range str1 {
		for _, b := range str2 {
			if v == b {
				for _, n := range str3 {
					if v == n {
						return n
					}
				}
			}
		}
	}
	return 0
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for i := 0; i < len(inp); i += 3 {
		x := overlap3(inp[i], inp[i+1], inp[i+2])
		if 'A' <= x && x <= 'Z' {
			sum += int(x-'A') + 27
		} else {
			sum += int(x-'a') + 1
		}
	}

	return sum
}
