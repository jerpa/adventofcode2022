package main

import (
	"sort"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
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

type monkey struct {
	items       []int
	operation   string
	test        int
	targetTrue  int
	targetFalse int
	inspections int
}

func readMonkeyFile() []monkey {
	inp := c.ReadInputFile()
	monkeys := []monkey{}
	for i := 0; i < len(inp); i += 7 {
		m := monkey{}
		s := strings.Split(inp[i+1], ": ")
		m.items = c.GetInts(strings.Split(s[1], ", "))
		s = strings.Split(inp[i+2], "= ")
		m.operation = s[1]
		s = strings.Split(inp[i+3], " ")
		m.test = c.GetInt(s[len(s)-1])
		s = strings.Split(inp[i+4], " ")
		m.targetTrue = c.GetInt(s[len(s)-1])
		s = strings.Split(inp[i+5], " ")
		m.targetFalse = c.GetInt(s[len(s)-1])
		monkeys = append(monkeys, m)
	}
	return monkeys
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part1() int {
	monkeys := readMonkeyFile()

	for i := 0; i < 20; i++ {
		for m := range monkeys {
			for o := range monkeys[m].items {
				monkeys[m].inspections++
				expr, _ := govaluate.NewEvaluableExpression(monkeys[m].operation)
				para := map[string]interface{}{"old": monkeys[m].items[o]}
				res, _ := expr.Evaluate(para)
				monkeys[m].items[o] = int(res.(float64)) / 3
				if monkeys[m].items[o]%monkeys[m].test == 0 {
					monkeys[monkeys[m].targetTrue].items = append(monkeys[monkeys[m].targetTrue].items, monkeys[m].items[o])
				} else {
					monkeys[monkeys[m].targetFalse].items = append(monkeys[monkeys[m].targetFalse].items, monkeys[m].items[o])
				}
			}
			monkeys[m].items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}
func part2() int {
	monkeys := readMonkeyFile()

	divs := []int{}
	for _, v := range monkeys {
		divs = append(divs, v.test)
	}
	sort.Ints(divs)
	lcm := LCM(divs[0], divs[1], divs[2:]...)

	for i := 0; i < 10000; i++ {
		for m := range monkeys {
			for o := range monkeys[m].items {
				monkeys[m].inspections++
				expr, _ := govaluate.NewEvaluableExpression(monkeys[m].operation)
				para := map[string]interface{}{"old": monkeys[m].items[o]}
				res, _ := expr.Evaluate(para)
				monkeys[m].items[o] = int(res.(float64)) % lcm
				if monkeys[m].items[o]%monkeys[m].test == 0 {
					monkeys[monkeys[m].targetTrue].items = append(monkeys[monkeys[m].targetTrue].items, monkeys[m].items[o])
				} else {
					monkeys[monkeys[m].targetFalse].items = append(monkeys[monkeys[m].targetFalse].items, monkeys[m].items[o])
				}
			}
			monkeys[m].items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}
