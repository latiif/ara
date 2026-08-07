package arabic

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/bidi"
)

// mirrored maps bracket-like characters to their glyph in RTL context.
var mirrored = map[rune]rune{
	'(': ')', ')': '(',
	'[': ']', ']': '[',
	'{': '}', '}': '{',
	'<': '>', '>': '<',
	'«': '»', '»': '«',
}

func applyMirroring(s string) string {
	r := []rune(s)
	for pos, char := range r {
		if m, ok := mirrored[char]; ok {
			r[pos] = m
		}
	}
	return string(r)
}

const (
	fsi = '⁨'
	pdi = '⁩'
)

// isolateQuoted wraps each "..." piece so its text doesn't get mixed up with text outside the quotes.
func isolateQuoted(s string) string {
	var b strings.Builder
	inQuote := false
	for _, char := range s {
		if char != '"' {
			b.WriteRune(char)
			continue
		}
		if inQuote {
			b.WriteRune(pdi)
			b.WriteRune(char)
		} else {
			b.WriteRune(char)
			b.WriteRune(fsi)
		}
		inQuote = !inQuote
	}
	return b.String()
}

var isolateStripper = strings.NewReplacer(string(fsi), "", string(pdi), "")

// isDigits reports whether s is only digits.
func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// renderRTLRun shapes and reverses one right-to-left run. swapSides puts its trailing space first, for printing out of order.
func renderRTLRun(text string, swapSides bool) string {
	trimmed := strings.TrimSpace(text)
	lead := text[:strings.Index(text, trimmed)]
	trail := text[strings.Index(text, trimmed)+len(trimmed):]
	if swapSides {
		lead, trail = trail, lead
	}

	table := GenerateTashkeelTable(trimmed)
	base := RemoveTashkeel(trimmed)
	shaped := ToGlyph(Smooth(base))
	reversed := Reverse(shaped)
	withTashkeel := ApplyTashkeel(table, reversed)

	return lead + applyMirroring(withTashkeel) + trail
}

// RenderBidi reorders a line of mixed Arabic/Latin text into visual order using the Unicode Bidirectional Algorithm (UAX #9).
func RenderBidi(line string) string {
	var p bidi.Paragraph
	if _, err := p.SetString(isolateQuoted(line)); err != nil {
		return line
	}

	ordering, err := p.Order()
	if err != nil {
		return line
	}

	numRuns := ordering.NumRuns()
	texts := make([]string, numRuns)
	dirs := make([]bidi.Direction, numRuns)
	for runIndex := range numRuns {
		run := ordering.Run(runIndex)
		texts[runIndex] = isolateStripper.Replace(run.String())
		dirs[runIndex] = run.Direction()
	}

	// bidi.Order lists runs in logical order and doesn't reorder them for
	// display. A number sitting inside one Arabic sentence needs its two
	// RTL sides swapped around it, so we find and reverse those chains here.
	var b strings.Builder
	for runIndex := 0; runIndex < numRuns; {
		if dirs[runIndex] != bidi.RightToLeft {
			b.WriteString(texts[runIndex])
			runIndex++
			continue
		}

		chainEnd := runIndex
		scanIndex := runIndex + 1
		for scanIndex+1 < numRuns && isDigits(texts[scanIndex]) && dirs[scanIndex+1] == bidi.RightToLeft {
			chainEnd = scanIndex + 1
			scanIndex += 2
		}

		if chainEnd == runIndex {
			b.WriteString(renderRTLRun(texts[runIndex], false))
			runIndex++
			continue
		}

		for pos := chainEnd; pos >= runIndex; pos-- {
			if dirs[pos] == bidi.RightToLeft {
				b.WriteString(renderRTLRun(texts[pos], true))
			} else {
				b.WriteString(texts[pos])
			}
		}
		runIndex = chainEnd + 1
	}

	return b.String()
}
