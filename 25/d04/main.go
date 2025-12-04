package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func count_rolls(map2d [][]rune, i int, j int) int {
	count := 0
	if i > 0 && j > 0 && map2d[i-1][j-1] == '@' {
		count += 1
	}
	if i > 0 && map2d[i-1][j] == '@' {
		count += 1
	}
	if i > 0 && j < len(map2d[i-1])-1 && map2d[i-1][j+1] == '@' {
		count += 1
	}
	if j < len(map2d[i])-1 && map2d[i][j+1] == '@' {
		count += 1
	}
	if i < len(map2d)-1 && j < len(map2d[i])-1 && map2d[i+1][j+1] == '@' {
		count += 1
	}
	if i < len(map2d)-1 && map2d[i+1][j] == '@' {
		count += 1
	}
	if i < len(map2d)-1 && j > 0 && map2d[i+1][j-1] == '@' {
		count += 1
	}
	if j > 0 && map2d[i][j-1] == '@' {
		count += 1
	}
	return count
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d04/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	map2d := make([][]rune, 0)
	count := 0
	count_all := 0

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		map2d = append(map2d, make([]rune, 0))
		for _, char := range line {
			map2d[i] = append(map2d[i], char)
		}
	}

	changed := true
	first := true
	for changed {
		changed = false
		newmap2d := make([][]rune, len(map2d))
		for i := 0; i < len(map2d); i++ {
			newmap2d[i] = make([]rune, len(map2d[i]))
			copy(newmap2d[i], map2d[i])
		}
		for i := 0; i < len(map2d); i++ {
			for j := 0; j < len(map2d[i]); j++ {
				if string(map2d[i][j]) == "@" && count_rolls(map2d, i, j) < 4 {
					newmap2d[i][j] = 'x'
					if first {
						count++
					}
					count_all++
					changed = true
				}
			}
		}
		first = false
		map2d = newmap2d
	}

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Count all: %d\n", count_all)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
