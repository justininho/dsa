package medium

import (
	"strconv"
	"strings"
)

type Codec struct {
}

// Encode encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		builder.WriteString(strconv.FormatInt(int64(len(s)), 10))
		builder.WriteString("/:")
		builder.WriteString(s)
	}

	return builder.String()
}

// Decode decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	i := 0
	delimiter := "/:"
	var decoded []string

	for i < len(strs) {
		di := strings.Index(strs[i:], delimiter)
		if stringLength, err := strconv.Atoi(strs[i : i+di]); err == nil {
			start := i + di + len(delimiter)
			end := start + stringLength
			word := strs[start:end]
			decoded = append(decoded, word)
			i = end
		}
	}

	return decoded
}
