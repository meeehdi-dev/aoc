package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.ReadFile("d07/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	beams := make(map[int]map[int]int)
	count := 0
	for idx := range lines {
		line := lines[idx]
		if len(line) == 0 {
			continue
		}
		if beams[idx] == nil {
			beams[idx] = make(map[int]int)
		}

		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				beams[idx+1] = make(map[int]int)
				beams[idx+1][i] = 1
			} else if line[i] == '^' {
				if beams[idx-1][i] > 0 {
					count++
				}
				if i > 0 {
					beams[idx][i-1] += beams[idx-1][i]
				}
				if i < len(line)-1 {
					beams[idx][i+1] += beams[idx-1][i]
				}
			} else if line[i] == '.' && idx > 0 {
				if beams[idx-1][i] > 0 {
					beams[idx][i] += beams[idx-1][i]
				}
			}
		}
	}
	count2 := 0
	l := len(lines)-2
	for j := range beams[l] {
		if beams[l][j] > 0 {
			count2 += beams[l][j]
		}
	}

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Count2: %d\n", count2)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
