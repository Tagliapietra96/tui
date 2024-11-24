package tui

// Bright creates a new TextElement with the given string and args
// It returns a new TextElement with the bright text style
func Brigth(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(brigthStyle)
	return b
}

// MutedLight creates a new TextElement with the given string and args
// It returns a new TextElement with the muted light text style
func MutedLight(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(mutedLightStyle)
	return b
}

// Muted creates a new TextElement with the given string and args
// It returns a new TextElement with the muted text style
func Muted(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(mutedStyle)
	return b
}

// Accent creates a new TextElement with the given string and args
// It returns a new TextElement with the accent text style
func Accent(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(accentStyle)
	return b
}

// Success creates a new TextElement with the given string and args
// It returns a new TextElement with the success text style
func Success(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(successStyle)
	return b
}

// Info creates a new TextElement with the given string and args
// It returns a new TextElement with the info text style
func Info(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(infoStyle)
	return b
}

// Warning creates a new TextElement with the given string and args
// It returns a new TextElement with the warning text style
func Warning(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(warningStyle)
	return b
}

// Error creates a new TextElement with the given string and args
// It returns a new TextElement with the error text style
func Error(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(errorStyle)
	return b
}

// Bold creates a new TextElement with the given string and args
// It returns a new TextElement with the bold text style
func Bold(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(boldStyle)
	return b
}

// Italic creates a new TextElement with the given string and args
// It returns a new TextElement with the italic text style
func Italic(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(italicStyle)
	return b
}

// Underline creates a new TextElement with the given string and args
// It returns a new TextElement with the underline text style
func Underline(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(underlineStyle)
	return b
}

// Link creates a new TextElement with the given string and args
// It returns a new TextElement with the link text style
func Link(s string, args ...any) *TextElement {
	b := NewTextElement()
	b.Add(s, args...)
	b.Style(linkStyle)
	return b
}

// Quote creates a new TextElement with the given lines
// It returns a new TextElement with the quote text style
func Quote(lines ...string) *TextElement {
	q := NewTextElement()
	for _, l := range lines {
		q.Addln(l)
	}
	q.Style(quoteStyle)
	q.Type(ContainerBuffer)
	return q
}
