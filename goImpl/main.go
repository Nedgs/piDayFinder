package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"time"
)

func main() {

	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	content, err := readFileContent("../pi.txt")

	if err != nil {
		fmt.Printf("Error during readFileContent execution: %s\n", err)
		os.Exit(2)
	}

	precision := 1000000
	pi := calculatePi(precision)
	piStr := pi.Text('f', precision)

	var birthdate string
	var isValidDate = false
	for !isValidDate {
		fmt.Print("Enter your date of birth in the format DDMMYY: ")
		fmt.Scanln(&birthdate)
		isValidDate = isValidDateFormat(birthdate)

		if !isValidDate {
			fmt.Println("Invalid date format. Please use the format DDMMYY.")
		}
	}

	go findPattern(content, []byte(birthdate))
	go findPattern([]byte(piStr), []byte(birthdate))

	time.Sleep(60 * time.Minute)

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

func findPattern(data []byte, pattern []byte) {
	index := bytes.Index(data, pattern)
	if index != -1 {
		fmt.Printf("Pattern '%s' find at index %d.\n", pattern, index)
	} else {
		fmt.Printf("Pattern '%s' not found.\n", pattern)
	}
}

func calculatePi(precision int) *big.Float {
	a0 := big.NewFloat(1)
	b0 := new(big.Float).SetFloat64(1 / math.Sqrt2)
	t0 := new(big.Float).SetFloat64(0.25)
	p0 := big.NewFloat(1)

	for i := 0; i < precision; i++ {
		an := new(big.Float).Add(a0, b0)
		an.Quo(an, big.NewFloat(2))
		bn := new(big.Float).Mul(a0, b0)
		bn.Sqrt(bn)
		tn := new(big.Float).Sub(a0, an)
		tn.Mul(tn, tn)
		tn.Mul(p0, tn)
		pn := new(big.Float).Add(p0, p0)
		a0, b0, t0, p0 = an, bn, tn, pn
	}

	pi := new(big.Float).Add(a0, b0)
	pi.Mul(pi, pi)
	pi.Quo(pi, t0)

	return pi
}
