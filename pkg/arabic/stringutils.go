// Package arabic contains utility functions for working with Arabic strings.
package arabic

import (
	"strings"
	"unicode"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReversePreservingNonArabic returns RTL arabic text while preserving non-araibc characters
func ReversePreservingNonArabic(s string) string {

	arabicLetters := constructAlphabetMap(alphabet)

	words := strings.Split(s, " ")
	stack := NewStack(len(words))

	formattedWords := make([]string, 0, len(words))

	for _, word := range words {

		isArabicWord := false
		for _, letter := range word {
			if _, ok := arabicLetters[letter]; ok {
				isArabicWord = true
				break
			}
		}
		if isArabicWord {

			if nonArabic := stack.Flush(); len(nonArabic) == 0 {
			} else {
				formattedWords = append(formattedWords, nonArabic...)
			}
			formattedWords = append(formattedWords, word)
		} else {
			stack.Push(Reverse(word))
		}
	}
	formattedWords = append(formattedWords, stack.Flush()...)
	return Reverse(strings.Join(formattedWords, " "))
}

// Helper functionality to allow for faster lookup of characters
var arabicLettersInUnicode map[rune]bool

func constructAlphabetMap(letters []Harf) map[rune]bool {
	if arabicLettersInUnicode != nil {
		return arabicLettersInUnicode
	}
	result := make(map[rune]bool)

	for _, harf := range letters {
		result[harf.Initial] = true
		result[harf.Isolated] = true
		result[harf.Unicode] = true
		result[harf.Medial] = true
		result[harf.Final] = true
	}

	arabicLettersInUnicode = result
	return result
}

// SmartLength returns the length of the given string
// without considering the Arabic Vowels (Tashkeel).
func SmartLength(s *string) int {
	// len() use int as return value, so we'd better follow for compatibility
	length := 0
	for _, value := range *s {
		if tashkeel[value] {
			continue
		}
		length++
	}
	return length
}

// RemoveTashkeel returns its argument as rune-wise string without Arabic vowels (Tashkeel).
func RemoveTashkeel(s string) string {
	// var r []rune
	// the capcity of the slice wont be greater than the length of the string itself
	// hence, cap = len(s)
	r := make([]rune, 0, len(s))

	for _, value := range s {
		if tashkeel[value] {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

// RemoveKashida returns its argument as rune-wise string without Arabic Tatweel character.
func RemoveKashida(s string) string {

	r := make([]rune, 0, len(s))

	for _, value := range s {
		if kasheeda.equals(value) {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

func isBetweenArabicLetters(previousChar, currentChar, nextChar rune) (previousIn, nextIn bool) {

	previousIn, nextIn = false, false

	for _, s := range alphabet {
		if s.equals(previousChar) { // previousChar in the Arabic Alphabet ?
			previousIn = true
		}

		if s.equals(nextChar) { // nextChar in the Arabic Alphabet ?
			nextIn = true
		}
	}

	return previousIn, nextIn
}

func getCharGlyph(previousChar, currentChar, nextChar rune) rune {

	glyph := currentChar

	previousIn, nextIn := isBetweenArabicLetters(previousChar, currentChar, nextChar)

	for _, s := range alphabet {

		if !s.equals(currentChar) { // currentChar in the Arabic Alphabet ?
			continue
		}

		if previousIn && nextIn { // between two Arabic Alphabet, return the medium glyph
			for s := range initialAfter {
				if s.equals(previousChar) {
					return getHarf(currentChar).Initial
				}
			}

			return getHarf(currentChar).Medial
		}

		if nextIn { // beginning (because the previous is not in the Arabic Alphabet)
			return getHarf(currentChar).Initial
		}

		if previousIn { // final (because the next is not in the Arabic Alphabet)
			for s := range initialAfter {
				if s.equals(previousChar) {
					return getHarf(currentChar).Isolated
				}
			}
			return getHarf(currentChar).Final
		}

		if !previousIn && !nextIn {
			return getHarf(currentChar).Isolated
		}

	}
	return glyph
}

// equals() return true if the given Arabic char is alphabetically equal to
// the current Harf regardless its shape (Glyph)
func (c *Harf) equals(char rune) bool {
	switch char {
	case c.Unicode:
		return true
	case c.Initial:
		return true
	case c.Isolated:
		return true
	case c.Medial:
		return true
	case c.Final:
		return true
	default:
		return false
	}
}

// getHarf gets the correspondent Harf for the given Arabic char
func getHarf(char rune) Harf {
	for _, s := range alphabet {
		if s.equals(char) {
			return s
		}
	}

	return Harf{Unicode: char, Isolated: char, Medial: char, Final: char}
}

// ToGlyph returns the glyph representation of the given text
func ToGlyph(text string) string {

	r := []rune(text)

	length := len(r)

	newText := make([]rune, 0, length)

	var previousChar, currentChar, nextChar rune

	for i := 0; i < length; i++ {

		// get the current char
		currentChar = r[i]

		// get the previous char
		if (i - 1) < 0 {
			previousChar = 0
		} else {
			previousChar = r[i-1]
		}

		// get the next char
		if (i + 1) <= length-1 {
			nextChar = r[i+1]
		} else {
			nextChar = 0
		}

		// get the current char representation or return the same if unnecessary
		glyph := getCharGlyph(previousChar, currentChar, nextChar)

		// append the new char representation to the newText
		newText = append(newText, glyph)
	}

	return string(newText)
}

// Smooth smartly places ligatures
func Smooth(s string) string {
	smoothed := ""

	runes := []rune(s)

	var curr, next Harf

	length := len(runes)

	for i := 0; i < length; i++ {

		if i < length-1 {
			curr = getHarf(runes[i])
			next = getHarf(runes[i+1])

			if secondPart, ok := ligatures[curr]; ok {
				if ligature, ok := secondPart[next]; ok {
					i++
					smoothed = smoothed + string(ligature.Unicode)
					continue
				}
			}
		}

		smoothed = smoothed + string(runes[i])
	}

	return smoothed
}

//breakLineAt finds the first whitespace to break a row
func breakLineAt(pos int, str string) (fst, snd string) {
	runes := []rune(str)

	for i := pos; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			fst = string(runes[i+1:])
			snd = string(runes[:i])
			return
		}
	}

	fst = str
	snd = ""
	return
}

//padRTL Fills the leading places with spaces to make an RTL text
func padRTL(strlen, size int, str string) string {
	plotsToFill := size - strlen
	padding := strings.Repeat(" ", plotsToFill)
	return padding + str
}

//MakeRTL displays the arabic text in RTL
func MakeRTL(size int, str string) string {
	strlen := SmartLength(&str)

	if strlen > size {
		// Finds a suitable place to split the long row
		fst, snd := breakLineAt(strlen-size, str)
		return padRTL(SmartLength(&fst), size, fst) + "\n" + MakeRTL(size, snd)
	}

	return padRTL(strlen, size, str)

}
