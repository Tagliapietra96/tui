package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// H1 creates a new H1 heading with the given string and args
// It returns a new TextElement with the H1 heading style
// The args are optional and can be used to format the string
// The H1 style is uppercase, bold, and has a margin bottom of 2
// It also has a border at the bottom
func H1(s string, args ...any) *TextElement {
	h1 := NewTextElement()
	s = fmt.Sprintf(s, args...)
	h1.Add(strings.ToUpper(s))
	h1.Style(titleStyle.MarginBottom(2).Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(ColorLightMuted))
	h1.Type(ContainerBuffer)
	return h1
}

// H2 creates a new H2 heading with the given string and args
// It returns a new TextElement with the H2 heading style
// The args are optional and can be used to format the string
// The H2 style is uppercase and has a margin bottom of 1
// It also has an underline
func H2(s string, args ...any) *TextElement {
	h2 := NewTextElement()
	s = fmt.Sprintf(s, args...)
	h2.Add(strings.ToUpper(s))
	h2.Style(titleStyle.Underline(true))
	h2.Type(ContainerBuffer)
	return h2
}

// H3 creates a new H3 heading with the given string and args
// It returns a new TextElement with the H3 heading style
// The args are optional and can be used to format the string
// The H3 style is uppercase and has a margin bottom of 1
func H3(s string, args ...any) *TextElement {
	h3 := NewTextElement()
	s = fmt.Sprintf(s, args...)
	h3.Add(strings.ToUpper(s))
	h3.Style(titleStyle)
	h3.Type(ContainerBuffer)
	return h3
}

// H4 creates a new H4 heading with the given string and args
// It returns a new TextElement with the H4 heading style
// The args are optional and can be used to format the string
// The H4 style is bold and has a margin bottom of 1
func H4(s string, args ...any) *TextElement {
	h4 := NewTextElement()
	h4.Add(s, args...)
	h4.Style(titleStyle)
	h4.Type(ContainerBuffer)
	return h4
}

// H5 creates a new H5 heading with the given string and args
// It returns a new TextElement with the H5 heading style
// The args are optional and can be used to format the string
// The H5 style is bold and has no margins
func H5(s string, args ...any) *TextElement {
	h5 := NewTextElement()
	h5.Add(s, args...)
	h5.Style(titleStyle.MarginBottom(0))
	h5.Type(TextBuffer)
	return h5
}
