package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var valid1 int
	var valid2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		s := strings.Split(value, " ")
		cond := strings.Split(s[0], "-")
		low, _ := strconv.Atoi(cond[0])
		high, _ := strconv.Atoi(cond[1])

		low_index := low - 1
		high_index := high - 1

		pass := s[2]
		char := strings.Split(s[1], ":")[0]

		char1 := string([]rune(pass)[low_index])
		char2 := string([]rune(pass)[high_index])
		if char1 == char && char2 != char {
			valid2++
			fmt.Println("Valid password - Question2", char, pass, low_index, high_index)
		} else if char1 != char && char2 == char {
			valid2++
			fmt.Println("Valid password - Question2", char, pass, low_index, high_index)
		}
		cnt := strings.Count(s[2], char)
		if cnt >= low && cnt <= high {
			valid1++
			//fmt.Println("Valid password ")
		}

	}
	fmt.Println("The answer to question 1 is: ", valid1)
	fmt.Println("The answer to question 2 is: ", valid2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
