package cli

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

func Run(w io.Writer, getter func() string) {
	reason := getter()

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(1, 3).
		Margin(1, 0).
		Bold(true)

	reasonStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("81")).Bold(true)
	emojiStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)

	emoji := emojiStyle.Render("👎")
	styledReason := reasonStyle.Render(reason)
	content := fmt.Sprintf("%s  %s", emoji, styledReason)

	//nolint:errcheck
	io.WriteString(w, boxStyle.Render(content)+"\n")
}
