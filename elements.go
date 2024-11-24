package tui

import "github.com/charmbracelet/lipgloss"

// NewTextElement creates a new TextElement
// It returns a new TextElement with the default values
// The default values are:
//   - string: empty string
//   - style: lipgloss.NewStyle().Inline(true)
//   - bufferType: TextBuffer
//   - componentsAlignment: Left
func NewTextElement() *TextElement {
	return &TextElement{
		Buffer: Buffer{
			style:               lipgloss.NewStyle().Inline(true),
			bufferType:          TextBuffer,
			componentsAlignment: lipgloss.Left,
		},
	}
}

// TextElement is a struct that represents a text element
// It has a Buffer embedded in it
// It is used to create only text elements
// Now the buffer type is used to determine if the element is a inline (TextBuffer) or a block (ContainerBuffer) element
// Now The components alignment is used to determine the alignment of the string generated
// The components are used to print string directly in the buffer value
type TextElement struct {
	Buffer
}

// AddC adds a component to the TextElement in form of a string
// It receives a slice of Component
// It adds the string representation of the component to the TextElement
// This method overrides the AddC method from the Buffer struct
// The TextElement supports only string componentss
func (t *TextElement) AddC(c ...Component) {
	for _, comp := range c {
		t.Add(comp.String())
	}
}

// AddlnC adds a component to the TextElement in form of a string with a newline character at the end
// It receives a slice of Component
// It adds the string representation of the component to the TextElement with a newline character at the end
// This method overrides the AddlnC method from the Buffer struct
// The TextElement supports only string components
func (t *TextElement) AddlnC(c ...Component) {
	for _, comp := range c {
		t.Addln(comp.String())
	}
}

// Type sets the buffer type of the TextElement
// It receives a BufferType
// It sets the buffer type of the TextElement to the received value
// It returns the buffer type of the TextElement
// This method overrides the Type method from the Buffer struct
// The TextElement supports only TextBuffer and ContainerBuffer buffer types
// If the buffer type is TextBuffer, the style is set to inline
// If the buffer type is ContainerBuffer, the style is set to block
func (t *TextElement) Type(ty ...BufferType) BufferType {
	if len(ty) > 0 {
		t.bufferType = ty[0]
	}

	switch t.bufferType {
	case TextBuffer:
		t.style = t.style.Inline(true)
	case ContainerBuffer:
		t.style = t.style.Inline(false)
	}

	return t.bufferType
}

// Alignment sets the alignment of the components of the TextElement
// It receives a Position
// It sets the alignment of the style of the TextElement to the received value
// This method overrides the Alignment method from the Buffer struct
// The TextElement supports alignment only for text uses
func (t *TextElement) Alignment(a ...lipgloss.Position) lipgloss.Position {
	if len(a) == 1 {
		t.componentsAlignment = a[0]
		t.style = t.style.Align(t.componentsAlignment)
	} else if len(a) == 2 {
		t.componentsAlignment = a[1]
		t.style = t.style.Align(a[0], a[1])
	}

	return t.componentsAlignment
}
