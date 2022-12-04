package main

import (
	"sort"
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

	max := 0
	sum := 0
	for _, v := range inp {
		if v == "" {

			sum = 0
		} else {
			sum += c.GetInt(v)
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
func part2() int {
	inp := c.ReadInputFile()

	nums := []int{}
	sum := 0
	for _, v := range inp {
		if v == "" {
			nums = append(nums, sum)
			sum = 0
		} else {
			sum += c.GetInt(v)
		}
	}

	sort.Ints(nums)

	return nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
}
