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
func readMap() (map[int]map[int]bool, int) {
	inp := c.ReadInputFile()
	res := map[int]map[int]bool{}
	maxY := 0
	for _, v := range inp {
		p := strings.Split(v, " -> ")
		for i := 0; i < len(p)-1; i++ {
			start := strings.Split(p[i], ",")
			stop := strings.Split(p[i+1], ",")
			x := c.GetInt(start[0])
			y := c.GetInt(start[1])
			targetX := c.GetInt(stop[0])
			targetY := c.GetInt(stop[1])
			for {
				if _, ok := res[y]; !ok {
					res[y] = map[int]bool{}
				}
				res[y][x] = true
				if x == targetX && y == targetY {
					break
				}
				if x < targetX {
					x++
				} else if x > targetX {
					x--
				} else if y < targetY {
					y++
				} else {
					y--
				}
				if y > maxY {
					maxY = y
				}

			}
		}
	}
	return res, maxY
}

func part1() int {
	m, maxY := readMap()

	sum := 0
	for {
		x := 500
		y := 0
		for {
			if _, ok := m[y+1]; ok {
				if !m[y+1][x] {
					y++
				} else if !m[y+1][x-1] {
					y++
					x--
				} else if !m[y+1][x+1] {
					y++
					x++
				} else {
					if _, ok := m[y]; !ok {
						m[y] = map[int]bool{}
					}
					m[y][x] = true
					x = 500
					y = 0
					sum++
				}
			} else {
				y++
			}
			if y > maxY {
				return sum
			}
		}
	}

	return sum
}
func part2() int {
	m, maxY := readMap()
	m[maxY+2] = map[int]bool{}

	sum := 0
	for {
		x := 500
		y := 0
		for {
			if _, ok := m[y+1]; ok {
				if y == maxY+1 {
					m[y+1][x-1] = true
					m[y+1][x] = true
					m[y+1][x+1] = true
				}
				if !m[y+1][x] {
					y++
				} else if !m[y+1][x-1] {
					y++
					x--
				} else if !m[y+1][x+1] {
					y++
					x++
				} else {
					if _, ok := m[y]; !ok {
						m[y] = map[int]bool{}
					}
					m[y][x] = true
					sum++
					if x == 500 && y == 0 {
						return sum
					}
					x = 500
					y = 0
				}
			} else {
				y++
			}
		}
	}

	return sum
}
