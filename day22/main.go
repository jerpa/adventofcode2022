package main

import (
	"regexp"
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

	for i, v := range inp {
		if len(v) == 0 {
			break
		}
		inp[i] += strings.Repeat(" ", 100)
	}
	y := 0
	x := 0
	d := 0 // 0=right, 1=down, 2=left, 3=up
	for inp[y][x] != '.' {
		x++
	}
	reg := regexp.MustCompile(`(\d+|[LR])`)
	instr := reg.FindAllString(inp[len(inp)-1], -1)
	for _, s := range instr {
		if s == "L" {
			d--
			if d < 0 {
				d += 4
			}
		} else if s == "R" {
			d++
			d %= 4
		} else {
			steps := c.GetInt(s)
			for steps > 0 {
				newX := x
				newY := y
				steps--
				if d == 0 {
					newX++
					if newX == len(inp[y]) || inp[y][newX] == ' ' {
						newX = 0
						for inp[y][newX] == ' ' {
							newX++
						}
					}
				} else if d == 2 {
					newX--
					if newX == -1 || inp[y][newX] == ' ' {
						newX = len(inp[y]) - 1
						for inp[y][newX] == ' ' {
							newX--
						}
					}
				} else if d == 1 {
					newY++
					if newY == len(inp)-2 || inp[newY][x] == ' ' {
						newY = 0
						for inp[newY][x] == ' ' {
							newY++
						}
					}
				} else if d == 3 {
					newY--
					if newY == -1 || inp[newY][x] == ' ' {
						newY = len(inp) - 3
						for inp[newY][x] == ' ' {
							newY--
						}
					}
				}
				if inp[newY][newX] == '#' {
					break
				}
				y = newY
				x = newX
			}
		}
	}

	return (1000 * (y + 1)) + (4 * (x + 1)) + d
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
