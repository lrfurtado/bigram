package bigram

import "errors"
import "strings"

var EmptyStringErr = errors.New("empty string")

func Parse(s string) ([]string, error) {
	if s == "" {
		return nil, EmptyStringErr
	}
	words := strings.Split(s, " ")

	b := []string{}
	for i := 0; i < len(words)-1; i++ {
		b = append(b, strings.ToLower(words[i]+" "+words[i+1]))
	}

	return b, nil
}

func Histogram(s string) (map[string]int, error) {
	bigrams, err := Parse(s)
	if err != nil {
		return nil, err
	}
	h := make(map[string]int)
	for _, b := range bigrams {
		h[b]++
	}
	return h, nil
}
