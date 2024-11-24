package tui

import (
	"fmt"
	"math"

	"github.com/charmbracelet/lipgloss"
)

// Component is an interface that defines a component that can be rendered
// to the terminal.
//
// Components are the building blocks of the TUI library. They can be
// rendered to the terminal and can be composed together to create more
// complex components.
//
// A componetn must implement the Buffer method that returns a Buffer
// that can be rendered to the terminal.
// And a String method that returns the string representation of the component.
type Component interface {
	Buf() *Buffer
	String() string
}

// StyleOption is a function that takes a lipgloss.Style and returns a
// lipgloss.Style. It is used to configure the style of a component.
type StyleOption func(s lipgloss.Style) lipgloss.Style

// BufferType is an enum that defines the type of buffer.
// It can be a normal buffer or a container buffer.
// A normal buffer is a buffer that contains text.
// A container buffer is a buffer that contains other components.
// If a buffer is a container buffer, it can be have an orientation
// that defines how the components are arranged.
type BufferType int

const (
	TextBuffer      BufferType = iota // TextBuffer is a normal buffer, it contains text
	ContainerBuffer                   // ContainerBuffer is a container buffer, it contains other components
)

// Orientation is an enum that defines the orientation of a container buffer.
// It can be vertical or horizontal.
// If the orientation is vertical, the components are arranged vertically.
// If the orientation is horizontal, the components are arranged horizontally.
type Orientation int

const (
	Vertical   Orientation = iota // Vertical orientation arranges the components vertically
	Horizontal                    // Horizontal orientation arranges the components horizontally
)

// New creates a new Buffer.
// The buffer is initialized with the default values.
// The default values are:
//   - text: empty string
//   - style: default lipgloss style
//   - bufferType: TextBuffer
//   - bufferOrientation: Vertical
//   - componentsAlignment: Left
//   - components: empty slice
func New() *Buffer {
	b := new(Buffer)
	b.style = lipgloss.NewStyle()
	b.bufferType = TextBuffer
	b.bufferOrientation = Vertical
	b.componentsAlignment = lipgloss.Left
	b.components = make([]Component, 0)
	return b
}

// Buffer is a struct that represents a buffer.
// A buffer is a container that can contain text or other components.
// A buffer can be rendered to the terminal or converted to a string.
// A buffer can be configured with a style (lipgloss.Style)
//
// There are two types of buffers:
//   - TextBuffer: a buffer that contains text
//   - ContainerBuffer: a buffer that contains other components
//
// If a buffer is a TextBuffer, it can have a style that defines how the text
// is rendered.
// If a buffer is a ContainerBuffer, it can have an orientation that defines
// how the components are arranged.
// The two orientations are:
//   - Vertical: the components are arranged vertically
//   - Horizontal: the components are arranged horizontally
//
// When a buffer is rendered to the terminal, the components are aligned
// according to the componentsAlignment field.
// The componentsAlignment field can be:
//   - Left: the components are aligned to the left
//   - Center: the components are aligned to the center
//   - Right: the components are aligned to the right
//   - Top: the components are aligned to the top
//   - Bottom: the components are aligned to the bottom
//
// Note that the componentsAlignment field is only used when the buffer is a
// ContainerBuffer. If the Orientation is Vertical, the components are aligned
// horizontally. If the Orientation is Horizontal, the components are aligned
// vertically.
//
// A buffer can have a width and a height. The width and height define the
// size of the buffer. The size of the buffer is the size of the content
// plus the size of the frame.
// If the buffer is a ContainerBuffer, and the Width or Height are set, the
// children components are resized to fit the buffer.
type Buffer struct {
	style               lipgloss.Style    // style of the buffer
	bufferType          BufferType        // type of the buffer
	bufferOrientation   Orientation       // orientation of the buffer
	componentsAlignment lipgloss.Position // alignment of the components
	components          []Component       // components of the buffer
}

// Add adds a string to the buffer.
// The string can contain format specifiers.
// The format specifiers are replaced by the arguments.
// If the buffer alrady contains a string, the new string is appended to the
// existing string.
//
// If the buffer is a TextBuffer, the string represents the content of the buffer.
// If the buffer is a ContainerBuffer, the string is rendered only if the buffer
// does not contain any components.
func (b *Buffer) Add(s string, args ...any) {
	b.style = b.style.SetString(fmt.Sprintf("%s%s", b.style.Value(), fmt.Sprintf(s, args...)))
}

// Addln adds a string to the buffer in a new line.
// The string can contain format specifiers.
// The format specifiers are replaced by the arguments.
// If the buffer alrady contains a string, the new string is appended to the
// existing string.
// A new line is added before the string.
//
// If the buffer is a TextBuffer, the string represents the content of the buffer.
// If the buffer is a ContainerBuffer, the string is rendered only if the buffer
// does not contain any components.
func (b *Buffer) Addln(s string, args ...any) {
	if b.style.Value() != "" {
		b.Add("\n")
	}
	b.Add(s, args...)
}

