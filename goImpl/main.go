package main

import (
	"bytes"
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

	index := findPattern(content, []byte(birthdate))
	if index != -1 {
		fmt.Printf("Le motif '%s' a été trouvé à l'index %d.\n", birthdate, index)
	} else {
		fmt.Printf("Le motif '%s' n'a pas été trouvé dans le tableau de bytes.\n", birthdate)
	}
}

func readFileContent(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, errors.New("empty content")
	}
	return content, nil
}

func isValidDateFormat(date string) bool {
	regex := regexp.MustCompile(`^\d{4}$`)
	return regex.MatchString(date)
}

func findPattern(data []byte, pattern []byte) int {
	return bytes.Index(data, pattern)
}
