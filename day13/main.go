package main

import (
	"encoding/json"
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

type pair struct {
	first  []any
	second []any
}

func parseRow(s string) []any {
	var data []any
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		c.Print(err.Error())
	}
	return data
}

func readFile() []pair {
	inp := c.ReadInputFile()
	res := []pair{}
	p := pair{}
	var l []any
	for _, v := range inp {
		if len(v) == 0 {
			res = append(res, p)
			p = pair{}
			continue
		}
		l = parseRow(v)

		if p.first == nil {
			p.first = l
		} else {
			p.second = l
		}
	}
	res = append(res, p)
	return res
}

type trool int

const (
	false trool = iota
	unknown
	true
)

func compare(first, second []any) trool {
	for i := range first {
		if len(second) <= i {
			return false
		}
		if _, ok := first[i].(float64); ok {
			if _, ok := second[i].(float64); ok {
				if second[i].(float64) < first[i].(float64) {
					return false
				} else if first[i].(float64) < second[i].(float64) {
					return true
				}
			} else {
				v := compare([]any{first[i]}, second[i].([]any))
				if v != unknown {
					return v
				}
			}
		} else if _, ok := second[i].(float64); ok {
			v := compare(first[i].([]any), []any{second[i]})
			if v != unknown {
				return v
			}
		} else {
			v := compare(first[i].([]any), second[i].([]any))
			if v != unknown {
				return v
			}
		}
	}
	if len(first) < len(second) {
		return true
	}
	return unknown
}

func part1() int {
	p := readFile()
	sum := 0
	for i := range p {
		if compare(p[i].first, p[i].second) == true {
			sum += i + 1
		}
	}
	return sum
}
func part2() int {
	inp := c.ReadInputFile()
	data := []string{"[[2]]", "[[6]]"}
	sum := 0
	for _, v := range inp {
		if len(v) > 0 {
			data = append(data, v)
		}
	}
	sort.Slice(data, func(i, j int) bool {
		return compare(parseRow(data[i]), parseRow(data[j])) == true
	})
	for i, v := range data {
		if v == "[[2]]" {
			sum = i + 1
		}
		if v == "[[6]]" {
			sum *= (i + 1)
		}
	}

	return sum
}
