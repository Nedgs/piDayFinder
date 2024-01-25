package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"regexp"
)

func main() {
	content, err := readFileContent("../pi.txt")

	if err != nil {
		fmt.Printf("Error during readFileContent execution: %s\n", err)
		os.Exit(2)
	}

	precision := 1000000
	pi := new(big.Float).SetPrec(uint(precision)).Quo(big.NewFloat(22), piLeibniz(precision))

	// Convertir le résultat en format texte avec la précision spécifiée
	piStr := pi.Text('f', precision)

	fmt.Println(len(content))

	var birthdate string
	var isValidDate = false
	for !isValidDate {
		fmt.Print("Enter your date of birth in the format DDMM: ")
		fmt.Scanln(&birthdate)
		isValidDate = isValidDateFormat(birthdate)

		if !isValidDate {
			fmt.Println("Invalid date format. Please use the format DDMMYY.")
		}
	}

	index := findPattern(content, []byte(birthdate))
	if index != -1 {
		fmt.Printf("Pattern '%s' find at index %d.\n", birthdate, index)
	} else {
		fmt.Printf("Pattern '%s' not found.\n", birthdate)
	}
	index = findPattern([]byte(piStr), []byte(birthdate))
	if index != -1 {
		fmt.Printf("Pattern '%s' find at index %d.\n", birthdate, index)
	} else {
		fmt.Printf("Pattern '%s' not found.\n", birthdate)
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
	regex := regexp.MustCompile(`^\d{6}$`)
	return regex.MatchString(date)
}

func findPattern(data []byte, pattern []byte) int {
	return bytes.Index(data, pattern)
}

func piLeibniz(precision int) *big.Float {
	pi := new(big.Float).SetPrec(uint(precision))
	pi.SetInt64(0)
	sign := 1.0
	denom := new(big.Float).SetInt64(1)

	for i := 0; i < precision; i++ {
		term := new(big.Float).SetFloat64(sign).Quo(big.NewFloat(1), denom)
		pi.Add(pi, term)
		sign *= -1
		denom.Add(denom, big.NewFloat(2))
	}

	return pi.Mul(pi, big.NewFloat(4))
}
