package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
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
func getID(x, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}

type pos struct {
	X int
	Y int
}

func part1() int {
	inp := c.ReadInputFile()
	var start pos
	var end pos
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			ch := inp[y][x]
			if ch == 'S' {
				inp[y] = strings.Replace(inp[y], "S", "a", 1)
				start = pos{X: x, Y: y}
			} else if ch == 'E' {
				inp[y] = strings.Replace(inp[y], "E", "z", 1)
				end = pos{X: x, Y: y}
			}
		}
	}
	graph := buildGraph(inp)

	s, _ := graph.GetMapping(getID(start.X, start.Y))
	e, _ := graph.GetMapping(getID(end.X, end.Y))
	d, err := graph.Shortest(s, e)
	if err != nil {
		c.Print(err.Error())

	}

	return int(d.Distance)
}

func part2() int {
	inp := c.ReadInputFile()
	var starts []pos
	var end pos
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			ch := inp[y][x]
			if ch == 'S' {
				ch = 'a'
				inp[y] = strings.Replace(inp[y], "S", "a", 1)
			} else if ch == 'E' {
				inp[y] = strings.Replace(inp[y], "E", "z", 1)
				end = pos{X: x, Y: y}
			}
			if ch == 'a' {
				starts = append(starts, pos{X: x, Y: y})
			}
		}
	}
	graph := buildGraph(inp)

	e, _ := graph.GetMapping(getID(end.X, end.Y))
	best := int64(-1)
	for _, start := range starts {
		s, _ := graph.GetMapping(getID(start.X, start.Y))
		d, err := graph.Shortest(s, e)

		if err != nil {
			continue
		}
		if best == -1 || d.Distance < best {
			best = d.Distance
		}
	}
	return int(best)
}
func buildGraph(inp []string) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			ch := inp[y][x]
			if _, err := graph.GetMapping(getID(x, y)); err != nil {
				graph.AddVertex(graph.AddMappedVertex(getID(x, y)))
			}

			if y > 0 && inp[y-1][x] <= ch+1 {
				if _, err := graph.GetMapping(getID(x, y-1)); err != nil {
					graph.AddMappedVertex(getID(x, y-1))
				}
				graph.AddMappedArc(getID(x, y), getID(x, y-1), 1)
			}
			if y < len(inp)-1 && inp[y+1][x] <= ch+1 {
				if _, err := graph.GetMapping(getID(x, y+1)); err != nil {
					graph.AddMappedVertex(getID(x, y+1))
				}
				graph.AddMappedArc(getID(x, y), getID(x, y+1), 1)
			}
			if x > 0 && inp[y][x-1] <= ch+1 {
				if _, err := graph.GetMapping(getID(x-1, y)); err != nil {
					graph.AddMappedVertex(getID(x-1, y))
				}
				graph.AddMappedArc(getID(x, y), getID(x-1, y), 1)
			}
			if x < len(inp[y])-1 && inp[y][x+1] <= ch+1 {
				if _, err := graph.GetMapping(getID(x+1, y)); err != nil {
					graph.AddMappedVertex(getID(x+1, y))
				}
				graph.AddMappedArc(getID(x, y), getID(x+1, y), 1)
			}
		}
	}
	return graph

}
