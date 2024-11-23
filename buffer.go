package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type Component interface {
	Buffer() *Buffer
	String() string
}

type StyleConfig func(s lipgloss.Style) lipgloss.Style

type Buffer struct {
	builder *strings.Builder
	style   lipgloss.Style
}

func New(lines ...string) *Buffer {
	b := new(Buffer)
	b.builder = new(strings.Builder)
	b.style = lipgloss.NewStyle()

	for i, line := range lines {
		if i == 0 {
			b.Add(line)
		} else {
			b.Addln(line)
		}
	}

	return b
}

func (b *Buffer) Add(s string, args ...any) *Buffer {
	b.builder.WriteString(fmt.Sprintf(s, args...))
	return b
}

func (b *Buffer) Addln(s string, args ...any) *Buffer {
	b.builder.WriteString("\n")
	b.Add(s, args...)
	return b
}

func (b *Buffer) AddC(c Component, args ...any) *Buffer {
	b.Add(c.String(), args...)
	return b
}

func (b *Buffer) AddlnC(c Component, args ...any) *Buffer {
	b.Addln(c.String(), args...)
	return b
}

func (b *Buffer) Width(w int, pos ...lipgloss.Position) *Buffer {
	b.style = b.style.Width(w)
	if len(pos) > 0 {
		b.style = b.style.AlignHorizontal(pos[0])
	}
	return b
}

func (b *Buffer) Height(h int, pos ...lipgloss.Position) *Buffer {
	b.style = b.style.Height(h)
	if len(pos) > 0 {
		b.style = b.style.AlignVertical(pos[0])
	}
	return b
}

func (b *Buffer) Margin(margins ...int) *Buffer {
	b.style = b.style.Margin(margins...)
	return b
}

func (b *Buffer) Padding(paddings ...int) *Buffer {
	b.style = b.style.Padding(paddings...)
	return b
}

func (b *Buffer) Style(style lipgloss.Style) *Buffer {
	b.style = style
	return b
}

func (b *Buffer) ConfigStyle(configs ...StyleConfig) {
	for _, config := range configs {
		b.style = config(b.style)
	}
}

func (b *Buffer) Dim() (int, int) {
	return StrWidth(b.builder.String()), StrHeight(b.builder.String())
}

func (b *Buffer) Frames() (int, int) {
	return b.style.GetHorizontalFrameSize(), b.style.GetVerticalFrameSize()
}

func (b *Buffer) Size() (int, int) {
	w, h := b.Dim()
	fw, fh := b.Frames()
	return w + fw, h + fh
}

func (b *Buffer) Len() int {
	return runewidth.StringWidth(b.builder.String())
}

func (b *Buffer) Reset() *Buffer {
	b.builder.Reset()
	b.style = lipgloss.NewStyle()
	return b
}

func (b *Buffer) String() string {
	return b.style.Render(b.builder.String())
}

func (b *Buffer) Print() {
	fmt.Print(b.String())
}

func (b *Buffer) Buffer() *Buffer {
	return b
}
