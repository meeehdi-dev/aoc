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

	file, err := os.ReadFile("d02/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.SplitSeq(string(file), "\n")

	invalids := 0
	invalids2 := 0

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		ranges := strings.SplitSeq(line, ",")
		for r := range ranges {
			if len(r) == 0 {
				continue
			}

			ids := strings.Split(r, "-")
			start, err := strconv.Atoi(ids[0])
			if err != nil {
				panic(err)
			}
			last, err := strconv.Atoi(ids[1])
			if err != nil {
				panic(err)
			}

			for i := start; i <= last; i++ {
				str := strconv.Itoa(i)
				if str[:len(str)/2] == str[len(str)/2:] {
					invalids += i
				}
				for strlen := len(str) / 2; strlen > 0; strlen-- {
					if len(str)%strlen != 0 {
						continue
					}
					found := true
					for index := 0; index < len(str)/strlen-1; index++ {
						if str[strlen*index:strlen*(index+1)] != str[strlen*(index+1):strlen*(index+2)] {
							found = false
							break
						}
					}
					if found {
						invalids2 += i
						break
					}
				}
			}
		}
	}

	fmt.Printf("Invalids: %d\n", invalids)
	fmt.Printf("Invalids2: %d\n", invalids2)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
