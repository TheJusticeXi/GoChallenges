package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	listSpecialCharacter = []string{"|", "#", "'", "\"", "\\", "%", "?", "\n", "<", "Ø", "ð", ">", "ï", "û", "!", "@", " "}
)

func main() {

	fmt.Println("Please enter input string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	letterCount := 0
	numericCount := 0

	for _, ch := range input {
		if IsSpecialCharacter(string(ch)) {
			continue
		} else if IsNumeric(string(ch)) {
			numericCount++
		} else {
			letterCount++
		}
	}

	fmt.Printf("LETTERS = %d\n", letterCount)
	fmt.Printf("DIGITS = %d\n", numericCount)

}

func IsNumeric(ch string) bool {
	result := false
	_, err := strconv.Atoi(ch)
	if err == nil {
		result = true
	}

	return result
}

func IsSpecialCharacter(ch string) bool {
	result := false
	for _, spCh := range listSpecialCharacter {
		if ch == spCh {
			result = true
			break
		}
	}

	return result
}
