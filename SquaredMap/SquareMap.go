package main

import (
	"fmt"
)

func main() {
	var MapQnt int = 0
	fmt.Println("Enter Squared Map Number:")
	fmt.Scanf("%d", &MapQnt)

	//m := make(map[int]int)

	if MapQnt <= 0 {
		fmt.Println("Cannot Create the Squared Map")
	} else {
		fmt.Printf("[")

		for index := 1; index <= MapQnt; index++ {
			fmt.Printf(" %d:%d ", index, index*index)
		}

		fmt.Printf("]")

	}

	return

}
