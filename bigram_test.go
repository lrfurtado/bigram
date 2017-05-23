package bigram

import "testing"
import "reflect"

var bigramtests = []struct {
	in   string
	out  []string
	hout map[string]int
	err  error
}{
	{"", nil, nil, EmptyStringErr},
	{"This is a test", []string{"this is", "is a", "a test"}, map[string]int{"this is": 1, "is a": 1, "a test": 1}, nil},
	{"The quick brown fox and the quick blue hare", []string{"the quick", "quick brown", "brown fox", "fox and", "and the", "the quick", "quick blue", "blue hare"},
		map[string]int{"the quick": 2, "quick brown": 1, "brown fox": 1, "fox and": 1, "and the": 1, "quick blue": 1, "blue hare": 1}, nil},
	{"AAA AAA AAA AAA", []string{"aaa aaa", "aaa aaa", "aaa aaa"}, map[string]int{"aaa aaa": 3}, nil},
}

func TestParse(t *testing.T) {
	for _, tt := range bigramtests {
		o, err := Parse(tt.in)

		if err != tt.err {
			t.Errorf("Found (%v), expected (%v)", err, tt.err)
		}

		if !reflect.DeepEqual(o, tt.out) {
			t.Errorf("Found (%v), expected (%v)", o, tt.out)
		}
	}
}

func TestHistogram(t *testing.T) {
	for _, tt := range bigramtests {
		o, err := Histogram(tt.in)

		if err != tt.err {
			t.Errorf("Found (%v), expected (%v)", err, tt.err)
		}

		if !reflect.DeepEqual(o, tt.hout) {
			t.Errorf("Found (%v), expected (%v)", o, tt.out)
		}
	}
}
