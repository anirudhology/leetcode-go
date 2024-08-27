package math

import "strconv"

func MultiplyStrings(num1 string, num2 string) string {
	// Special case
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// Lengths of both strings
	m, n := len(num1), len(num2)
	// Array to store the product
	products := make([]int, m+n-1)
	// Process the strings
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			products[i+j] += (int(num1[i]) - int('0')) * (int(num2[j]) - int('0'))
		}
	}
	for i := len(products) - 1; i > 0; i-- {
		products[i-1] = products[i-1] + int(products[i]/10)
		products[i] = products[i] % 10
	}
	// String to store final output
	output := ""
	for _, i := range products {
		output += strconv.Itoa(i)
	}
	return output
}
