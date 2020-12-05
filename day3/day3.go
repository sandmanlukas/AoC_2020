package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
	//"strings"
)

func main() {
	first := slope(3, 1)
	second := slope(1, 1)
	third := slope(5, 1)
	fourth := slope(7, 1)
	fifth := slope(1, 2)
	ans := first * second * third * fourth * fifth
	fmt.Println(ans)

}

func slope(right int, down int) int {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counter := -1
	scanner := bufio.NewScanner(file)
	line_len := 0
	trees := 0
	index := 0
	for scanner.Scan() {
		counter++

		if down > 1 && (counter+1)%2 == 0 {
			fmt.Println("Skipping line: ", counter+1)
			continue
		}
		value := scanner.Text()

		line_len = len(value) - 1

		if counter == 0 {
			continue
		}

		if index+right > line_len {
			index = (index+right)%30 - 1
			//fmt.Println(index)
		} else {
			index += right
		}
		char1 := string([]rune(value)[index])
		if char1 == "#" {
			trees++
			continue
		}

	}
	fmt.Println("The answer to question 1 is: ", trees)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return trees
}
