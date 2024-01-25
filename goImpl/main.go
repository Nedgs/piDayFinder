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
		os.Exit(2)
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
	index := findPattern(content, birthdate)
	if index != -1 {
		fmt.Printf("Le motif '%s' a été trouvé à l'index %d.\n", birthdate, index)
	} else {
		fmt.Printf("Le motif '%s' n'a pas été trouvé dans le tableau de bytes.\n", birthdate)
	}
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

func findPattern(data []byte, pattern string) int {
	patternBytes := []byte(pattern)
	patternLen := len(patternBytes)

	for i := 0; i <= len(data)-patternLen; i++ {
		match := true
		for j := 0; j < patternLen; j++ {
			if data[i+j] != patternBytes[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
