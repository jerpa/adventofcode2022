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
	//c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func overlap(rock []string, cave []string, x, y int) bool {
	for ry := range rock {
		for rx := range rock[ry] {
			if rock[ry][rx] == '#' && cave[y+ry][x+rx] == '#' {
				return true
			}
		}
	}
	return false
}

func part1() int {
	inp := c.ReadInputFile()
	cave := []string{"       ", "       ", "       ", "       ", "       ", "       ", "       ", "       ", "#######"}
	rocks := [][]string{[]string{"####"}, []string{" # ", "###", " # "}, []string{"  #", "  #", "###"}, []string{"#", "#", "#", "#"}, []string{"##", "##"}}
	top := 7
	rockX := 2
	rockY := 4
	rock := 0
	wind := inp[0]
	cnt := 0
	for {
		for _, w := range wind {
			if w == '>' {
				if rockX+len(rocks[rock][0]) < 7 && !overlap(rocks[rock], cave, rockX+1, rockY) {
					rockX++
				}
			} else {
				if rockX-1 >= 0 && !overlap(rocks[rock], cave, rockX-1, rockY) {
					rockX--
				}
			}
			if overlap(rocks[rock], cave, rockX, rockY+1) {
				for y := range rocks[rock] {
					for x := range rocks[rock][y] {
						if rocks[rock][y][x] == '#' {
							t := []rune(cave[rockY+y])
							t[rockX+x] = '#'
							cave[rockY+y] = string(t)
						}
					}
				}
				if top > rockY {
					top = rockY
				}
				cnt++
				if cnt == 2022 {
					return len(cave) - top - 1
				}
				for top < 7 {
					cave = append([]string{"       "}, cave...)
					top++
				}
				rockX = 2
				rock++
				rock %= 5
				rockY = 4 - len(rocks[rock])

			} else {
				rockY++
			}
		}
	}
	return 0
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
