package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	//"sort"
	"strings"
)

func main() {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans1 := float64(0)

	scanner := bufio.NewScanner(file)
	ids := make(map[int]int)
	for scanner.Scan() {
		//row := 0
		//col := 0
		low_row := float64(0)
		high_row := float64(127)
		low_col := float64(0)
		high_col := float64(7)
		line := strings.Trim(scanner.Text(), "\n")
		for _, rune := range line {
			char := string(rune)

			val_row := math.Round((high_row - low_row) / 2)
			val_col := math.Round((high_col - low_col) / 2)

			switch char {
			case "F":
				high_row = high_row - val_row
			case "B":
				low_row = low_row + val_row
			case "L":
				high_col = high_col - val_col
			default:
				low_col = low_col + val_col

			}

		}

		if low_row == high_row && low_col == high_col {
			ans1 = math.Max(float64(ans1), low_row*8+low_col)
			id := int(low_row*8 + low_col)
			ids[id] = 1
		}

	}
	max := float64(0)
	min := float64(ans1)
	ids_new := make(map[int]int)

	for key, _ := range ids {
		if ids[key+1] == 1 || ids[key-1] == 1 {
			max = math.Max(max, float64(key))
			min = math.Min(min, float64(key))
			ids_new[key] = 1
		}
	}

	for i := int(min); i <= int(max); i++ {
		if ids_new[i] == 0 {
			fmt.Println("Your seat: ", i)
		}
	}

	fmt.Println("The answer to question 1 is: ", ans1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
