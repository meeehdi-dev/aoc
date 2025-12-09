package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
	Z int
}

func getDist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2) + math.Pow(float64(p1.Z-p2.Z), 2))
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d08/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	points := make([]Point, 0)
	for line := range lines {
		if len(line) == 0 {
			continue
		}

		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		p := Point{x, y, z}
		points = append(points, p)
	}

	dists := make([]float64, 0)
	distPoints := make(map[float64][]Point)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := getDist(points[i], points[j])
			dists = append(dists, dist)
			distPoints[dist] = make([]Point, 0)
			distPoints[dist] = append(distPoints[dist], points[i])
			distPoints[dist] = append(distPoints[dist], points[j])
		}
	}
	sort.Float64s(dists)
	m := make(map[Point]int)
	last := 0
	for i := 0; i < len(dists); i++ {
		if i >= len(dists) {
			break
		}
		dist := dists[i]
		current_points := distPoints[dist]
		val := -1
		max_idx := 0
		for _, point := range current_points {
			if val != -1 {
				break
			}
			for j, idx := range m {
				if j == point {
					if val == -1 {
						val = idx
					}
				}
				if idx > max_idx {
					max_idx = idx
				}
			}
		}
		if val == -1 {
			val = max_idx + 1
		}
		swap := -1
		for _, point := range current_points {
			if m[point] > 0 {
				swap = m[point]
			}
			m[point] = val
		}
		if swap != -1 {
			for j, v := range m {
				if swap == v {
					m[j] = val
				}
			}
		}
		current_count := make(map[int]int)
		for _, v := range m {
			if v > 0 {
				current_count[v]++
			}
		}
		current_counts := make([]int, 0)
		for _, c := range current_count {
			current_counts = append(current_counts, c)
		}
		if len(current_counts) == 1 && current_counts[0] == len(points) {
			last = current_points[0].X * current_points[1].X
			break
		}
	}
	count := make(map[int]int)
	for _, v := range m {
		if v > 0 {
			count[v]++
		}
	}
	counts := make([]int, 0)
	for _, c := range count {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	size := 1
	for i := len(counts) - 1; i >= len(counts)-3; i-- {
		if i < 0 {
			break
		}
		size *= counts[i]
	}

	fmt.Printf("Size: %d\n", size)
	fmt.Printf("Last: %d\n", last)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
