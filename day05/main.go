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
	crates := map[int][]string{}
	inp := c.ReadInputFile()

	row := 0
	for _, v := range inp {
		if strings.Contains(v, "[") {
			for i := 0; i*4 < len(v); i++ {
				if v[i*4+1] == ' ' {
					continue
				}
				if _, ok := crates[i]; !ok {
					crates[i] = []string{}
				}
				crates[i] = append([]string{string(v[i*4+1])}, crates[i]...)
			}

		} else {
			break
		}
		row++
	}
	row += 2
	for ; row < len(inp); row++ {
		s := strings.Split(inp[row], " ")
		num := c.GetInt(s[1])
		start := c.GetInt(s[3]) - 1
		stop := c.GetInt(s[5]) - 1
		for ; num > 0; num-- {
			crates[stop] = append(crates[stop], crates[start][len(crates[start])-1])
			crates[start] = crates[start][:len(crates[start])-1]
		}
	}
	res := ""
	for i := 0; i < len(crates); i++ {
		res += crates[i][len(crates[i])-1]
	}
	c.Print(res)
	return 0
}
func part2() int {
	crates := map[int][]string{}
	inp := c.ReadInputFile()

	row := 0
	for _, v := range inp {
		if strings.Contains(v, "[") {
			for i := 0; i*4 < len(v); i++ {
				if v[i*4+1] == ' ' {
					continue
				}
				if _, ok := crates[i]; !ok {
					crates[i] = []string{}
				}
				crates[i] = append([]string{string(v[i*4+1])}, crates[i]...)
			}

		} else {
			break
		}
		row++
	}
	row += 2
	for ; row < len(inp); row++ {
		s := strings.Split(inp[row], " ")
		num := c.GetInt(s[1])
		start := c.GetInt(s[3]) - 1
		stop := c.GetInt(s[5]) - 1
		crates[stop] = append(crates[stop], crates[start][len(crates[start])-num:]...)
		crates[start] = crates[start][:len(crates[start])-num]

	}
	res := ""
	for i := 0; i < len(crates); i++ {
		res += crates[i][len(crates[i])-1]
	}
	c.Print(res)
	return 0
}
