package main

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		_, _ = result.WriteString("I")
	}

	return result.String()
}
