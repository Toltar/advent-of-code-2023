package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NumAsString string

const (
	One   NumAsString = "one"
	Two   NumAsString = "two"
	Three NumAsString = "three"
	Four  NumAsString = "four"
	Five  NumAsString = "five"
	Six   NumAsString = "six"
	Seven NumAsString = "seven"
	Eight NumAsString = "eight"
	Nine  NumAsString = "nine"
)

func normalizeToNumString(stringNumber string) string {
	if stringNumber == string(One) {
		return "1"
	} else if stringNumber == string(Two) {
		return "2"
	} else if stringNumber == string(Three) {
		return "3"
	} else if stringNumber == string(Four) {
		return "4"
	} else if stringNumber == string(Five) {
		return "5"
	} else if stringNumber == string(Six) {
		return "6"
	} else if stringNumber == string(Seven) {
		return "7"
	} else if stringNumber == string(Eight) {
		return "8"
	} else if stringNumber == string(Nine) {
		return "9"
	}

	return stringNumber
}

func solutionForDay1Part1(l string) int {
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

func findFirstNumber(l string) string {
	result := ""
	// Setting to -1 if we dont have anything that resembles anything in the string that we want so we can ignore this
	for i := 0; i < len(l); i++ {
		// Getting the sub string walking it forward
		s := l[i:]

		if strings.HasPrefix(s, "one") || strings.HasPrefix(s, "1") {
			result = "1"
			break
		}

		if strings.HasPrefix(s, "two") || strings.HasPrefix(s, "2") {
			result = "2"
			break
		}

		if strings.HasPrefix(s, "three") || strings.HasPrefix(s, "3") {
			result = "3"
			break
		}

		if strings.HasPrefix(s, "four") || strings.HasPrefix(s, "4") {
			result = "4"
			break
		}

		if strings.HasPrefix(s, "five") || strings.HasPrefix(s, "5") {
			result = "5"
			break
		}

		if strings.HasPrefix(s, "six") || strings.HasPrefix(s, "6") {
			result = "6"
			break
		}

		if strings.HasPrefix(s, "seven") || strings.HasPrefix(s, "7") {
			result = "7"
			break
		}

		if strings.HasPrefix(s, "eight") || strings.HasPrefix(s, "8") {
			result = "8"
			break
		}

		if strings.HasPrefix(s, "nine") || strings.HasPrefix(s, "9") {
			result = "9"
			break
		}
	}

	fmt.Println("First number is", result)

	return result
}

func findLastNumber(l string) string {
	result := ""

	for i := len(l); i >= 0; i-- {
		// Getting the sub string walking it forward
		s := l[:i]

		fmt.Println("Substring is", s)

		if strings.HasSuffix(s, "one") || strings.HasSuffix(s, "1") {
			result = "1"
			break
		}

		if strings.HasSuffix(s, "two") || strings.HasSuffix(s, "2") {
			result = "2"
			break
		}

		if strings.HasSuffix(s, "three") || strings.HasSuffix(s, "3") {
			result = "3"
			break
		}

		if strings.HasSuffix(s, "four") || strings.HasSuffix(s, "4") {
			result = "4"
			break
		}

		if strings.HasSuffix(s, "five") || strings.HasSuffix(s, "5") {
			result = "5"
			break
		}

		if strings.HasSuffix(s, "six") || strings.HasSuffix(s, "6") {
			result = "6"
			break
		}

		if strings.HasSuffix(s, "seven") || strings.HasSuffix(s, "7") {
			result = "7"
			break
		}

		if strings.HasSuffix(s, "eight") || strings.HasSuffix(s, "8") {
			result = "8"
			break
		}

		if strings.HasSuffix(s, "nine") || strings.HasSuffix(s, "9") {
			result = "9"
			break
		}
	}

	fmt.Println("Last number is", result)
	return result
}

func solutionForPart2(l string) int {
	// re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	firstNumber := findFirstNumber(l)
	lastNumber := findLastNumber(l)
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

	totalPart1 := 0
	totalPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalPart1 += solutionForDay1Part1(line)
		totalPart2 += solutionForPart2(line)
	}

	fmt.Println("Total for part 1", totalPart1)
	fmt.Println("Total for part 2", totalPart2)
}
