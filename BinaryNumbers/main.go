package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println("Please enter set of binary numbers:")
	var input string
	fmt.Scanf("%s", &input)

	numbers := strings.Split(input, ",")

	var out []string

	for _, num := range numbers {
		//fmt.Printf("input = %s\n", num)
		DecNum := ConvertBinaryNumStringToNumber(num)
		//fmt.Printf("%d\n", intNum)

		if math.Mod(DecNum, 5.0) == 0 {
			out = append(out, num)
		}
	}

	fmt.Printf("These can be divided by 5 = %s\n", strings.Join(out, ","))

	return
}

func ConvertBinaryNumStringToNumber(input string) float64 {
	result := 0.0
	count := 0.0

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '1' {
			result += math.Pow(2, count)
		}
		count++
	}

	return result
}
