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

	file, err := os.ReadFile("d06/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	seps := make([]int, 0)
	for idx := range lines {
		line := lines[idx]
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(line); i++ {
			if line[i] == ' ' {
				seps = append(seps, i)
			}
		}
		break
	}
	for idx := range lines {
		line := lines[idx]
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(seps); i++ {
			if line[seps[i]] != ' ' {
				seps = append(seps[:i], seps[i+1:]...)
				i--
			}
		}
	}

	ops := make([]string, 0)
	for i := 0; i < len(lines[len(lines)-2]); i++ {
		op := lines[len(lines)-2][i]
		if op != ' ' {
			ops = append(ops, string(op))
		}
	}

	m := make([][]int, len(seps)+1)
	m2 := make([][]string, len(seps)+1)
	for i := range m {
		m[i] = make([]int, 0)
		m2[i] = make([]string, 0)
	}
	total := 0
	total2 := 0

	for idx := range lines {
		line := lines[idx]
		if len(line) == 0 {
			continue
		}

		sep := 0
		for i := 0; i <= len(seps); i++ {
			val := ""
			if i < len(seps) {
				val = line[sep:seps[i]]
				sep = seps[i] + 1
			} else {
				val = line[sep:]
			}
			t := strings.TrimSpace(val)
			if t == "*" || t == "+" {
				break
			}
			v, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			m[i] = append(m[i], v)
			m2[i] = append(m2[i], val)
		}
	}

	for i := 0; i < len(ops); i++ {
		v := 0
		if ops[i] == "*" {
			v = 1
		}
		v2 := 0
		if ops[i] == "*" {
			v2 = 1
		}
		for j := 0; j < len(m[i]); j++ {
			switch ops[i] {
			case "+":
				v += m[i][j]
			case "*":
				v *= m[i][j]
			}
		}
		val := make([]string, len(m2[i]))
		for j := 0; j < len(m2[i]); j++ {
			for k := 0; k < len(m2[i][j]); k++ {
				if m2[i][j][k] == ' ' {
					continue
				}
				val[k] += string(m2[i][j][k])
			}
		}
		for j := range val {
			if val[j] == "" {
				continue
			}
			n, err := strconv.Atoi(val[j])
			if err != nil {
				panic(err)
			}
			switch ops[i] {
			case "+":
				v2 += n
			case "*":
				v2 *= n
			}
		}
		total += v
		total2 += v2
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Total2: %d\n", total2)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
