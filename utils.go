package tui

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/mattn/go-runewidth"
)

// colors
var (
	ColorAccent     = lipgloss.AdaptiveColor{Light: "201", Dark: "213"}
	ColorBright     = lipgloss.AdaptiveColor{Light: "0", Dark: "15"}
	ColorMuted      = lipgloss.AdaptiveColor{Light: "244", Dark: "241"}
	ColorLightMuted = lipgloss.AdaptiveColor{Light: "241", Dark: "248"}
	ColorError      = lipgloss.AdaptiveColor{Light: "160", Dark: "196"}
	ColorSuccess    = lipgloss.AdaptiveColor{Light: "22", Dark: "40"}
	ColorWarning    = lipgloss.AdaptiveColor{Light: "208", Dark: "214"}
	ColorInfo       = lipgloss.AdaptiveColor{Light: "33", Dark: "45"}
	ColorLink       = lipgloss.AdaptiveColor{Light: "27", Dark: "33"}
)

// styles
var (
	// text element styles
	titleStyle     = lipgloss.NewStyle().Foreground(ColorBright).Bold(true).Inline(false).MarginBottom(1)
	boldStyle      = lipgloss.NewStyle().Bold(true).Inline(true)
	italicStyle    = lipgloss.NewStyle().Italic(true).Inline(true)
	underlineStyle = lipgloss.NewStyle().Underline(true).Inline(true)
	linkStyle      = lipgloss.NewStyle().Foreground(ColorLink).Underline(true).Inline(true)
	quoteStyle     = lipgloss.NewStyle().Foreground(ColorMuted).Inline(false).Italic(true).Border(lipgloss.ThickBorder(), false, false, false, true).BorderForeground(ColorMuted).PaddingLeft(2).Margin(1, 0)

	// text color styles
	brigthStyle     = lipgloss.NewStyle().Foreground(ColorBright).Inline(true)
	mutedLightStyle = lipgloss.NewStyle().Foreground(ColorLightMuted).Inline(true)
	mutedStyle      = lipgloss.NewStyle().Foreground(ColorMuted).Inline(true)
	accentStyle     = lipgloss.NewStyle().Foreground(ColorAccent).Inline(true)
	successStyle    = lipgloss.NewStyle().Foreground(ColorSuccess).Inline(true)
	infoStyle       = lipgloss.NewStyle().Foreground(ColorInfo).Inline(true)
	warningStyle    = lipgloss.NewStyle().Foreground(ColorWarning).Inline(true)
	errorStyle      = lipgloss.NewStyle().Foreground(ColorError).Inline(true)

	// input styles
	questionIconStyle = lipgloss.NewStyle().Foreground(ColorSuccess).Bold(true).Inline(true)
	questionStyle     = lipgloss.NewStyle().Bold(true).Inline(true)

	// table styles
	rowStyle        = lipgloss.NewStyle().Padding(0, 2, 0, 0)
	oddRowStyle     = rowStyle.Foreground(ColorLightMuted)
	headingRowStyle = rowStyle.Bold(true).Foreground(ColorBright)
)

// GetNumberWidth function returns the width of a number.
// It takes an integer as input and returns the width of the number.
// The width of the number is the number of characters in the number.
func GetNumberWidth(number int) int {
	return runewidth.StringWidth(strconv.Itoa(number))
}

// FormatIntWithPrefix function formats an integer with a prefix.
// It takes an integer and a minimum length as input and returns a string with the number formatted.
// If the number is less than the minimum length, it will add zeros at the beginning of the number.
func FormatIntWithPrefix(number, minLength int) string {
	n := strconv.Itoa(number)
	nl := runewidth.StringWidth(n)
	var b strings.Builder
	if nl < minLength {
		for i := 0; i < minLength-nl; i++ {
			b.WriteString("0")
		}
	}

	b.WriteString(n)
	return b.String()
}

func RemoveFrames()

// GetTerminalSize function returns the width and height of the terminal.
// It returns the width and height of the terminal as integers.
// If the terminal size cannot be determined, it returns 0, 0.
func GetTerminalSize() (int, int) {
	w, h, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return 0, 0
	}

	return w, h
}
