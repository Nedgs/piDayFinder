package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func main() {

	content, err := readFileContent("../pi.txt")
	if err != nil {
		fmt.Printf("Error during readFileContent execution: %s\n", err)
	}
	fmt.Println(len(content))
	var birthdate string
	var isValidDate = false
	for !isValidDate {
		fmt.Print("Enter your date of birth in the format DDMM: ")
		fmt.Scanln(&birthdate)
		isValidDate = isValidDateFormat(birthdate)

		if !isValidDate {
			fmt.Println("Invalid date format. Please use the format DDMM.")
		}
	}

	fmt.Println(birthdate)
}

func readFileContent(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return []byte(""), errors.New("error readFile method")
	}
	if len(content) == 0 {
		return []byte(""), errors.New("empty content")
	}
	return content, nil
}

func isValidDateFormat(date string) bool {
	regex := regexp.MustCompile(`^\d{4}$`)
	return regex.MatchString(date)
}
