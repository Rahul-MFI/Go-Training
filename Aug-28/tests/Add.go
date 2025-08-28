package tests

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	delimiters := []string{",", "\n"}
	numbers := input

	if strings.HasPrefix(input, "//") {
		parts := strings.SplitN(input, "\n", 2)
		header, body := parts[0], parts[1]
		numbers = body

		if strings.Contains(header, "[") {
			re := regexp.MustCompile(`\[(.+?)\]`)
			matches := re.FindAllStringSubmatch(header, -1)
			for _, m := range matches {
				delimiters = append(delimiters, m[1])
			}
		} else {
			delimiters = append(delimiters, string(header[2]))
		}
	}

	replacer := strings.NewReplacer(makeReplacements(delimiters)...)
	normalized := replacer.Replace(numbers)
	parts := strings.Split(normalized, ",")

	sum := 0
	var negatives []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return 0, err
		}
		if n < 0 {
			negatives = append(negatives, p)
			continue
		}
		if n <= 1000 {
			sum += n
		}
	}
	if len(negatives) > 0 {
		return 0, errors.New("negatives not allowed: " + strings.Join(negatives, ","))
	}
	return sum, nil
}

func makeReplacements(delims []string) []string {
	var reps []string
	for _, d := range delims {
		reps = append(reps, d, ",")
	}
	return reps
}
