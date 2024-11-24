package list

import (
	"github.com/Tagliapietra96/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func changeFocus(l *List) tea.Cmd {
	return func() tea.Msg {
		for i, item := range l.filteredItems {
			item.Hover(i == l.current)
		}
		return nil
	}
}

// ListType is an enumeration of the types of lists
// It has the following values:
//   - UnorderedList: an unordered list
//   - OrderedList: an ordered list
//   - CheckList: a checklist
//   - PlainList: a plain list
type ListType int

const (
	UnorderedList ListType = iota // UnorderedList is an unordered list
	OrderedList                   // OrderedList is an ordered list
	CheckList                     // CheckList is a checklist
	PlainList                     // PlainList is a plain list
)

// ListFilterOption is a function that filters a list of list items
// It receives a slice of ListItem
// It returns a slice of ListItem
// The function is used to filter the list of items
type ListFilterOption func(l []*ListItem) []*ListItem

// GetList creates a list of list items
// It receives a ListType and a slice of strings
// It returns a slice of ListItem
// The ListType is used to determine the type of list
// The strings are the content of the items
// The function calculates the prefix of the items in case of an ordered list
func GetList(lType ListType, items ...string) []*ListItem {
	nel := len(items)
	list := make([]*ListItem, nel)
	prefix := tui.GetNumberWidth(nel)
	for i, item := range items {
		list[i] = NewListItem(lType, i, item, prefix)
	}
	return list
}

// NewList creates a new List
// It receives a ListType, a boolean, and a slice of strings
// It returns a new List with the received values
// The ListType is used to determine the type of list
// The boolean is used to determine if the list is multi-selectable
// The strings are the content of the items
// The function creates a new list of list items
func NewList(lt ListType, multi bool, items ...string) *List {
	numOfItems := len(items)
	list := new(List)
	list.items = GetList(lt, items...)
	list.filteredItems = list.items
	list.current = 0
	list.itemsPerPage = 10
	list.paginator = NewPaginator(BulletPaginator, numOfItems, list.itemsPerPage)
	list.multi = multi
	changeFocus(list)()
	return list
}

// List is a struct that represents a list
// The list has the following fields:
//   - items: a slice of list items
//   - filteredItems: a slice of filtered list items
//   - paginator: a paginator for the list
//   - current: the current index of the list
//   - itemsPerPage: the number of items per page
//   - multi: a boolean that indicates if the list is multi-selectable
//   - interactive: a boolean that indicates if the list is in interactive mode
//
// The List is a model that implements the Model interface
// In the Update method, the List handles the KeyMsg messages
type List struct {
	items         []*ListItem
	filteredItems []*ListItem
	paginator     *Paginator
	current       int
	itemsPerPage  int
	multi         bool
	interactive   bool
	done          bool
}

func (l *List) PrevPage() (tea.Model, tea.Cmd) {
	if l.paginator.CurrentPage() > 1 {
		l.paginator.Prev()
		l.current -= l.itemsPerPage
	}
	if l.current < 0 {
		l.current = 0
	}
	return l, changeFocus(l)
}

func (l *List) NextPage() (tea.Model, tea.Cmd) {
	if l.paginator.CurrentPage() < l.paginator.NumOfPages() {
		l.paginator.Next()
		l.current += l.itemsPerPage
	}
	if l.current >= len(l.filteredItems) {
		l.current = len(l.filteredItems) - 1
	}
	return l, changeFocus(l)
}

func (l *List) PrevItem() (tea.Model, tea.Cmd) {
	if l.current%l.itemsPerPage == 0 {
		l.paginator.Prev()
	}
	l.current--
	if l.current < 0 {
		l.current = 0
	}
	return l, changeFocus(l)
}

func (l *List) NextItem() (tea.Model, tea.Cmd) {
	if l.current%l.itemsPerPage == l.itemsPerPage-1 {
		l.paginator.Next()
	}
	l.current++
	if l.current >= len(l.filteredItems) {
		l.current = len(l.filteredItems) - 1
	}
	return l, changeFocus(l)
}

func (l *List) ManageSelection(m tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := m.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if l.multi {
				l.Done(true)
			} else {
				l.filteredItems[l.current].Check(true)
				l.Done(true)
			}
		}
		switch msg.Type {
		case tea.KeyEnter:
			if l.multi {
				l.Done(true)
			} else {
				l.filteredItems[l.current].Check(true)
				l.Done(true)
			}
		}
	}

	if l.multi {
		l.filteredItems[l.current].Update(m)
	}

	if l.Done() {
		return l, tea.Quit
	}
	return l, changeFocus(l)
}

func (l *List) Items(i ...*ListItem) []*ListItem {
	if len(i) > 0 {
		l.items = i
		l.filteredItems = i
		l.paginator.NumOfItems(len(l.items))
		changeFocus(l)()
	}
	return l.items
}

func (l *List) Filter(fo ...ListFilterOption) []*ListItem {
	l.filteredItems = l.items
	if len(fo) > 0 {
		for _, f := range fo {
			l.filteredItems = f(l.filteredItems)
		}
	}
	changeFocus(l)()
	return l.filteredItems
}

func (l *List) Current(c ...int) int {
	if len(c) > 0 {
		l.current = c[0]
		changeFocus(l)()
	}
	return l.current
}

func (l *List) ItemsPerPage(i ...int) int {
	if len(i) > 0 {
		l.itemsPerPage = i[0]
		l.paginator.ItemsPerPage(l.itemsPerPage)
		changeFocus(l)()
	}
	return l.itemsPerPage
}

func (l *List) Multi(m ...bool) bool {
	if len(m) > 0 {
		l.multi = m[0]
	}
	return l.multi
}

func (l *List) Done(d ...bool) bool {
	if len(d) > 0 {
		l.done = d[0]
		l.Interactive(false)
		l.Filter(func(list []*ListItem) []*ListItem {
			result := make([]*ListItem, 0)
			for _, item := range list {
				if item.Check() {
					result = append(result, item)
				}
			}
			return result
		})
	}
	return l.done
}

func (l *List) Interactive(i ...bool) bool {
	if len(i) > 0 {
		l.interactive = i[0]
		for _, item := range l.filteredItems {
			item.Interactive(l.interactive)
		}
		changeFocus(l)()
	}
	return l.interactive
}

func (l *List) Buf() *tui.Buffer {
	buf := tui.New()
	buf.Type(tui.ContainerBuffer)
	if !l.done {
		items := l.filteredItems[l.paginator.Start():l.paginator.End()]
		for _, item := range items {
			buf.AddC(item.Buf())
		}
		buf.AddC(l.paginator.Buf())
	} else {
		for _, item := range l.filteredItems {
			item.Type(PlainList)
			buf.AddC(tui.Muted(item.String()))
		}
	}
	return buf
}

func (l *List) String() string {
	return l.Buf().String()
}

func (l *List) Init() tea.Cmd {
	return nil
}

func (l *List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			return l.PrevItem()
		case "down":
			return l.NextItem()
		case "left":
			return l.PrevPage()
		case "right":
			return l.NextPage()
		case "enter", " ", "space":
			return l.ManageSelection(msg)
		}

		switch msg.Type {
		case tea.KeyUp:
			return l.PrevItem()
		case tea.KeyDown:
			return l.NextItem()
		case tea.KeyLeft:
			return l.PrevPage()
		case tea.KeyRight:
			return l.NextPage()
		default:
			return l.ManageSelection(msg)
		}
	}
	return l, nil
}

func (l *List) View() string {
	l.Interactive(true)
	return l.String()
}
