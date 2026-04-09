package service

import (
	"strconv"
	"strings"

	"github.com/itsdarkhost/rbk-week1/pkg"
)

type TextConverter struct {
	Input string
}

func (tc TextConverter) Convert() string {
	raw := strings.Fields(tc.Input)

	// Решаем случай с "(low 3)" -> ["(low", "3)"]
	words := tc.mergeCommands(raw)

	var res []string
	for i := 0; i < len(words); i++ {
		w := words[i]

		if pkg.IsCommand(w) {
			cmd, n := pkg.ParseCommand(w)
			tc.apply(res, cmd, n)
			continue
		}

		w = strings.Trim(w, ",.")
		res = append(res, w)
	}

	return pkg.NormalizePunctuation(res)
}

func (tc TextConverter) apply(words []string, cmd string, n int) {
	start := len(words) - n
	if start < 0 {
		start = 0
	}

	for i := start; i < len(words); i++ {
		switch cmd {

		case "up":
			words[i] = strings.ToUpper(words[i])

		case "low":
			words[i] = strings.ToLower(words[i])

		case "cap":
			words[i] = pkg.Capitilize(words[i])

		case "hex":
			if v, err := strconv.ParseInt(words[i], 16, 64); err == nil {
				words[i] = strconv.FormatInt(v, 10)
			}

		case "bin":
			if v, err := strconv.ParseInt(words[i], 2, 64); err == nil {
				words[i] = strconv.FormatInt(v, 10)
			}
		}
	}
}

func (tc TextConverter) mergeCommands(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && !strings.HasSuffix(words[i], ")") {
			if i+1 < len(words) {
				merged := words[i] + " " + words[i+1]
				result = append(result, merged)
				i++
				continue
			}
		}

		result = append(result, words[i])
	}

	return result
}
