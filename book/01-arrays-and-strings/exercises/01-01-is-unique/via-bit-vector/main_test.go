// ENG:
// 		Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// 		cannot use additional data structures?
// 		Hints: #44, #117, #132
// RU:
// 		Реализуйте алгоритм, определяющий, все ли символы в строке встречаются только один раз.
// 		А если при этом запрещено использование дополнительных структур данных?

package main

import (
	"testing"
	"unicode"
)

func isUnique(s string) bool {
	var bitVector [4]uint64
	for _, c := range s {
		idx := getIdx(unicode.ToLower(c))
		if readBit(bitVector, idx) {
			return false
		}

		setBit(&bitVector, idx)
	}

	return true
}

func getIdx(char rune) uint8 {
	return uint8(char)
}

func setBit(bitVector *[4]uint64, idx uint8) {
	octet, bit := idx/64, idx%64
	(*bitVector)[octet] |= (1 << bit)
}

func readBit(bitVector [4]uint64, idx uint8) bool {
	octet, bit := idx/64, idx%64
	return (bitVector[octet]&(1<<bit) != 0)
}

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Basic cases
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "two identical characters",
			input:    "aa",
			expected: false,
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: true,
		},

		// Case sensitivity tests
		{
			name:     "different case same letter",
			input:    "Aa",
			expected: false, // ToLower makes them identical
		},
		{
			name:     "different case different letters",
			input:    "Ab",
			expected: true,
		},
		{
			name:     "mixed case unique",
			input:    "AbCdEf",
			expected: true,
		},
		{
			name:     "mixed case with duplicate",
			input:    "AbCdA",
			expected: false,
		},

		// Alphabetical strings
		{
			name:     "lowercase alphabet",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: true,
		},
		{
			name:     "alphabet with duplicate",
			input:    "abcdefghijklmnopqrstuvwxyza",
			expected: false,
		},
		{
			name:     "full latin alphabet mixed case",
			input:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: false, // duplicates due to ToLower
		},

		// Digits
		{
			name:     "unique digits",
			input:    "1234567890",
			expected: true,
		},
		{
			name:     "digits with duplicate",
			input:    "12345678901",
			expected: false,
		},
		{
			name:     "mixed letters and digits unique",
			input:    "abc123",
			expected: true,
		},
		{
			name:     "mixed letters and digits with duplicate",
			input:    "abc123a",
			expected: false,
		},

		// Special characters
		{
			name:     "unique punctuation symbols",
			input:    "!@#$%^&*()",
			expected: true,
		},
		{
			name:     "punctuation symbols with duplicate",
			input:    "!@#$%^&*()!",
			expected: false,
		},
		{
			name:     "with spaces",
			input:    "a b c",
			expected: false,
		},

		// ASCII boundary characters (codes 0-127)
		{
			name:     "ASCII control characters",
			input:    "\x00\x01\x02",
			expected: true,
		},
		{
			name:     "ASCII upper boundary characters",
			input:    "\x7E\x7F", // ~ and DEL
			expected: true,
		},
		{
			name:     "all ASCII punctuation symbols",
			input:    "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
			expected: true,
		},

		// ASCII codes and symbols
		{
			name:     "ASCII digits and letters",
			input:    "abc123XYZ",
			expected: true,
		},
		{
			name:     "string with control characters",
			input:    "a\t\n\rb",
			expected: true,
		},
		{
			name:     "all ASCII digits and operators",
			input:    "0123456789+-*/=",
			expected: true,
		},
		{
			name:     "ASCII quote characters",
			input:    "\"'`",
			expected: true,
		},

		// Edge cases
		{
			name:     "very long unique string",
			input:    "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;':\",./<>?",
			expected: true,
		},
		{
			name:     "long string with duplicate at end",
			input:    "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;':\",./<>?a",
			expected: false,
		},
		{
			name:     "string with tab and newline",
			input:    "a\tb\nc",
			expected: true,
		},
		{
			name:     "string with repeating special characters",
			input:    "a\tb\nc\t",
			expected: false,
		},

		// Real world examples
		{
			name:     "word with unique letters",
			input:    "world",
			expected: true,
		},
		{
			name:     "word with repeating letters",
			input:    "hello",
			expected: false,
		},
		{
			name:     "short sentence",
			input:    "The quick brown fox",
			expected: false, // has spaces and duplicates
		},

		// Numeric strings
		{
			name:     "phone number with duplicates",
			input:    "1234567890",
			expected: true,
		},
		{
			name:     "number with repeating digits",
			input:    "1122334455",
			expected: false,
		},

		// Empty and whitespace strings
		{
			name:     "only spaces",
			input:    "   ",
			expected: false,
		},
		{
			name:     "single space",
			input:    " ",
			expected: true,
		},
		{
			name:     "different types of whitespace",
			input:    " \t\n",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isUnique(tt.input)
			if result != tt.expected {
				t.Errorf("isUnique(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Test for ASCII boundary values
func TestIsUniqueASCIILimits(t *testing.T) {
	// Test ASCII characters with codes 0-127
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "ASCII lower boundary characters",
			input:    string(rune(0)) + string(rune(1)) + string(rune(2)),
			expected: true,
		},
		{
			name:     "ASCII upper boundary characters",
			input:    string(rune(125)) + string(rune(126)) + string(rune(127)), // }~DEL
			expected: true,
		},
		{
			name:     "boundary characters with duplicate",
			input:    string(rune(0)) + string(rune(127)) + string(rune(0)),
			expected: false,
		},
		{
			name:     "full ASCII printable range",
			input:    " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			expected: false, // has duplicates due to ToLower
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isUnique(tt.input)
			if result != tt.expected {
				t.Errorf("isUnique(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