// AddC adds a component to the buffer.
// The component is appended to the list of components.
// If the buffer is a ContainerBuffer, the component is rendered according to
// the orientation of the buffer.
// If the buffer is a TextBuffer, the component is rendered only if the buffer
// does not contain any text, otherwise the component is ignored.
func (b *Buffer) AddC(c ...Component) {
	b.components = append(b.components, c...)
}

// Type is used to get or set the type of the buffer.
// If the bufferType argument is provided, the buffer type is set to the
// provided value.
// If more than one argument is provided, only the first argument is used.
//
// If the bufferType argument is not provided, the current buffer type is returned.
//
// The buffer type can be:
//   - TextBuffer: a buffer that contains text
//   - ContainerBuffer: a buffer that contains other components
//
// The function always returns the current buffer type.
func (b *Buffer) Type(bufferType ...BufferType) BufferType {
	if len(bufferType) > 0 {
		b.bufferType = bufferType[0]
	}
	return b.bufferType
}

// Orientation is used to get or set the orientation of the buffer.
// It is only used when the buffer is a ContainerBuffer.
// If the orientation argument is provided, the orientation of the buffer is
// set to the provided value.
// If more than one argument is provided, only the first argument is used.
//
// If the orientation argument is not provided, the current orientation of the
// buffer is returned.
//
// The orientation of the buffer can be:
//   - Vertical: the components are arranged vertically
//   - Horizontal: the components are arranged horizontally
//
// The function always returns the current orientation of the buffer.
func (b *Buffer) Orientation(orientation ...Orientation) Orientation {
	if len(orientation) > 0 {
		b.bufferOrientation = orientation[0]
	}
	return b.bufferOrientation
}

// Alignment is used to get or set the alignment of the components of the buffer.
// It is only used when the buffer is a ContainerBuffer.
// For TextBuffer, use the Style method to set the alignment of the text.
// If the align argument is provided, the alignment of the components is set to
// the provided value.
// If more than one argument is provided, only the first argument is used.
//
// If the align argument is not provided, the current alignment of the components
// is returned.
//
// The alignment of the components can be:
//   - Center: the components are aligned to the center (for both horizontal and vertical horientations)
//   - Left: the components are aligned to the left (for vertical orientation)
//   - Right: the components are aligned to the right (for vertical orientation)
//   - Top: the components are aligned to the top (for horizontal orientation)
//   - Bottom: the components are aligned to the bottom (for horizontal orientation)
//
// The function always returns the current alignment of the components.
func (b *Buffer) Alignment(align ...lipgloss.Position) lipgloss.Position {
	if len(align) > 0 {
		b.componentsAlignment = align[0]
	}
	return b.componentsAlignment
}

// Width is used to get or set the width of the buffer.
// If the width argument is provided, the width of the buffer is set to the
// provided value.
// If more than one argument is provided, only the first argument is used.
// To unset the width of the buffer, set the width to 0.
// The width comprehends the content width plus the horizontal paddings.
// It is used todefined where the text is wrapped (TextBuffer) or to resize
// the children components (ContainerBuffer).
//
// If the width argument is not provided, the current width of the buffer is returned.
//
// The function always returns the current width of the buffer.
func (b *Buffer) Width(width ...int) int {
	if len(width) > 0 {
		b.style = b.style.Width(width[0])
	}
	return b.style.GetWidth()
}

// Height is used to get or set the height of the buffer.
// If the height argument is provided, the height of the buffer is set to the
// provided value.
// If more than one argument is provided, only the first argument is used.
// To unset the height of the buffer, set the height to 0.
// The height comprehends the content height plus the vertical paddings.
// If the content heigth is minor than the buffer height, the content is
// filled with empty lines and the vertical paddings, it will be aligned
// according to the vertical alignment.
//
// If the height argument is not provided, the current height of the buffer is returned.
//
// The function always returns the current height of the buffer.
func (b *Buffer) Height(height ...int) int {
	if len(height) > 0 {
		b.style = b.style.Height(height[0])
	}
	return b.style.GetHeight()
}

// Margin is used to get or set the margins of the buffer.
// If the margins argument is provided, the margins of the buffer are set to the
// provided values.
//   - With no arguments, the function returns the horizontal and vertical margins.
//   - With one argument, the value is applied to all sides.
//   - With two arguments, the value is applied to the vertical and horizontal sides, in that order.
//   - With three arguments, the value is applied to the top side, the horizontal sides, and the bottom side, in that order.
//   - With four arguments, the value is applied clockwise starting from the top side, followed by the right side, then the bottom, and finally the left.
//   - With more than four arguments no margin will be added.
//
// The function always returns the horizontal and vertical margins of the buffer.
func (b *Buffer) Margin(margins ...int) (int, int) {
	if len(margins) > 0 {
		b.style = b.style.Margin(margins...)
	}
	return b.style.GetHorizontalMargins(), b.style.GetVerticalMargins()
}

