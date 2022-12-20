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

type link struct {
	num   int
	index int
	prev  *link
	next  *link
}

func part1() int {
	inp := c.ReadInputFile()

	l := &link{num: c.GetInt(inp[0]), index: 0}
	first := l

	for i := 1; i < len(inp); i++ {
		nl := &link{num: c.GetInt(inp[i]), prev: l, index: i}
		l.next = nl
		l = nl
	}
	first.prev = l
	l.next = first

	for i, v := range inp {
		num := c.GetInt(v) % (len(inp) - 1)
		for l.index != i {
			l = l.next
		}
		if num < 0 {
			for i := 0; i > num; i-- {
				// [left.prev]--[l.prev = left]--[l]--[l.next = right]--[right.next]
				// [left.prev]--[l]--[left]--[right]--[right.next]

				right := l.next
				left := l.prev
				l.prev = left.prev
				l.next = left
				left.prev.next = l
				left.prev = l
				left.next = right
				right.prev = left
			}
		} else if num > 0 {
			for i := 0; i < num; i++ {
				// [left.prev]--[l.prev = left]--[l]--[l.next = right]--[right.next]
				// [left.prev]--[left]--[right]--[l]--[right.next]

				right := l.next
				left := l.prev
				l.prev = right
				l.next = right.next
				right.next.prev = l
				right.next = l
				right.prev = left
				left.next = right
			}
		}

	}
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 1000; i++ {
		l = l.next
	}
	v1 := l.num
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 2000; i++ {
		l = l.next
	}
	v2 := l.num
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 3000; i++ {
		l = l.next
	}
	v3 := l.num

	return v1 + v2 + v3
}
func part2() int {
	inp := c.ReadInputFile()

	l := &link{num: 811589153 * c.GetInt(inp[0]), index: 0}
	first := l

	for i := 1; i < len(inp); i++ {
		nl := &link{num: 811589153 * c.GetInt(inp[i]), prev: l, index: i}
		l.next = nl
		l = nl
	}
	first.prev = l
	l.next = first
	for o := 0; o < 10; o++ {
		for i, v := range inp {
			num := (811589153 * c.GetInt(v)) % (len(inp) - 1)
			for l.index != i {
				l = l.next
			}
			if num < 0 {
				for i := 0; i > num; i-- {
					// [left.prev]--[l.prev = left]--[l]--[l.next = right]--[right.next]
					// [left.prev]--[l]--[left]--[right]--[right.next]

					right := l.next
					left := l.prev
					l.prev = left.prev
					l.next = left
					left.prev.next = l
					left.prev = l
					left.next = right
					right.prev = left
				}
			} else if num > 0 {
				for i := 0; i < num; i++ {
					// [left.prev]--[l.prev = left]--[l]--[l.next = right]--[right.next]
					// [left.prev]--[left]--[right]--[l]--[right.next]

					right := l.next
					left := l.prev
					l.prev = right
					l.next = right.next
					right.next.prev = l
					right.next = l
					right.prev = left
					left.next = right
				}
			}

		}
	}
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 1000; i++ {
		l = l.next
	}
	v1 := l.num
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 2000; i++ {
		l = l.next
	}
	v2 := l.num
	for l.num != 0 {
		l = l.next
	}
	for i := 0; i < 3000; i++ {
		l = l.next
	}
	v3 := l.num

	return v1 + v2 + v3
}
