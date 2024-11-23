package tui

import "github.com/charmbracelet/lipgloss"

// colors
var (
	accentColor     = lipgloss.AdaptiveColor{Light: "201", Dark: "213"}
	brigthColor     = lipgloss.AdaptiveColor{Light: "0", Dark: "15"}
	mutedColor      = lipgloss.AdaptiveColor{Light: "244", Dark: "241"}
	mutedLightColor = lipgloss.AdaptiveColor{Light: "241", Dark: "248"}
	errorColor      = lipgloss.AdaptiveColor{Light: "160", Dark: "196"}
	successColor    = lipgloss.AdaptiveColor{Light: "22", Dark: "40"}
	warningColor    = lipgloss.AdaptiveColor{Light: "208", Dark: "214"}
	infoColor       = lipgloss.AdaptiveColor{Light: "33", Dark: "45"}
	linkColor       = lipgloss.AdaptiveColor{Light: "27", Dark: "33"}
)

// styles
var (
	// text styles
	TitleStyle      = lipgloss.NewStyle().Foreground(brigthColor).Bold(true).Margin(1, 0)
	BrigthStyle     = lipgloss.NewStyle().Foreground(brigthColor).Inline(true)
	MutedLightStyle = lipgloss.NewStyle().Foreground(mutedLightColor).Inline(true)
	MutedStyle      = lipgloss.NewStyle().Foreground(mutedColor).Inline(true)
	AccentStyle     = lipgloss.NewStyle().Foreground(accentColor).Inline(true)
	BoldStyle       = lipgloss.NewStyle().Bold(true).Inline(true)
	ItalicStyle     = lipgloss.NewStyle().Italic(true).Inline(true)
	UnderlineStyle  = lipgloss.NewStyle().Underline(true).Inline(true)

	// message styles
	SuccessStyle = lipgloss.NewStyle().Foreground(successColor).Inline(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(errorColor).Inline(true)
	WarningStyle = lipgloss.NewStyle().Foreground(warningColor).Inline(true)
	InfoStyle    = lipgloss.NewStyle().Foreground(infoColor).Inline(true)

	// element styles
	LinkStyle  = lipgloss.NewStyle().Foreground(linkColor).Underline(true).Inline(true)
	QuoteStyle = lipgloss.NewStyle().Foreground(mutedColor).Italic(true).Border(lipgloss.ThickBorder(), false, false, false, true).BorderForeground(mutedColor).PaddingLeft(2).Margin(1, 0)

	// input styles
	QuestionIconStyle = lipgloss.NewStyle().Foreground(successColor).Bold(true).Inline(true)
	QuestionStyle     = lipgloss.NewStyle().Bold(true).Inline(true)

	// table styles
	RowStyle        = lipgloss.NewStyle().Padding(0, 2, 0, 0)
	OddRowStyle     = RowStyle.Foreground(mutedLightColor)
	HeadingRowStyle = RowStyle.Bold(true).Foreground(brigthColor)
)
