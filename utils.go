package tui

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

// FormatIntWithPrefix function formats an integer with a prefix.
// It takes an integer and a minimum length as input and returns a string with the number formatted.
// If the number is less than the minimum length, it will add zeros at the beginning of the number.
func FormatIntWithPrefix(number, minLength int) string {
	n := strconv.Itoa(number)
	nl := lipgloss.Width(n)
	var b strings.Builder
	if nl < minLength {
		for i := 0; i < minLength-nl; i++ {
			b.WriteString("0")
		}
	}

	b.WriteString(n)
	return b.String()
}

// CleanString function cleans a string.
// It takes a string as input and returns a string with the leading and trailing whitespaces removed.
// It also removes empty lines at the beginning and end of the string.
// It Useful for cleaning up unwonted margins and paddings in a styled string.
func CleanString(s string) string {
	// set a slice to track the result
	result := make([]string, 0)
	var found bool

	// iterate over the lines in the string
	// delete all white spaces in the beginning and end of the line
	// start appending the lines to the result slice once a non-empty line is found
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" && !found {
			continue
		}

		found = true
		result = append(result, line)
	}

	// iterate backwards over the result slice
	// remove empty lines at the end of the string
	// continue until a non-empty line is found
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "" {
			result = result[:i+1]
			break
		}
	}

	// ricombine the result slice into a single string
	return strings.Join(result, "\n")
}

// TruncateString function truncates a string.
// It takes a string, a length, and an optional truncation string as
// input and returns a truncated string.
// If the length of the string is greater than the specified length,
// it truncates the string and appends the truncation string.
// If the truncation string is not provided, it uses "..." as the
// default truncation string.
// Example:
//
//	TruncateString("Hello, World!", 10) => "Hello, ..."
//	TruncateString("Hello, World!", 10, "!!!") => "Hello, !!!"
//	TruncateString("Hello, World!", 10, "") => "Hello, Wor"
//	TruncateString("Hello, World!", 1) => "H"
//
// Note: The length is neither the number of characters nor the number
// of bytes in the string. Is the width of the string.
// The function uses lipgloss.Width to calculate the width of the string
// and the truncation string.
// The length of the truncation string is subtracted from the specified length
// to determine the maximum length of the string.
// If the length is less than the width of the truncation string, the function
// returns the truncated string without the truncation string.
// If the length is less than or equal to 0, the function returns an empty string.
func TruncateString(str string, length int, truncation ...string) string {
	// If the length is less than or equal to 0, return an empty string
	if length <= 0 {
		return ""
	}

	// set the truncation string
	var b strings.Builder
	tr := "..."
	if len(truncation) > 0 {
		tr = truncation[0]
	}

	// If the width of the string is greater than the specified length
	// truncate the string and append the truncation string
	if lipgloss.Width(str) > length {
		dots := Render(tr, func(s lipgloss.Style) lipgloss.Style {
			return s.Foreground(ColorMuted)
		})

		// If the length is less than the width of the truncation string
		// return the truncated string without the truncation string
		if length < lipgloss.Width(dots) {
			return str[:length]
		}

		// Otherwise, truncate the string and append the truncation string
		b.WriteString(str[:length-lipgloss.Width(dots)])
		b.WriteString(dots)
	} else {
		// if the width of the string is less than or equal to the specified length
		// return the string as is
		b.WriteString(str)
	}

	return b.String()
}

// getTerminalSize function returns the width and height of the terminal.
// It returns the width and height of the terminal as integers.
// If the terminal size cannot be determined, it returns 0, 0.
func getTerminalSize() (int, int) {
	w, h, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return 0, 0
	}

	return w, h
}
