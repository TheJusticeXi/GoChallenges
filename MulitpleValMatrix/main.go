package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	fmt.Println("Please enter row and column quantity:")
	fmt.Scanf("%s", &inputStr)

	row, col := GetRowAndColumn(inputStr)

	CreateMatrix(row, col)

}

func CreateMatrix(row int, col int) {

	fmt.Printf("[")
	for i := 0; i < row; i++ {
		fmt.Printf("[")
		for j := 0; j < col; j++ {
			out := i * j
			if j != col-1 {
				fmt.Printf("%d,", out)
			} else {
				fmt.Printf("%d", out)
			}

		}
		fmt.Printf("]")
	}
	fmt.Printf("]")
}

func GetRowAndColumn(input string) (int, int) {

	strs := strings.Split(input, ",")

	row := ConvertStringToInt(strs[0])

	col := ConvertStringToInt(strs[1])

	fmt.Printf("Row = %d, Column = %d\n", row, col)

	return row, col
}

func ConvertStringToInt(numStr string) int {
	marks, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println("Error during conversion")
		return -1
	}

	return marks
}
