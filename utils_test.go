package tui

import (
	"testing"
)

func TestCleanString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "  Hello, World!  ",
			expected: "Hello, World!",
		},
		{
			input:    "\n\nHello, World!\n\n",
			expected: "Hello, World!",
		},
		{
			input:    "  Hello, \n  World!  ",
			expected: "Hello,\nWorld!",
		},
		{
			input:    "  \n  \n  ",
			expected: "",
		},
		{
			input:    "Hello,\n\nWorld!",
			expected: "Hello,\n\nWorld!",
		},
	}

	for _, test := range tests {
		result := CleanString(test.input)
		if result != test.expected {
			t.Errorf("CleanString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestFormatIntWithPrefix(t *testing.T) {
	tests := []struct {
		number   int
		minLen   int
		expected string
	}{
		{
			number:   42,
			minLen:   5,
			expected: "00042",
		},
		{
			number:   42,
			minLen:   2,
			expected: "42",
		},
		{
			number:   42,
			minLen:   0,
			expected: "42",
		},
		{
			number:   42,
			minLen:   1,
			expected: "42",
		},
		{
			number:   42,
			minLen:   3,
			expected: "042",
		},
	}

	for _, test := range tests {
		result := FormatIntWithPrefix(test.number, test.minLen)
		if result != test.expected {
			t.Errorf("FormatIntWithPrefix(%d, %d) = %q; expected %q", test.number, test.minLen, result, test.expected)
		}
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{
			input:    "Hello, World!",
			length:   5,
			expected: "He...",
		},
		{
			input:    "Hello, World!",
			length:   13,
			expected: "Hello, World!",
		},
		{
			input:    "Hello, World!",
			length:   0,
			expected: "",
		},
		{
			input:    "Hello, World!",
			length:   1,
			expected: "H",
		},
		{
			input:    "Hello, World!",
			length:   14,
			expected: "Hello, World!",
		},
	}

	for _, test := range tests {
		result := TruncateString(test.input, test.length)
		if result != test.expected {
			t.Errorf("TruncateString(%q, %d) = %q; expected %q", test.input, test.length, result, test.expected)
		}
	}
}
