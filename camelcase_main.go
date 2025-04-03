// camelcase_main.go

package main

import (
	"dsa/modules/camelcase"
	"fmt"
)

func main() {
	examples := []string{
		"hello world example",
		" spaced  input ",
		"",
		"   ",
		"42 test case",
		"hello @world",
	}

	for _, ex := range examples {
		result, err := camelcase.Convert(ex)
		if err != nil {
			fmt.Printf("Input: %-20q ➜ Error: %v\n", ex, err)
		} else {
			fmt.Printf("Input: %-20q ➜ CamelCase: %s\n", ex, result)
		}
	}
}
