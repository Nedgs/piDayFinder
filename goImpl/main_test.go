package main

import (
	"testing"
)

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

	if len(piStr) == precision+2 {
		t.Errorf("Expected %d digits of precision, but got %d.", precision, len(piStr)-2)
	}
}
