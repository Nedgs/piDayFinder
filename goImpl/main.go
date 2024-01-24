package main

import (
	"fmt"
	"os"
)

func main() {

	content, _ := readFileContent("../pi.txt")
	fmt.Println(len(content))
}

func readFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
