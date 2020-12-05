package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passwords()
}

func passwords() int {
	path := "./input.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	req := map[string]int{
		"byr": 1,
		"iyr": 1,
		"eyr": 1,
		"hgt": 1,
		"hcl": 1,
		"ecl": 1,
		"pid": 1}
	valid1 := 0
	//valid2 := 0
	pass := ""
	line_count := 0
	var yr_req int
	for scanner.Scan() {
		line_count++
		line := scanner.Text()
		line = strings.Trim(line, "\n")

		if line != "" {
			pass += " " + line
		}
		if line == "" || line_count == 1056 {

			entries := strings.Split(strings.TrimSpace(pass), " ")
			s := make(map[string]int)
			//fmt.Println(entries)
			for i, _ := range entries {
				//fmt.Println(entries[i])
				field := strings.Split(entries[i], ":")[0]
				validation := strings.Split(entries[i], ":")[1]

				if field == "iyr" {
					yr_req, _ = strconv.Atoi(validation)
					if yr_req > 2020 || yr_req < 2010 {
						fmt.Println("Issue year is wrong.", validation)
						//continue
						break
					}
				} else if field == "byr" {
					yr_req, _ = strconv.Atoi(validation)
					if yr_req > 2002 || yr_req < 1920 {
						fmt.Println("Birth year is wrong.", validation)
						//continue
						break
					}
				} else if field == "eyr" {
					yr_req, _ = strconv.Atoi(validation)
					if yr_req > 2030 || yr_req < 2020 {
						fmt.Println("Exp year is wrong", validation)
						//continue
						break
					}
				} else if field == "hgt" {
					if strings.Contains(validation, "cm") {
						length, _ := strconv.Atoi(strings.Split(validation, "cm")[0])
						if length < 150 || length > 193 {
							fmt.Println("Wrong length (cm)", validation)
							//continue
							break

						}
					} else {
						length, _ := strconv.Atoi(strings.Split(validation, "in")[0])
						if length < 59 || length > 76 {
							fmt.Println("Wrong length (in)", validation)
							//continue
							break
						}
					}
				} else if field == "hcl" {
					match, _ := regexp.MatchString("^#(?:[0-9a-fA-F]{3}){1,2}$", validation)
					if !match {
						fmt.Println("Incorrect hair color", validation)
						//continue
						break
					}

				} else if field == "ecl" {
					colour := map[string]int{"amb": 1, "blu": 1, "brn": 1, "gry": 1,
						"grn": 1, "hzl": 1, "oth": 1}
					if 1 != colour[validation] {
						fmt.Println("Wrong eye color", validation)
						//continue
						break
					}
				} else if field == "pid" {
					if _, err := strconv.Atoi(validation); err != nil || len(validation) != 9 {
						fmt.Println("Wrong pid", validation)
						//continue
						break
					}
					/*
						match, _ := regexp.MatchString("\\d{9}", conv)
						if !match {
							fmt.Println("Incorrect PID", conv)
							continue
						}
					*/
				}
				if field != "cid" {
					s[field] = 1
				}

			}
			res1 := reflect.DeepEqual(s, req)

			if res1 {
				fmt.Println("Found valid pass: ", pass)
				valid1++
			} else {
				//fmt.Println("Invalid pass: ", pass, "Counter: ", len(entries))
			}
			pass = ""
		}
	}
	fmt.Println("The answer to question 1 is: ", valid1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}
