package encoder

import (
	"fmt"
	"strings"
)

// Encode encodes a single line into bracket notation using run-length encoding.
// Only runs of 2 or more consecutive identical substrings are compressed.
func Encode(input string) string {
	if input == "" {
		return ""
	}
	var result strings.Builder
	i := 0
	for i < len(input) {
		bestCount := 1
		bestLen := 1

		// Try pattern lengths from longest to shortest to prefer longer patterns
		maxPatLen := (len(input) - i) / 2
		for patLen := maxPatLen; patLen >= 1; patLen-- {
			pattern := input[i : i+patLen]
			count := 1
			for i+count*patLen+patLen <= len(input) && input[i+count*patLen:i+count*patLen+patLen] == pattern {
				count++
			}
			if count > 1 && count*patLen > bestCount*bestLen {
				bestCount = count
				bestLen = patLen
			}
		}

		if bestCount > 1 {
			pattern := input[i : i+bestLen]
			result.WriteString(fmt.Sprintf("[%d %s]", bestCount, pattern))
			i += bestCount * bestLen
		} else {
			result.WriteByte(input[i])
			i++
		}
	}
	return result.String()
}

// EncodeLines encodes multiple lines.
func EncodeLines(lines []string) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = Encode(line)
	}
	return result
}
