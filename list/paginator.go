package list

import (
	"fmt"
	"math"

	"github.com/Tagliapietra96/tui"
	tea "github.com/charmbracelet/bubbletea"
)

// PaginatorType is an enumeration of the types of paginators
// It has the following values:
//   - BulletPaginator: a paginator with bullet points
//   - NumberPaginator: a paginator with numbers
type PaginatorType int

const (
	BulletPaginator PaginatorType = iota // BulletPaginator is a paginator with bullet points
	NumberPaginator                      // NumberPaginator is a paginator with numbers
)

// NewPaginator creates a new Paginator
// It receives a PaginatorType, the number of items, and the number of items per page
// It returns a new Paginator with the received values
// The PaginatorType is used to determine the type of paginator
// The number of items is the total number of items
// The number of items per page is the number of items per page
// If the number of items per page is not provided, it will be set to 10
func NewPaginator(pt PaginatorType, numOfItems int, itemsPerPage ...int) *Paginator {
	p := new(Paginator)
	p.paginatorType = pt
	p.numOfItems = numOfItems
	p.currentPage = 1
	if len(itemsPerPage) > 0 {
		p.itemsPerPage = itemsPerPage[0]
	} else {
		p.itemsPerPage = 10
	}
	p.Recalc()
	return p
}

// Paginator is a struct that represents a paginator
// The paginator has the following fields:
//   - paginatorType: the type of paginator (BulletPaginator, NumberPaginator)
//   - itemsPerPage: the number of items per page
//   - currentPage: the current page
//   - numOfPages: the total number of pages
//   - numOfItems: the total number of items
//   - start: the start index of the items
//   - end: the end index of the items
//
// The Paginator is a model that implements the Model interface
// In the Update method, the Paginator handles the KeyMsg messages
// If the message is a right key, the Paginator goes to the next page
// If the message is a left key, the Paginator goes to the previous page
type Paginator struct {
	paginatorType PaginatorType
	itemsPerPage  int
	currentPage   int
	numOfPages    int
	numOfItems    int
	start         int
	end           int
}

// Next goes to the next page
// It increments the current page by one
// It recalculates the number of pages, the start index, and the end index
// It returns the updated model and no command
func (p *Paginator) Next() (tea.Model, tea.Cmd) {
	if p.currentPage < p.numOfPages {
		p.currentPage++
		p.Recalc()
	}
	return p, nil
}

// Prev goes to the previous page
// It decrements the current page by one
// It recalculates the number of pages, the start index, and the end index
// It returns the updated model and no command
func (p *Paginator) Prev() (tea.Model, tea.Cmd) {
	if p.currentPage > 1 {
		p.currentPage--
		p.Recalc()
	}
	return p, nil
}

// Recalc recalculates the number of pages, the start index, and the end index
// It calculates the number of pages by dividing the total number of items by the number of items per page
// It calculates the start index by multiplying the current page minus one by the number of items per page
// It calculates the end index by multiplying the current page by the number of items per page
func (p *Paginator) Recalc() {
	p.numOfPages = int(math.Ceil(float64(p.numOfItems) / float64(p.itemsPerPage)))
	p.start = (p.currentPage - 1) * p.itemsPerPage
	p.end = p.currentPage * p.itemsPerPage
}

// Type sets or returns the type of the paginator
// It receives a PaginatorType as input
// If the PaginatorType is provided, it sets the paginatorType field to the received value
// It returns the paginatorType field
func (p *Paginator) Type(pt ...PaginatorType) PaginatorType {
	if len(pt) > 0 {
		p.paginatorType = pt[0]
	}
	return p.paginatorType
}

// ItemsPerPage sets or returns the number of items per page
// It receives an integer as input
// If the integer is provided, it sets the itemsPerPage field to the received value
// It recalculates the number of pages, the start index, and the end index
// It returns the itemsPerPage field
func (p *Paginator) ItemsPerPage(ipp ...int) int {
	if len(ipp) > 0 {
		p.itemsPerPage = ipp[0]
		p.Recalc()
	}
	return p.itemsPerPage
}

// CurrentPage sets or returns the current page
// It receives an integer as input
// If the integer is provided, it sets the currentPage field to the received value
// It recalculates the number of pages, the start index, and the end index
// It returns the currentPage field
func (p *Paginator) CurrentPage(cp ...int) int {
	if len(cp) > 0 {
		p.currentPage = cp[0]
		p.Recalc()
	}
	return p.currentPage
}

// NumOfPages returns the total number of pages
// It recalculates the number of pages
// It returns the numOfPages field
func (p *Paginator) NumOfPages() int {
	p.Recalc()
	return p.numOfPages
}

// NumOfItems sets or returns the total number of items
// It receives an integer as input
// If the integer is provided, it sets the numOfItems field to the received value
// It recalculates the number of pages, the start index, and the end index
// It returns the numOfItems field
func (p *Paginator) NumOfItems(noi ...int) int {
	if len(noi) > 0 {
		p.numOfItems = noi[0]
		p.Recalc()
	}
	return p.numOfItems
}

// Start returns the start index of the items
// It recalculates the start index
// It returns the start field
func (p *Paginator) Start() int {
	p.Recalc()
	return p.start
}

// End returns the end index of the items
// It recalculates the end index
// It returns the end field
func (p *Paginator) End() int {
	p.Recalc()
	return p.end
}

// Buf returns a buffer with the paginator
// It creates a new buffer
// It recalculates the number of pages
// It adds bullet points or numbers to the buffer based on the type of paginator
// It returns the buffer
func (p *Paginator) Buf() *tui.Buffer {
	buf := tui.New()
	p.Recalc()
	for i := 1; i <= p.numOfPages; i++ {
		switch p.paginatorType {
		case BulletPaginator:
			if i == p.currentPage {
				buf.AddC(tui.H5("•"))
			} else {
				buf.AddC(tui.Muted("•"))
			}
		default: // NumberPaginator
			if i == p.currentPage {
				if i != p.numOfPages {
					buf.AddC(tui.H5(fmt.Sprintf("%d ", i)))
				} else {
					buf.AddC(tui.H5(fmt.Sprintf("%d", i)))
				}
			} else {
				if i != p.numOfPages {
					buf.AddC(tui.Muted(fmt.Sprintf("%d ", i)))
				} else {
					buf.AddC(tui.Muted(fmt.Sprintf("%d", i)))
				}
			}
		}
	}
	return buf
}

// String returns the string representation of the paginator
// It returns the string representation of the buffer of the paginator
func (p *Paginator) String() string {
	return p.Buf().String()
}

// Init initializes the paginator model
func (p *Paginator) Init() tea.Cmd {
	return nil
}

// Update updates the paginator model
// It handles the KeyMsg messages
// If the message is a right key, it goes to the next page
// If the message is a left key, it goes to the previous page
func (p *Paginator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	p.Recalc()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRight:
			return p.Next()
		case tea.KeyLeft:
			return p.Prev()
		}
	}
	return p, nil
}

// View returns the view of the paginator model
func (p *Paginator) View() string {
	return p.String()
}
