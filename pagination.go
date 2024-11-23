package tui

import (
	"fmt"
	"math"

	"github.com/charmbracelet/lipgloss"
)

type Direction int

const (
	Vertical Direction = iota
	Horizontal
)

type Box struct {
	width     int
	height    int
	elements  []Component
	direction Direction
	buffer    *Buffer
}

func NewPage() *Box {
	b := NewBox()
	b.buffer.style = lipgloss.NewStyle().Margin(1, 2)
	return b
}

func NewBox() *Box {
	b := new(Box)
	b.direction = Vertical
	b.elements = make([]Component, 0)
	b.buffer = New()
	return b
}

func (b *Box) Add(c Component) *Box {
	b.elements = append(b.elements, c)
	return b
}

func (b *Box) SetDim(dimensions ...int) *Box {
	if len(dimensions) == 0 {
		b.width = 0
		b.height = 0
	}
	if len(dimensions) == 1 {
		b.width = dimensions[0]
		if b.width > 0 {
			return b
		}
	}
	if len(dimensions) == 2 {
		b.width = dimensions[0]
		b.height = dimensions[1]
		if b.width > 0 && b.height > 0 {
			return b
		}
	}

	var sw, maxh int

	for _, e := range b.elements {
		w, h := e.Buffer().Size()
		sw += w
		if h > maxh {
			maxh = h
		}
	}

	fw, fh := b.buffer.Frames()
	tw, th := GetTerminalSize()
	if b.width == 0 || b.width-fw > tw {
		b.width = tw
	}
	if b.height == 0 || b.height-th > th {
		b.height = th
	}

	if b.width-fw > sw {
		b.width = sw
	}

	if b.direction == Horizontal {
		if b.height-fh > maxh {
			b.height = maxh
		}

		delta := sw - b.width
		toSubtract := int(math.Floor(float64(delta) / float64(len(b.elements))))
		delta -= toSubtract * len(b.elements)
		refW := 0
		longest := []int{}
		for i, e := range b.elements {
			w, _ := e.Buffer().Size()
			if w > refW {
				refW = w
				longest = append(longest, i)
			}
			e.Buffer().Width(w - toSubtract)
		}
		if delta > 0 && len(longest) > 0 {
			w, _ := b.elements[longest[len(longest)-1]].Buffer().Size()
			b.elements[longest[len(longest)-1]].Buffer().Width(w - delta)
		}
	}

	return b
}

func (b *Box) SetDirection(direction Direction) *Box {
	b.direction = direction
	return b
}

func (b *Box) String() string {
	b.buffer.builder.Reset()
	b.SetDim(b.width, b.height)
	el := make([]string, len(b.elements))
	for i, e := range b.elements {
		el[i] = e.String()
	}

	if b.direction == Vertical {
		b.buffer.Add(lipgloss.JoinVertical(lipgloss.Left, el...))
	} else {
		b.buffer.Add(lipgloss.JoinHorizontal(lipgloss.Top, el...))
	}

	return b.buffer.String()
}

func (b *Box) Print() {
	fmt.Print(b.String())
}
