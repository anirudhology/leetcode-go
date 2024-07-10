package strings

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(strs []string) string {
	var encoded strings.Builder
	for _, s := range strs {
		encoded.WriteString(fmt.Sprintf("%d/%s", len(s), s))
	}
	return encoded.String()
}

func Decode(s string) []string {
	var decoded []string
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] != '/' {
			j++
		}
		size, _ := strconv.Atoi(s[i:j])
		word := s[j+1 : j+1+size]
		i = j + 1 + size
		decoded = append(decoded, word)
	}
	return decoded
}
