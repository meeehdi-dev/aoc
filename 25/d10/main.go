package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	len        int
	indicators []int
	buttons    [][]int
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d10/example.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	machines := make([]Machine, 0)

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		machine := Machine{
			len:        0,
			indicators: nil,
			buttons:    make([][]int, 0),
		}

		for i := 0; i < len(line); i++ {
			if line[i] == '[' {
				indicators := make([]int, 0)
				for j := i + 1; j < len(line); j++ {
					if line[j] == ']' {
						machine.len = j - i - 1
						i = j
						break
					}
					switch line[j] {
					case '#':
						indicators = append(indicators, j-i-1)
					}
				}
				machine.indicators = indicators
			}
			if line[i] == '(' {
				buttons := make([]int, 0)
				num := ""
				for j := i + 1; j < len(line); j++ {
					if line[j] == ')' {
						n, err := strconv.Atoi(num)
						if err != nil {
							panic(err)
						}
						buttons = append(buttons, n)
						num = ""
						i = j
						break
					}
					switch line[j] {
					case ',':
						n, err := strconv.Atoi(num)
						if err != nil {
							panic(err)
						}
						buttons = append(buttons, n)
						num = ""
					default:
						num += string(line[j])
					}
				}
				machine.buttons = append(machine.buttons, buttons)
			}
		}

		machines = append(machines, machine)
		break
	}

	presses := 0

	fmt.Printf("Presses: %d\n", presses)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
