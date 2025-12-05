package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.ReadFile("d05/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	range_input := true
	ranges := make([][]int, 0)
	count := 0
	fresh := 0

	for line := range lines {
		if len(line) == 0 {
			range_input = false
			continue
		}

		if range_input {
			vals := strings.Split(line, "-")
			r := make([]int, 2)
			r[0], _ = strconv.Atoi(vals[0])
			r[1], _ = strconv.Atoi(vals[1])
			changed := true
			skip := false
			for changed {
				changed = false
				retry_idx := -1
				for idx, ra := range ranges {
					if r[0] >= ra[0] && r[0] <= ra[1] && r[1] > ra[1] {
						ra[1] = r[1]
						changed = true
						retry_idx = idx
						break
					}
					if r[1] >= ra[0] && r[1] <= ra[1] && r[0] < ra[0] {
						ra[0] = r[0]
						changed = true
						retry_idx = idx
						break
					}
					if r[0] >= ra[0] && r[1] <= ra[1] {
						// remove
						skip = true
						break
					}
					if r[0] <= ra[0] && r[1] >= ra[1] {
						ra[0] = r[0]
						ra[1] = r[1]
						changed = true
						retry_idx = idx
						break
					}
				}
				if retry_idx != -1 {
					r[0], r[1] = ranges[retry_idx][0], ranges[retry_idx][1]
					ranges = append(ranges[:retry_idx], ranges[retry_idx+1:]...)
				}
			}
			if !skip {
				ranges = append(ranges, r)
			}
		} else {
			id, _ := strconv.Atoi(line)
			for _, r := range ranges {
				if id >= r[0] && id <= r[1] {
					count++
					break
				}
			}
		}
	}

	for _, r := range ranges {
		fresh += r[1] - r[0] + 1
	}

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Fresh: %d\n", fresh)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
