package problems

import "testing"

func TestStringHasAllUniqueCharacters(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"abcd", true},
		{"aacd", false},
		{"abbc", false},
		{"abca", false},
	}

	for _, test := range tests {
		if output := StringHasAllUniqueCharacters(test.input); output != test.expected {
			t.Errorf("FAILED(%v), expected: %v, received: %v", test.input, test.expected, output)
		}
	}
}

func TestStringReverseNullTerminated(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
	}

	for _, test := range tests {
		if output := StringReverseNullTerminated(test.input); output != test.expected {
			t.Errorf("FAILED(%v), expected: %v, received: %v", test.input, test.expected, output)
		}
	}
}
