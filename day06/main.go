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
	return findFirstUnique(4)
}
func part2() int {
	return findFirstUnique(14)
}
func findFirstUnique(length int) int {
	inp := c.ReadInputFile()

	sum := 0
	l := inp[0]
	for i := length; i < len(l); i++ {
		s := l[i-length : i]
		sum = 0
		for _, o := range s {
			for _, u := range s {
				if o == u {
					sum++
				}
			}
		}
		if sum == length {
			return i
		}
	}

	return -1
}
