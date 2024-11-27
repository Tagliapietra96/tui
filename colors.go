package tui

import "github.com/charmbracelet/lipgloss"

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
