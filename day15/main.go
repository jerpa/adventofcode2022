package main

import (
	"math"
	"regexp"
	"sort"
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

type pos struct {
	X int
	Y int
}

func (p pos) Distance(o pos) int {
	return int(math.Abs(float64(p.X-o.X)) + math.Abs(float64(p.Y-o.Y)))
}

type sensor struct {
	pos    pos
	beacon pos
}

type line struct {
	start int
	stop  int
}

func (l line) merge(o line) ([]line, bool) {
	if l.stop >= o.start && l.start <= o.stop {
		return []line{line{start: int(math.Min(float64(l.start), float64(o.start))), stop: int(math.Max(float64(l.stop), float64(o.stop)))}}, true
	}
	return []line{l, o}, false
}

func readFile() []sensor {
	inp := c.ReadInputFile()
	res := []sensor{}
	r := regexp.MustCompile(`x=([\-\d]+), y=([\-\d]+)`)
	for _, v := range inp {
		m := r.FindAllStringSubmatch(v, -1)
		res = append(res, sensor{pos: pos{X: c.GetInt(m[0][1]), Y: c.GetInt(m[0][2])}, beacon: pos{X: c.GetInt(m[1][1]), Y: c.GetInt(m[1][2])}})
	}
	return res
}

func part1() int {
	data := readFile()

	target := 2000000
	a := []line{}
	sum := 0
	for _, v := range data {
		d := v.pos.Distance(v.beacon)
		if math.Abs(float64(target-v.pos.Y)) < float64(d) {
			l := d - int(math.Abs(float64(target-v.pos.Y)))
			a = append(a, line{start: v.pos.X - l, stop: v.pos.X + l})
		}
	}
	for {
		m := []line{}
		t := false
		for i := 0; i < len(a)-1; i++ {
			if r, ok := a[i].merge(a[i+1]); ok {
				m = append(m, r...)
				t = true
				i++
			} else {
				m = append(m, a[i])
			}
		}
		a = m
		if !t || len(a) == 1 {
			break
		}
	}
	for _, v := range a {
		sum += v.stop - v.start + 1
	}
	p := map[int]bool{}
	for _, v := range data {
		if v.beacon.Y == target {
			for _, l := range a {
				if v.beacon.X >= l.start && v.beacon.X <= l.stop {
					p[v.beacon.X] = true
				}
			}
		}
		if v.pos.Y == target {
			for _, l := range a {
				if v.pos.X >= l.start && v.pos.X <= l.stop {
					p[v.pos.X] = true
				}
			}
		}
	}
	sum -= len(p)

	return sum
}
func part2() int {
	data := readFile()
	var l1 int
	var l2 int
	var a []line
	m := []line{}
	t := false

	for target := 0; target <= 4000000; target++ {
		a = []line{}

		for _, v := range data {
			d := v.pos.Distance(v.beacon)
			if math.Abs(float64(target-v.pos.Y)) < float64(d) {
				l1 = d - int(math.Abs(float64(target-v.pos.Y)))
				l2 = v.pos.X + l1
				l1 = v.pos.X - l1
				if l1 < 0 {
					l1 = 0
				}
				if l2 >= 4000000 {
					l2 = 4000000
				}

				a = append(a, line{start: l1, stop: l2})
			}
		}
		for {
			m = []line{}
			t = false
			for i := 0; i < len(a)-1; i++ {
				if r, ok := a[i].merge(a[i+1]); ok {
					m = append(m, r...)
					t = true
					i++
				} else {
					m = append(m, a[i])
				}
			}
			a = m
			if !t || len(a) == 1 {
				break
			}
		}
		if len(a) != 1 {
			sort.Slice(a, func(i, j int) bool {
				return a[i].stop < a[j].start
			})
			return (a[0].stop+1)*4000000 + target
		}

	}
	return 0
}
