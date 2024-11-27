package tui

import (
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// StyleOption type is a function that takes a lipgloss style as input and returns a lipgloss style.
// It is used to apply different styles to a lipgloss style.
type StyleOption func(lipgloss.Style) lipgloss.Style

// Config function configures a lipgloss style.
// it takes a pointer to a lipgloss style and a list of style options as input.
// The style options are functions that take a lipgloss style as input and return a lipgloss style.
// The Config function applies the style options to the lipgloss style.
func Config(s *lipgloss.Style, options ...StyleOption) {
	st := *s
	for _, option := range options {
		st = option(st)
	}
	reflect.ValueOf(s).Elem().Set(reflect.ValueOf(st))
}

// NewStyle function returns a lipgloss style.
// It takes a list of style options as input and returns a lipgloss style.
// The style options are functions that take a lipgloss style as input and return a lipgloss style.
// The NewStyle function applies the style options to the lipgloss style and returns the modified style.
func NewStyle(options ...StyleOption) lipgloss.Style {
	s := lipgloss.NewStyle()
	Config(&s, options...)
	return s
}

// Render function returns a styled string.
// It takes a string and a list of style options as input and returns a styled string.
// The style options are functions that take a lipgloss style as input and return a lipgloss style.
// The Render function applies the style options to the string and returns the styled string.
func Render(text string, options ...StyleOption) string {
	return NewStyle(options...).Render(text)
}

// ConcatWith function concatenates a list of strings to a lipgloss style string value
// with the provided separator.
// It takes a pointer to a lipgloss style, a separator string, and a list of strings as input.
func ConcatWith(s *lipgloss.Style, sep string, strs ...string) {
	Config(s, func(st lipgloss.Style) lipgloss.Style {
		if s.Value() != "" {
			strs = append([]string{s.Value()}, strs...)
		}
		return st.SetString(strings.Join(strs, sep))
	})
}

// Concat function concatenates a list of strings to a lipgloss style string value.
// It takes a pointer to a lipgloss style and a list of strings as input.
func Concat(s *lipgloss.Style, strs ...string) {
	ConcatWith(s, "", strs...)
}

// ConcatLn function concatenates a list of strings to a lipgloss style string value
// with a newline separator ("\n").
// It takes a pointer to a lipgloss style and a list of strings as input.
func ConcatLn(s *lipgloss.Style, strs ...string) {
	ConcatWith(s, "\n", strs...)
}
