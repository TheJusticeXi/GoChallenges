package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Please enter series of number")
	var input string = ""
	fmt.Scanf("%s", &input)
	slc := ConvertStringToSlice(input)
	//var outSlc []float64;

	const C int = 50
	const H int = 30
	outStr := ""
	for _, element := range slc {
		out := math.Round(math.Sqrt(float64(2 * C * element / H)))
		outStr += strconv.Itoa(int(out))
		outStr += ","
	}

	outStr = outStr[:len(outStr)-1]

	fmt.Println(outStr)

	return

}

func ConvertStringToSlice(input string) []int {
	var intSlices []int
	var temp string = ""

	if input[len(input)-1] != ',' {
		input += ","
	}

	for i := 0; i < len(input); i++ {
		if input[i] == ',' {
			num, err := strconv.Atoi(temp)
			if err != nil {
				// ... handle error
				panic(err)
			} else {
				intSlices = append(intSlices, num)
				temp = ""
			}

		} else {
			temp += string(input[i])
		}
	}

	return intSlices
}
