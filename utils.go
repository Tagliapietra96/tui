package tui

import (
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

func StrWidth(s string) int {
	var max int
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		w := runewidth.StringWidth(line)
		if w > max {
			max = w
		}
	}

	return max
}

func StrHeight(s string) int {
	return strings.Count(s, "\n") + 1
}

func GetTerminalSize() (int, int) {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return width, height
}
