package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func megaConversion3000(value string) string {
	if len(value) == 1 {
		return value
	}
	switch strings.ToLower(value) {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}

// func megaSubStringer9000(value string) string {

// }

func main() {
	fmt.Println("Hello there!")

	if len(os.Args) < 2 {
		fmt.Println("You need to provide an input file")
		return
	}

	fileInput := os.Args[1]

	fmt.Println("Input file: ", fileInput)

	rawData, err := os.ReadFile(fileInput)

	check(err)

	data := string(rawData)

	formatedData := strings.Split(data, "\n")

	// regex := regexp.MustCompile("[0-9]")
	regex := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	// regex := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)+|[0-9]+")

	megaSum := 0

	for index, line := range formatedData {
		matches := regex.FindAllString(line, -1)
		if matches == nil {
			fmt.Println("No match")
			break
		}

		fmt.Println(" Row number: ", index+1, " values: ", matches[0], matches[len(matches)-1], " \t", matches)
		combinedString := megaConversion3000(matches[0]) + megaConversion3000(matches[len(matches)-1])
		calibration, err := strconv.Atoi(combinedString)

		check(err)
		megaSum += calibration
	}

	fmt.Println("Mega calibration sum: ", megaSum)
}
