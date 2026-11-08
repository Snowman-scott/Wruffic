package main

import (
	"bufio"
	"fmt"
	"io"
	// "os"
	"strings"
)

func ScannerHelper(prompt string, scanner *bufio.Scanner) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return strings.TrimSpace(scanner.Text()), nil
}
