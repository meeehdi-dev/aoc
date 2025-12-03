package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.ReadFile("d03/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	output := 0
	output2 := 0

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		maxes1 := make([]int, 0)
		max_len1 := 2
		maxes2 := make([]int, 0)
		max_len2 := 12

		for i := 0; i < len(line); i++ {
			current, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}

			if len(maxes1) < max_len1 {
				maxes1 = append(maxes1, current)
			} else {
				max1 := 0
				for i := 0; i < max_len1; i++ {
					max1 += maxes1[i] * int(math.Pow(10, float64(max_len1-i-1)))
				}
				max1i := -1
				for i := 0; i < max_len1; i++ {
					maxj := 0
					for j := 0; j < max_len1; j++ {
						if i == j {
							continue
						} else if j < i {
							maxj += maxes1[j] * int(math.Pow(10, float64(max_len1-j-1)))
						} else {
							maxj += maxes1[j] * int(math.Pow(10, float64(max_len1-j)))
						}
					}
					maxj += current
					if maxj > max1 {
						max1 = maxj
						max1i = i
					}
				}
				if max1i != -1 {
					maxes1[max1i] = current
					for i := max1i; i < max_len1-1; i++ {
						maxes1[i] = maxes1[i+1]
					}
					maxes1[max_len1-1] = current
				}
			}

			if len(maxes2) < max_len2 {
				maxes2 = append(maxes2, current)
			} else {
				max_val := 0
				for i := 0; i < max_len2; i++ {
					max_val += maxes2[i] * int(math.Pow(10, float64(max_len2-i-1)))
				}
				max_i := -1
				for i := 0; i < max_len2; i++ {
					max_curr := 0
					for j := 0; j < max_len2; j++ {
						if i == j {
							continue
						} else if j < i {
							max_curr += maxes2[j] * int(math.Pow(10, float64(max_len2-j-1)))
						} else {
							max_curr += maxes2[j] * int(math.Pow(10, float64(max_len2-j)))
						}
					}
					max_curr += current
					if max_curr > max_val {
						max_val = max_curr
						max_i = i
					}
				}
				if max_i != -1 {
					maxes2[max_i] = current
					for i := max_i; i < max_len2-1; i++ {
						maxes2[i] = maxes2[i+1]
					}
					maxes2[max_len2-1] = current
				}
			}
		}

		max1 := 0
		for i := 0; i < max_len1; i++ {
			max1 += maxes1[i] * int(math.Pow(10, float64(max_len1-i-1)))
		}
		max2 := 0
		for i := 0; i < max_len2; i++ {
			max2 += maxes2[i] * int(math.Pow(10, float64(max_len2-i-1)))
		}
		output += max1
		output2 += max2
	}

	fmt.Printf("Output: %d\n", output)
	fmt.Printf("Output2: %d\n", output2)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
