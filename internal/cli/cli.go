package cli

import (
	// "fmt"

	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"golang.org/x/term"
)

func Run(w io.Writer, getter func() string) {
	reason := getter()

	// Default width if not a terminal
	width := 60

	// Try to detect terminal width if output is a terminal
	if f, ok := w.(*os.File); ok && isatty.IsTerminal(f.Fd()) {
		if tw, _, err := term.GetSize(int(f.Fd())); err == nil && tw > 0 {
			width = tw
		}
	}

	// Leave a bit more margin for border, padding, and emoji spacing
	minContentWidth := 32
	maxContentWidth := width - 10
	if maxContentWidth < minContentWidth {
		maxContentWidth = minContentWidth
	}

	// Calculate the actual content width based on the reason length
	reasonLen := lipgloss.Width(reason)
	contentWidth := maxContentWidth

	if reasonLen < maxContentWidth {
		contentWidth = reasonLen
	}

	emojiStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("212")).
		Bold(true)

	emojiBlock := emojiStyle.Render("👎")
	emojiWidth := lipgloss.Width(emojiBlock)

	reasonStyle := lipgloss.NewStyle().
		Width(contentWidth).
		Foreground(lipgloss.Color("81")).
		Bold(true)

	// Wrap the reason text first, then apply color to each line
	wrappedReason := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Left).
		Render(reason)

	// Split into lines and color each line, avoiding extra newlines
	var coloredLines []string

	for _, line := range strings.Split(wrappedReason, "\n") {
		if len(line) > 0 {
			coloredLines = append(coloredLines, reasonStyle.Render(line))
		}
	}

	coloredReason := lipgloss.JoinVertical(lipgloss.Left, coloredLines...)

	reasonBlock := lipgloss.NewStyle().Width(contentWidth).Align(lipgloss.Left).Render(coloredReason)
	content := lipgloss.JoinHorizontal(lipgloss.Top, emojiBlock, reasonBlock)

	boxStyle := lipgloss.NewStyle().
		Width(contentWidth+emojiWidth+3).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(1).
		Margin(1, 0).
		Bold(true)

	//nolint:errcheck
	io.WriteString(w, boxStyle.Render(content)+"\n")
}
