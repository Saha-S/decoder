package decoder

import (
	"fmt"
	"strconv"
	"strings"
)

// Decode decodes a single line of encoded art, expanding bracket notation [N str].
func Decode(input string) (string, error) {
	var result strings.Builder
	i := 0
	for i < len(input) {
		if input[i] == ']' {
			return "", fmt.Errorf("Error")
		}
		if input[i] == '[' {
			j := strings.Index(input[i:], "]")
			if j == -1 {
				return "", fmt.Errorf("Error")
			}
			j = i + j
			inside := input[i+1 : j]
			spaceIdx := strings.Index(inside, " ")
			if spaceIdx == -1 {
				return "", fmt.Errorf("Error")
			}
			countStr := inside[:spaceIdx]
			str := inside[spaceIdx+1:]
			if str == "" {
				return "", fmt.Errorf("Error")
			}
			count, err := strconv.Atoi(countStr)
			if err != nil || count <= 0 {
				return "", fmt.Errorf("Error")
			}
			result.WriteString(strings.Repeat(str, count))
			i = j + 1
		} else {
			result.WriteByte(input[i])
			i++
		}
	}
	return result.String(), nil
}

// DecodeLines decodes multiple lines. Returns an error if any line fails.
func DecodeLines(lines []string) ([]string, error) {
	result := make([]string, len(lines))
	for i, line := range lines {
		decoded, err := Decode(line)
		if err != nil {
			return nil, err
		}
		result[i] = decoded
	}
	return result, nil
}
