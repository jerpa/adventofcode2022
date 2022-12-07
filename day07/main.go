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
	f := readFiles()

	sum := 0
	sum = smallFiles(f["/"].items)

	return sum
}

func part2() int {
	f := readFiles()

	need := 30000000 - (70000000 - f["/"].size)
	best := 70000000

	sizes := bestDir(f["/"].items)
	for _, v := range sizes {
		if v >= need && v < best {
			best = v
		}
	}

	return best
}
func smallFiles(m map[string]item) int {
	size := 0
	for k := range m {
		if m[k].isDir {
			size += smallFiles(m[k].items)
		}
	}
	for k := range m {
		if m[k].isDir && m[k].size <= 100000 {
			size += m[k].size
		}
	}
	return size
}
func bestDir(m map[string]item) []int {
	res := []int{}
	for k := range m {
		if m[k].isDir {
			res = append(res, bestDir(m[k].items)...)
		}
	}
	for k := range m {
		if m[k].isDir {
			res = append(res, m[k].size)
		}
	}
	return res
}

func readFiles() map[string]item {
	inp := c.ReadInputFile()
	items := map[string]item{"/": item{name: "/", isDir: true, size: 0, items: map[string]item{}}}
	curItem := items["/"]
	for _, v := range inp {
		s := strings.Split(v, " ")
		if s[0] == "$" {
			if s[1] == "cd" {
				if s[2] == "/" {
					curItem = items["/"]
				} else if s[2] == ".." {
					curItem = *curItem.parent
				} else {
					curItem = curItem.items[s[2]]
				}
			}

		} else {
			p := curItem
			i := item{name: s[1], parent: &p}
			if s[0] == "dir" {
				i.isDir = true
				i.items = map[string]item{}
			} else {
				i.size = c.GetInt(s[0])
			}
			curItem.items[s[1]] = i
		}
	}

	k := items["/"]
	k.size = calcSize(k.items)

	items["/"] = k
	return items
}
func calcSize(m map[string]item) int {
	for k := range m {
		if m[k].isDir {
			j := m[k]
			j.size = calcSize(j.items)
			m[k] = j
		}
	}
	size := 0
	for k := range m {
		size += m[k].size
	}
	return size
}

type item struct {
	name   string
	isDir  bool
	size   int
	items  map[string]item
	parent *item
}
