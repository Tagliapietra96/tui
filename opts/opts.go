package opts

import (
	"strings"

	"github.com/Tagliapietra96/tui"
	"github.com/charmbracelet/lipgloss"
)

// style options
var (
	// Empty is a style option that sets the lipgloss style to an empty style.
	// it maintains the string value of the style. (useful for resetting the style)
	Empty tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return lipgloss.NewStyle().SetString(s.Value())
	}

	// Inline is a style option that sets the inline property of a lipgloss style.
	Inline tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Inline(true)
	}

	// Block is a style option that sets the inline property of a lipgloss style to false.
	Block tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Inline(false)
	}

	// Bold is a style option that sets the bold property of a lipgloss style.
	Bold tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Bold(true)
	}

	// Italic is a style option that sets the italic property of a lipgloss style.
	Italic tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Italic(true)
	}

	// Underline is a style option that sets the underline property of a lipgloss style.
	Underline tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Underline(true)
	}

	// StrikeThrough is a style option that sets the strikethrough property of a lipgloss style.
	StrikeThrough tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Strikethrough(true)
	}

	// Upper is a style option that sets the transform function of a lipgloss style to strings.ToUpper.
	// It transforms the text to uppercase.
	Upper tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Transform(strings.ToUpper)
	}

	// Lower is a style option that sets the transform function of a lipgloss style to strings.ToLower.
	// It transforms the text to lowercase.
	Lower tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Transform(strings.ToLower)
	}

	// NormalText is a style option that sets the bold, italic, underline, and strikethrough properties of a lipgloss style to false.
	// It also removes the transform function from the style.
	NormalText tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Bold(false).Strikethrough(false).Italic(false).Underline(false).Transform(func(str string) string {
			return str
		})
	}

	// Accent is a style option that sets the foreground color of a lipgloss style to the accent color.
	Accent tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorAccent)
	}

	// Bright is a style option that sets the foreground color of a lipgloss style to the bright color.
	Bright tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorBright)
	}

	// Muted is a style option that sets the foreground color of a lipgloss style to the muted color.
	Muted tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorMuted)
	}

	// LightMuted is a style option that sets the foreground color of a lipgloss style to the light muted color.
	LightMuted tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorLightMuted)
	}

	// Error is a style option that sets the foreground color of a lipgloss style to the error color.
	Error tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorError)
	}

	// Success is a style option that sets the foreground color of a lipgloss style to the success color.
	Success tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorSuccess)
	}

	// Warning is a style option that sets the foreground color of a lipgloss style to the warning color.
	Warning tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorWarning)
	}

	// Info is a style option that sets the foreground color of a lipgloss style to the info color.
	Info tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorInfo)
	}

	// Link is a style option that sets the foreground color of a lipgloss style to the link color and underlines the text.
	BackAccent tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorAccent)
	}

	// BackBright is a style option that sets the background color of a lipgloss style to the bright color.
	BackBright tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorBright)
	}

	// BackMuted is a style option that sets the background color of a lipgloss style to the muted color.
	BackMuted tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorMuted)
	}

	// BackLightMuted is a style option that sets the background color of a lipgloss style to the light muted color.
	BackLightMuted tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorLightMuted)
	}

	// BackError is a style option that sets the background color of a lipgloss style to the error color.
	BackError tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorError)
	}

	// BackSuccess is a style option that sets the background color of a lipgloss style to the success color.
	BackSuccess tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorSuccess)
	}

	// BackWarning is a style option that sets the background color of a lipgloss style to the warning color.
	BackWarning tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorWarning)
	}

	// BackInfo is a style option that sets the background color of a lipgloss style to the info color.
	BackInfo tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Background(tui.ColorInfo)
	}

	// Left is a style option that aligns the text to the left.
	Left tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignHorizontal(lipgloss.Left)
	}

	// HorCenter is a style option that aligns the text to the center horizontally.
	HorCenter tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignHorizontal(lipgloss.Center)
	}

	// Right is a style option that aligns the text to the right.
	Right tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignHorizontal(lipgloss.Right)
	}

	// Top is a style option that aligns the text to the top.
	Top tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignVertical(lipgloss.Top)
	}

	// VerCenter is a style option that aligns the text to the center vertically.
	VerCenter tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignVertical(lipgloss.Center)
	}

	// Bottom is a style option that aligns the text to the bottom.
	Bottom tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.AlignVertical(lipgloss.Bottom)
	}

	// Link is a style option that sets the foreground color of a lipgloss style to the link color and underlines the text.
	Link tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		return s.Foreground(tui.ColorLink).Underline(true).Inline(true)
	}

	// Quote is a style option that sets the style of a quote. It adds a border to the left side of the text.
	Quote tui.StyleOption = func(s lipgloss.Style) lipgloss.Style {
		s = Color(nil)(s)
		return s.Border(lipgloss.ThickBorder(), false, false, false, true).BorderForeground(tui.ColorMuted).PaddingLeft(2).Margin(1, 0)
	}
)

// Width returns a style option that sets the width of a lipgloss style.
// It takes an integer as input and sets the width of the lipgloss style.
// If the width is less than 0, it sets the width to 0.
func Width(width int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		if width < 0 {
			width = 0
		}
		return s.Width(width)
	}
}

// Height returns a style option that sets the height of a lipgloss style.
// It takes an integer as input and sets the height of the lipgloss style.
// If the height is less than 0, it sets the height to 0.
func Height(height int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		if height < 0 {
			height = 0
		}
		return s.Height(height)
	}
}

