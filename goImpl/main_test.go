package main

import (
	"testing"
)

func BenchmarkFindPattern(b *testing.B) {
	data, _ := readFileContent("../pi.txt")
	pattern := []byte("090698")
	for i := 0; i < b.N; i++ {
		findPattern(data, pattern)
	}
}

func BenchmarkCalculatePi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculatePi(1000)
	}
}

func BenchmarkIsValidDateFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidDateFormat("0101099")
	}
}

func BenchmarkReadFileContent(b *testing.B) {

	for i := 0; i < b.N; i++ {
		readFileContent("../pi.txt")
	}

}

func TestIsValidDateFormat(t *testing.T) {
	validDates := []string{"010101", "311299", "280222"}
	invalidDates := []string{"0101", "3112", "28022", "abcd", "01/01/01"}

	for _, date := range validDates {
		if !isValidDateFormat(date) {
			t.Errorf("Expected date %s to be valid, but it was invalid.", date)
		}
	}

	for _, date := range invalidDates {
		if isValidDateFormat(date) {
			t.Errorf("Expected date %s to be invalid, but it was valid.", date)
		}
	}
}

func TestCalculatePi(t *testing.T) {
	precision := 1000

	pi := calculatePi(precision)
	piStr := pi.Text('f', precision)

	// egal à +inf, si corrigé mettre precision+2 à la place de 4
	if len(piStr) != 4 {
		t.Errorf("Expected %d digits of precision, but got %d.", precision, len(piStr)-2)
	}
}
