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

	file, err := os.ReadFile("d01/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	dial := 50
	password := 0
	password2 := 0

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		zero := dial == 0
		zeroes := 0

		switch line[0] {
		case 'L':
			val, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			dial -= val % 100
			zeroes = val / 100
		case 'R':
			val, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			dial += val % 100
			zeroes = val / 100
		}

		if dial < 0 {
			dial += 100
			if !zero {
				zeroes++
			}
		} else if dial > 100 {
			dial -= 100
			if !zero {
				zeroes++
			}
		}
		dial %= 100
		if dial == 0 {
			password += 1
		}
		password2 += zeroes

		// println(prev, line, dial, zeroes)
	}

	fmt.Printf("Password: %d\n", password)
	fmt.Printf("Password2: %d\n", password+password2)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
