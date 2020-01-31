package goarabic

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

	// Store remaining elements of second array
	for j < n2 {
		if table[j] != rune(0) {
			r[k] = table[j]
			k++
		}
		j++
	}

	return string(r)

}

var tashkeels = []rune{
	FATHA, FATHATAN, DAMMA, DAMMATAN, KASRA, KASRATAN, SHADDA, SUKUN,
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
