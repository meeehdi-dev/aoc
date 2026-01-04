package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

func getArea(p1, p2 Point) int {
	return int((math.Abs(float64(p1.X-p2.X)) + 1) * (math.Abs(float64(p1.Y-p2.Y)) + 1))
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d09/example.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	points := make([]Point, 0)

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}

	max_area := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := getArea(points[i], points[j])
			if area > max_area {
				max_area = area
			}
		}
	}

	fmt.Printf("Area: %d\n", max_area)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