// Dim returns a style option that sets the width and height of a lipgloss style.
// It takes two integers as input and sets the width and height of the lipgloss style.
// If the width or height is less than 0, it sets the width or height to 0.
func Dim(width, height int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return Width(width)(Height(height)(s))
	}
}

// FitWidth returns a style option that sets the width of a lipgloss style.
// It takes an integer as input representing the expected width of the lipgloss style.
// The FitWidth function calculates the width of the lipgloss style by subtracting the
// horizontal border size and margins from the expected width.
// If the calculated width is less than 0, it sets the width to 0.
func FitWidth(expected int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return Width(expected - s.GetHorizontalBorderSize() - s.GetHorizontalMargins())(s)
	}
}

// FitHeight returns a style option that sets the height of a lipgloss style.
// It takes an integer as input representing the expected height of the lipgloss style.
// The FitHeight function calculates the height of the lipgloss style by subtracting the
// vertical border size and margins from the expected height.
// If the calculated height is less than 0, it sets the height to 0.
func FitHeight(expected int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return Height(expected - s.GetVerticalBorderSize() - s.GetVerticalMargins())(s)
	}
}

// FitDim returns a style option that sets the width and height of a lipgloss style.
// It takes two integers as input representing the expected width and height of the lipgloss style.
// The FitDim function calculates the width and height of the lipgloss style by subtracting the
// horizontal and vertical border sizes and margins from the expected width and height.
// If the calculated width or height is less than 0, it sets the width or height to 0.
func FitDim(width, height int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return FitWidth(width)(FitHeight(height)(s))
	}
}

// Margin returns a style option that sets the margin of a lipgloss style.
// It takes a list of integers as input and sets the margin of the lipgloss style.
//   - One integer: sets the margin on all sides.
//   - Two integers: sets the vertical and horizontal margins.
//   - Three integers: sets the top, horizontal, and bottom margins.
//   - Four integers: sets the top, right, bottom, and left margins.
//   - More than four integers: no effect.
//
// To unset a margin, set it to 0
func Margin(margins ...int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return s.Margin(margins...)
	}
}

// Padding returns a style option that sets the padding of a lipgloss style.
// It takes a list of integers as input and sets the padding of the lipgloss style.
//   - One integer: sets the padding on all sides.
//   - Two integers: sets the vertical and horizontal paddings.
//   - Three integers: sets the top, horizontal, and bottom paddings.
//   - Four integers: sets the top, right, bottom, and left paddings.
//   - More than four integers: no effect.
//
// To unset a padding, set it to 0
func Padding(paddings ...int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		return s.Padding(paddings...)
	}
}

// Color returns a style option that sets the color of a lipgloss style.
// It takes a list of lipgloss terminal colors as input and sets the color of the lipgloss style.
//   - One color: sets the foreground color.
//   - Two colors: sets the foreground and background colors.
//   - Three colors: sets the foreground, background, and border foreground colors.
//   - Four colors: sets the foreground, background, border foreground, and border background colors.
//   - More than four colors: same as four colors.
//   - No colors: unset the colors.
//
// If a color is not provided (nil), it will be set to NoColor.
func Color(color ...lipgloss.TerminalColor) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		if len(color) == 0 {
			return s.Foreground(lipgloss.NoColor{}).Background(lipgloss.NoColor{}).BorderForeground(lipgloss.NoColor{}).BorderBackground(lipgloss.NoColor{})
		}
		if len(color) > 0 {
			if color[0] != nil {
				s = s.Foreground(color[0])
			} else {
				s = s.Foreground(lipgloss.NoColor{})
			}
		}
		if len(color) > 1 {
			if color[1] != nil {
				s = s.Background(color[1])
			} else {
				s = s.Background(lipgloss.NoColor{})
			}
		}
		if len(color) > 2 {
			if color[2] != nil {
				s = s.BorderForeground(color[2])
			} else {
				s = s.BorderForeground(lipgloss.NoColor{})
			}
		}
		if len(color) > 3 {
			if color[3] != nil {
				s = s.BorderBackground(color[3])
			} else {
				s = s.BorderBackground(lipgloss.NoColor{})
			}
		}
		return s
	}
}

// Heading returns a style option that sets the style of a heading.
// It takes an integer as input and sets the style of the heading based on the level.
// The level determines the size and style of the heading.
//   - 1: sets the style of the heading to the largest size and adds a border at the bottom.
//   - 2: sets the style of the heading to the second largest size and underlines the text.
//   - 3: sets the style of the heading to the third largest size and transforms the text to uppercase.
//   - 4: sets the style of the heading to the fourth largest size and removes the underline.
//   - 5: sets the style of the heading to the fifth largest size and removes the border.
//   - More than 4: no effect.
func Heading(level int) tui.StyleOption {
	return func(s lipgloss.Style) lipgloss.Style {
		if level <= 5 && level > 0 {
			s = s.Foreground(tui.ColorBright).Bold(true).Inline(true)
			if level < 5 {
				s = s.Inline(false).MarginBottom(1)
			}
			if level < 4 {
				s = s.Transform(strings.ToUpper)
			}
			if level < 3 {
				s = s.Underline(true)
			}
			if level < 2 {
				s = s.MarginBottom(2).Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(tui.ColorLightMuted).Underline(false)
			}
		}
		return s
	}
}
