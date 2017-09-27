package ropeExperiment

// Copied from
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang

import (
	"math/rand"
	"time"
)

const asciiLetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var unicodeLetterBytes = [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '™', '¥', '§', '©', '®', '¼', '¾', 'Δ', 'Φ', 'Ω', 'θ', 'λ', 'Ϣ', '🐈', '🐑', '🐿', '🍩', '☕', '🍷', '🍺', '🔪', '🚇', '🚲', '🕐', '📷', '🔬'}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateASCIIString creates a UTF8-encoded string which contains `n`
// characters in the [a-zA-Z] range.
func GenerateASCIIString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(asciiLetterBytes) {
			b[i] = asciiLetterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// GenerateUnicodeString creates a UTF8-encoded string which contains `n`
// characters, some of which will be ASCII [a-z], others UTF8 characters,
// and other Unicode characters.
func GenerateUnicodeString(n int) string {
	b := make([]rune, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(unicodeLetterBytes) {
			b[i] = unicodeLetterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
