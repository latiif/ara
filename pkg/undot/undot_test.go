package undot

import "testing"

var testcases = []struct {
	input    string
	expected string
}{
	{input: "أإآا", expected: "اااا"},
	{input: "بتث", expected: "ٮٮٮ"},
	{input: "ججج", expected: "ححح"},
	{input: "رز", expected: "رر"},
	{input: "شس", expected: "سس"},
	{input: "ضص", expected: "صص"},
	{input: "طظ", expected: "طط"},
	{input: "عغ", expected: "عع"},
	{input: "ف", expected: "ڡ"},
	{input: "ق", expected: "ٯ"},
	{input: "ن", expected: "ں"},
	{input: "ي", expected: "ى"},
	{input: "ة", expected: "ه"},
	{input: "ك", expected: "ك"},
	{input: "test", expected: "test"},
	{input: "", expected: ""},
}

func TestUndot(t *testing.T) {
	for _, tc := range testcases {
		result := Undot(tc.input)
		if result != tc.expected {
			t.Errorf("got %q, want %q", result, tc.expected)
		}
	}
}