// Padding is used to get or set the paddings of the buffer.
// If the paddings argument is provided, the paddings of the buffer are set to the
// provided values.
//   - With no arguments, the function returns the horizontal and vertical paddings.
//   - With one argument, the value is applied to all sides.
//   - With two arguments, the value is applied to the vertical and horizontal sides, in that order.
//   - With three arguments, the value is applied to the top side, the horizontal sides, and the bottom side, in that order.
//   - With four arguments, the value is applied clockwise starting from the top side, followed by the right side, then the bottom, and finally the left.
//   - With more than four arguments no padding will be added.
//
// The function always returns the horizontal and vertical paddings of the buffer.
func (b *Buffer) Padding(paddings ...int) (int, int) {
	if len(paddings) > 0 {
		b.style = b.style.Padding(paddings...)
	}
	return b.style.GetHorizontalPadding(), b.style.GetVerticalPadding()
}

// Style is used to get or set the style of the buffer.
// If the style argument is provided, the style of the buffer is set to the
// provided value.
// If more than one argument is provided, only the first argument is used.
//
// If the style argument is not provided, the current style of the buffer is returned.
//
// The function always returns the current style of the buffer.
func (b *Buffer) Style(style ...lipgloss.Style) lipgloss.Style {
	if len(style) > 0 {
		b.style = style[0].SetString(b.style.Value())
	}
	return b.style
}

// ConfigStyle is used to configure the style of the buffer.
// The function takes one or more StyleOption functions as arguments.
// Each StyleOption function takes a lipgloss.Style and returns a lipgloss.Style.
// The StyleOption functions are applied to the current style of the buffer.
// The function returns the new style of the buffer.
func (b *Buffer) ConfigStyle(configs ...StyleOption) lipgloss.Style {
	for _, config := range configs {
		b.style = config(b.style)
	}
	return b.style
}

// Size returns the size of the buffer.
// The size of the buffer is the size of the content plus the size of the frame.
// The size is returned as a pair of integers: the width and the height.
func (b *Buffer) Size() (int, int) {
	return lipgloss.Size(b.style.String())
}

// FrameSize returns the size of the frame of the buffer.
// The frame size is the size of the buffer minus the size of the content.
// The frame size is returned as a pair of integers: the width and the height.
func (b *Buffer) FrameSize() (int, int) {
	return b.style.GetHorizontalFrameSize(), b.style.GetVerticalFrameSize()
}

// ContentSize returns the size of the content of the buffer.
// The content size is the size of the buffer minus the size of the frame.
// The content size is returned as a pair of integers: the width and the height.
func (b *Buffer) ContentSize() (int, int) {
	w, h := b.Size()
	fw, fh := b.FrameSize()
	return w - fw, h - fh
}

// UnsetStr unsets the string of the buffer.
// The string of the buffer is set to an empty string.
func (b *Buffer) UnsetStr() {
	b.style = b.style.UnsetString()
}

// UnsetStyle unsets the style of the buffer.
// The style of the buffer is set to the default lipgloss style.
func (b *Buffer) UnsetStyle() {
	b.style = lipgloss.NewStyle().SetString(b.style.Value())
}

// UnsetComponents unsets the components of the buffer.
// The components of the buffer are set to an empty slice.
func (b *Buffer) UnsetComponents() {
	b.components = make([]Component, 0)
}

// Reset resets the buffer.
// The string, style, and components of the buffer are unset.
// The buffer will maintain the same type, orientation, and alignment.
func (b *Buffer) Reset() {
	b.UnsetStr()
	b.UnsetStyle()
	b.UnsetComponents()
}

// Buf returns the buffer itself.
func (b *Buffer) Buf() *Buffer {
	return b
}

