package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1, str2 string
	fmt.Println("Please enter String1 and String 2 with a space between")
	fmt.Scanf("%s %s", &str1, &str2)

	if len(str1) != len(str2) {
		fmt.Println("These are not Anagrams")
		return
	}

	for index := 0; index < len(str1); index++ {
		if strings.Contains(str2, string(str1[index])) {
			ind := strings.Index(str2, string(str1[index]))
			str2 = strRemoveAt(str2, ind)
		}
	}

	if len(str2) == 0 {
		fmt.Println("These are Anagrams")
	} else {
		fmt.Println("These are not Anagrams")
	}

	return
}

func strRemoveAt(s string, index int) string {
	length := len(s)
	return s[:index] + s[index+length:]
}
