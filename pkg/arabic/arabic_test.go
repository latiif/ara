package arabic

import "testing"

func TestGenerateTashkeelTable(t *testing.T) {
	var testcases = []struct {
		input    string
		expected []rune
	}{
		{input: "مَر", expected: []rune{'\x00', 'َ', '\x00'}},
		{input: "قُ", expected: []rune{'\x00', 'ُ'}},
		{input: "عِ", expected: []rune{'\x00', 'ِ'}},
		{input: "شّ", expected: []rune{'\x00', 'ّ'}},
		{input: "سً", expected: []rune{'\x00', 'ً'}},
		{input: "ثٌ", expected: []rune{'\x00', 'ٌ'}},
		{input: "تٍ", expected: []rune{'\x00', 'ٍ'}},
		{input: "بَ قْ", expected: []rune{'\x00', 'َ', '\x00', '\x00', 'ْ'}},
		{input: "نص بدون حركات", expected: []rune{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}},
		{input: "", expected: nil},
		{input: "", expected: []rune{}},
	}

	for _, tc := range testcases {
		result := GenerateTashkeelTable(tc.input)
		if !areEqual(result, tc.expected) {
			t.Errorf("got %q, want %q", result, tc.expected)
		}
	}
}

func areEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