// String returns the string representation of the buffer.
//
// If the buffer is a TextBuffer, the string is returned according to the style.
// If there are not strings in the buffer, the components are rendered.
// In this case the components are attached one by one in the same line as a
// unique string.
//
// If the buffer is a ContainerBuffer, the components are rendered according to
// the orientation and alignment of the buffer.
// If there are not components in the buffer, the buffer searches for strings previously
// set and returns the result according to the style. (like a TextBuffer)
// The components are resized to fit the buffer.
// If the width or height of the buffer is not set, the size of the terminal
// is used.
// If the total dimensions of the components are minor than the buffer dimensions,
// the entire buffer is risized to fit as minor space as possible.
// The components are aligned according to the componentsAlignment field.
// Every component is rendered like an independent buffer.
// The components are joined according to the orientation of the buffer.
// The result is returned as a string.
// If both strings and components are not set, an empty string is returned.
func (b *Buffer) String() string {
	elNum := len(b.components)

	if b.style.Value() == "" && elNum == 0 { // If the buffer is empty, return an empty string
		return ""
	}

	switch b.bufferType {
	case ContainerBuffer:
		if elNum > 0 { // If the buffer contains components, render the components
			b.UnsetStr() // Unset the string of the buffer

			// If the width or height of the buffer is not set, use the terminal size
			if b.Width() == 0 || b.Height() == 0 {
				tw, th := GetTerminalSize()
				fw, fh := b.FrameSize()
				if b.Width() == 0 && tw > 0 {
					b.Width(tw - fw)
				}
				if b.Height() == 0 && th > 0 {
					b.Height(th - fh)
				}
			}

			// Get the padding of the buffer
			paddingHor, paddingVer := b.Padding()

			buffers := make([]*Buffer, elNum)   // transform the components into buffers
			dimensions := make([][2]int, elNum) // store the dimensions of the components
			result := make([]string, elNum)     // store the resulting string of the components

			// totalWidth and totalHeight store the total dimensions of the components
			// maxWidth and maxHeight store the maximum dimensions of the components
			// mWIndex stores the index of the component with the maximum width
			var totalWidth, totalHeight, maxWidth, maxHeight, mWIndex int

			for i, c := range b.components {
				buffers[i] = c.Buf()         // transform the component into a buffer
				w, h := buffers[i].Size()    // get the dimensions of the buffer
				dimensions[i] = [2]int{w, h} // store the dimensions of the buffer
				totalWidth += w              // add the width to the total width
				totalHeight += h             // add the height to the total height
				if w > maxWidth {            // get the maximum width and the index of the component with the maximum width
					maxWidth = w
					mWIndex = i
				}
				if h > maxHeight { // get the maximum height
					maxHeight = h
				}
			}

			switch b.bufferOrientation {
			case Horizontal:
				// deltaWidth stores the difference between the total width of the components
				// and the buffer width (used only when the buffer width is set and is minor
				// than the total width of the components)
				var deltaWidth int

				// If the total width of the components is minor than the buffer width,
				// resize the buffer to fit the components
				if (b.Width() > 0 && totalWidth < b.Width()) || b.Width() == 0 {
					b.Width(totalWidth)
					deltaWidth = 0
				} else {
					// If the total width of the components is major than the buffer width,
					// resize the components to fit the buffer
					deltaWidth = totalWidth - b.Width() - paddingHor
				}

				// If the total height of the components is minor than the buffer height,
				// resize the buffer to fit the components
				if (maxHeight < b.Height() && b.Height() > 0) || b.Height() == 0 {
					b.Height(maxHeight + paddingVer)
				}

				// calculate how much width should be subtracted from each component
				// to fit the buffer
				subtractWidth := int(math.Floor(float64(deltaWidth) / float64(elNum)))
				subtractWidthDelta := deltaWidth - subtractWidth*elNum

				for i, buf := range buffers {
					// resize the component to fit the buffer
					bfw, _ := buf.FrameSize()
					w := dimensions[i][0] - subtractWidth - bfw
					// if the component has the maximum width, subtract the remaining width
					// to fit the buffer
					if i == mWIndex && subtractWidthDelta > 0 {
						w -= subtractWidthDelta
					}
					buf.Width(w)
					result[i] = buf.String()
				}

				// add the components in form of string to the buffer
				b.Add(lipgloss.JoinHorizontal(b.componentsAlignment, result...))
			default: // Vertical
				// If the maximum width of the components is minor than the buffer width,
				// resize the entire buffer to fit the components
				if (maxWidth < b.Width() && b.Width() > 0) || b.Width() == 0 {
					b.Width(maxWidth + paddingHor)
				}

				// If the total height of the components is minor than the buffer height,
				// resize the buffer to fit the components
				if (totalHeight < b.Height() && b.Height() > 0) || b.Height() == 0 {
					b.Height(totalHeight + paddingVer)
				}

				for i, buf := range buffers {
					// resize the component to fit the buffer
					bfw, _ := buf.FrameSize()
					if b.Width() > 0 && dimensions[i][0] > b.Width() {
						buf.Width(b.Width() - paddingHor - bfw)
					}
					result[i] = buf.String()
				}

				// add the components in form of string to the buffer
				b.Add(lipgloss.JoinVertical(b.componentsAlignment, result...))
			}
		}
	default: // TextBuffer
		// if the buffer contains components and the buffer does not contain text, render the components
		if elNum > 0 && b.style.Value() == "" {
			for _, c := range b.components {
				b.Add(c.String())
			}
		}
	}

	// return the buffer string according to the style
	return b.style.String()
}

// Print prints the buffer to the terminal.
func (b *Buffer) Print() {
	fmt.Print(b.String())
}
