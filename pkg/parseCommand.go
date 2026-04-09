package pkg

import (
	"strconv"
	"strings"
)

// MARK: ParseCommand
func ParseCommand(s string) (string, int) {
	s = strings.Trim(s, "()")

	if strings.Contains(s, ",") {
		parts := strings.Split(s, ",")
		cmd := strings.TrimSpace(parts[0])
		n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		return cmd, n
	}

	if strings.Contains(s, " ") {
		parts := strings.Fields(s)
		if len(parts) == 2 {
			n, _ := strconv.Atoi(parts[1])
			return parts[0], n
		}
	}

	return s, 1
}

// MARK: IsCommand
func IsCommand(s string) bool {
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

// MARK: NormalizePunctuation
func NormalizePunctuation(words []string) string {
	text := strings.Join(words, " ")

	replacer := strings.NewReplacer(
		" ,", ",",
		" .", ".",
		" !", "!",
		" ?", "?",
		" ;", ";",
		" :", ":",
	)

	text = replacer.Replace(text)

	text = NormalizeQuotes(text)

	text = strings.Join(strings.Fields(text), " ")
	return text
}

// MARK: NormalizeQuotes
func NormalizeQuotes(s string) string {
	s = strings.ReplaceAll(s, "' ", "'")
	s = strings.ReplaceAll(s, " '", "'")

	s = strings.ReplaceAll(s, "\" ", "\"")
	s = strings.ReplaceAll(s, " \"", "\"")

	return s
}

// MARK: Capitilize
func Capitilize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
