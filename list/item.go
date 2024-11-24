package list

import (
	"strings"

	"github.com/Tagliapietra96/tui"
	tea "github.com/charmbracelet/bubbletea"
)

// toggle function toggles the selected field of the item
// It receives a ListItem as input
// It returns the item model and a command
// it used in the update method of the ListItem
func toggle(l *ListItem) (tea.Model, tea.Cmd) {
	l.Toggle()
	return l, nil
}

// NewListItem creates a new ListItem
// It receives a ListType, an index, and a value
// It returns a new ListItem with the received values
// The ListType is used to determine the type of list the item belongs to
// The index is used to determine the position of the item in the list
// The value is the content of the item
// The prefix is used to determine the number of digits of the index
// If the prefix is not provided, it will be set to 0
// If the prefix is provided, it will be set to the received value
// If more than one prefix is provided, it will be set to the first value
func NewListItem(lType ListType, index int, value string, prefix ...int) *ListItem {
	l := new(ListItem)
	l.listType = lType
	l.index = index
	l.value = value
	if len(prefix) > 0 {
		l.prefix = prefix[0]
	}
	return l
}

// ListItem is a struct that represents a list item
// The list has the following fields:
//   - listType: the type of list the item belongs to (UnorderedList, OrderedList, CheckList, PlainList)
//   - selected: a boolean that indicates if the item is selected
//   - hovered: a boolean that indicates if the item is hovered
//   - interactive: a boolean that indicates if the item is in interactive mode
//   - index: the position of the item in the list
//   - prefix: the number of digits of the index
//   - value: the content of the item
//
// The ListItem is a model that implements the Model interface
// In the Update method, the ListItem handles the KeyMsg messages
// If the message is a space key, the ListItem toggles the selected field
type ListItem struct {
	listType    ListType
	selected    bool
	hovered     bool
	interactive bool
	index       int
	prefix      int
	value       string
}

// Type method sets or returns the type of the list the item belongs to
// It receives a ListType as input
// If the ListType is provided, it sets the listType field to the received value
// It returns the listType field
func (l *ListItem) Type(lt ...ListType) ListType {
	if len(lt) > 0 {
		l.listType = lt[0]
	}
	return l.listType
}

// Check method checks or unchecks the item
// It receives a boolean as input
// If the boolean is provided, it sets the selected field to the received value
// It returns the selected field
func (l *ListItem) Check(checked ...bool) bool {
	if len(checked) > 0 {
		l.selected = checked[0]
	}
	return l.selected
}

// Hover method hovers or unhovers the item
// It receives a boolean as input
// If the boolean is provided, it sets the hovered field to the received value
// It returns the hovered field
func (l *ListItem) Hover(hovered ...bool) bool {
	if len(hovered) > 0 {
		l.hovered = hovered[0]
	}
	return l.hovered
}

// Index method sets or returns the index of the item
// It receives an integer as input
// If the integer is provided, it sets the index field to the received value
// It returns the index field
func (l *ListItem) Index(index ...int) int {
	if len(index) > 0 {
		l.index = index[0]
	}
	return l.index
}

// Prefix method sets or returns the prefix of the item
// It receives an integer as input
// If the integer is provided, it sets the prefix field to the received value
// It returns the prefix field
func (l *ListItem) Prefix(prefix ...int) int {
	if len(prefix) > 0 {
		l.prefix = prefix[0]
	}
	return l.prefix
}

// ListType method sets or returns the type of the list the item belongs to
// It receives a ListType as input
// If the ListType is provided, it sets the listType field to the received value
// It returns the listType field
func (l *ListItem) Value(value ...string) string {
	if len(value) > 0 {
		l.value = strings.Join(value, " ")
	}
	return l.value
}

// Toggle method toggles the selected field
// It returns the selected field
// If the selected field is true, it sets it to false
// If the selected field is false, it sets it to true
func (l *ListItem) Toggle() bool {
	l.selected = !l.selected
	return l.selected
}

// Interactive method sets or returns the interactive field
// It receives a boolean as input
// If the boolean is provided, it sets the interactive field to the received value
// It returns the interactive field
func (l *ListItem) Interactive(interactive ...bool) bool {
	if len(interactive) > 0 {
		l.interactive = interactive[0]
	}
	return l.interactive
}

// Buf method returns the buffer of the item
// It returns a buffer with the content of the item
// If the item is hovered, the buffer is brigth
// If the item is not hovered, the buffer is muted light
// The buffer is created based on the type of list the item belongs to
//   - If the list is an unordered list, the buffer has a bullet point
//   - If the list is an ordered list, the buffer has a number
//   - If the list is a check list, the buffer has a checkbox
//   - If the list is a plain list, the buffer has no prefix
func (l *ListItem) Buf() *tui.Buffer {
	el := tui.NewTextElement()
	plain := tui.NewTextElement()
	if l.interactive {
		c := "> "
		if !l.hovered {
			c = "  "
		}
		el.AddC(tui.Muted(c))
	}

	switch l.listType {
	case UnorderedList:
		if l.interactive {
			if l.selected {
				el.AddC(tui.Success("✓ %s", l.value))
			} else if l.hovered {
				el.AddC(tui.Brigth("• %s", l.value))
			} else {
				el.AddC(tui.MutedLight("• %s", l.value))
			}
		} else {
			plain.Add("• %s", l.value)
			el.AddC(plain)
		}
	case OrderedList:
		el.AddC(tui.Bold("%s. ", tui.FormatIntWithPrefix(l.index+1, l.prefix)))
		if l.interactive {
			if l.selected {
				el.AddC(tui.Success("%s", l.value))
			} else if l.hovered {
				el.AddC(tui.Brigth("%s", l.value))
			} else {
				el.AddC(tui.MutedLight("%s", l.value))
			}
		} else {
			plain.Add("%s", l.value)
			el.AddC(plain)
		}
	case CheckList:
		if l.interactive {
			if l.selected {
				el.AddC(tui.Bold("["), tui.Accent("✕"), tui.Bold("] "), tui.Brigth("%s", l.value))
			} else {
				el.AddC(tui.Bold("[ ] "))
				if l.hovered {
					el.AddC(tui.Brigth("%s", l.value))
				} else {
					el.AddC(tui.MutedLight("%s", l.value))
				}
			}
		} else {
			if l.selected {
				el.AddC(tui.Bold("[x] "))
			} else {
				el.AddC(tui.Bold("[ ] "))
			}
			plain.Add("%s", l.value)
			el.AddC(plain)
		}
	default: // PlainList
		if l.interactive {
			if l.selected {
				el.AddC(tui.Success("%s", l.value))
			} else if l.hovered {
				el.AddC(tui.Brigth("%s", l.value))
			} else {
				el.AddC(tui.MutedLight("%s", l.value))
			}
		} else {
			plain.Add("%s", l.value)
			el.AddC(plain)
		}
	}

	return el.Buf()
}

// String method returns the string representation of the item
// It returns the string representation of the buffer of the item
func (l *ListItem) String() string {
	return l.Buf().String()
}

// Init method initializes the item model
// It returns nil
func (l *ListItem) Init() tea.Cmd {
	return nil
}

// Update method updates the item model
// It receives a message as input
// If the message is a KeyMsg, it handles the space key
// If the space key is pressed, it toggles the selected field
// It returns the item model and a command
// The command returns nil
func (l *ListItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			return toggle(l)
		}
		switch msg.String() {
		case " ", "space":
			return toggle(l)
		}
	}
	return l, nil
}

// View method returns the view of the item
// It returns the string representation of the item
func (l *ListItem) View() string {
	l.interactive = true
	return l.String()
}
