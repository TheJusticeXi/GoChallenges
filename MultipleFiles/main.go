package main

import (
	"fmt"
)

func main() {

	input := ReadInput()

	output := ConvertUpperCase(input)

	fmt.Printf("Output = %s", output)

	return
}
