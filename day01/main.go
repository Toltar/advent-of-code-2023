package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getLineValue(l string) int {
	re := regexp.MustCompile("[0-9]")
	numbersAsStrings := re.FindAllString(l, -1)
	firstNumber := numbersAsStrings[0]
	lastNumber := numbersAsStrings[len(numbersAsStrings)-1]

	fullNumber := firstNumber + lastNumber

	num, err := strconv.Atoi(fullNumber)

	if err != nil {
		panic("Error converting the fullNumber into a integer")
	}

	return num
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += getLineValue(line)
	}

	fmt.Println("Total is", total)
}
