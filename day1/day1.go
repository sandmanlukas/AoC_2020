package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		num, _ := strconv.Atoi(value)
		ints = append(ints, num)
	}
	sort.Ints(ints)
	var left, right int
	for first := 0; first < len(ints)-2; first++ {
		left = first + 1
		right = len(ints) - 1
		for right > left {
			if ints[left]+ints[right]+ints[first] == 2020 {
				value := ints[left] * ints[right] * ints[first]
				fmt.Printf("Answer is: %d\n", value)
			}
			if ints[left]+ints[right]+ints[first] > 2020 {
				right--
			} else {
				left++
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
