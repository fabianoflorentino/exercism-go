package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := '❗'
	recommendationIndex := strings.IndexRune(log, recommendation)
	search := '🔍'
	weather := '☀'

	if strings.ContainsRune(log, recommendation) && recommendationIndex == 0 {
		return "recommendation"
	}
	if strings.ContainsRune(log, weather) && strings.ContainsRune(log, recommendation) {
		return "weather"
	}
	if strings.ContainsRune(log, search) {
		return "search"
	}
	if strings.ContainsRune(log, recommendation) {
		return "recommendation"
	}
	if strings.ContainsRune(log, weather) {
		return "weather"
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	replaceRune := func(r rune) rune {
		if r == oldRune {
			return newRune
		}

		return r
	}

	return strings.Map(replaceRune, log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
