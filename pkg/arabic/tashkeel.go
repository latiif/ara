package arabic

import (
	"reflect"
)

// GenerateTashkeelTable extracts information about Tashkeels
func GenerateTashkeelTable(s string) []rune {

	sRune := []rune(s)

	r := make([]rune, len(sRune), len(sRune))

	for index, value := range sRune {
		if tashkeel[value] {
			r[index] = value
		}
	}

	return r
}

// ApplyTashkeel applies the tashkeel on a string
func ApplyTashkeel(table []rune, s string) string {

	sRune := []rune(s)

	n1, n2 := len(sRune), len(table)
	i, j, k := 0, 0, 0

	r := make([]rune, n1+n2, n1+n2)

	reverseAny(table)

	// Traverse both array
	for i < n1 && j < n2 {

		r[k] = sRune[i]
		k++
		i++

		if table[j] != rune(0) {
			r[k] = table[j]
			k++
			j += 2
		} else {
			j++
		}

	}

	// Store remaining elements of first array
	for i < n1 {
		r[k] = sRune[i]
		k++
		i++
	}

	// ignore trailing empty 0 from tashkeel table
	trimmed := make([]rune, k)
	copy(trimmed, r)
	return string(trimmed)
}

var tashkeels = []rune{
	fatha, fathataan, damma, dammataan, kasra, kasrataan, shadda, sukun,
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
