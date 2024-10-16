package cmd

import (
	"log"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Convert string to float, panic if conversion fails
func toFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Panic(err)
	}
	return f
}

// Render a table into a string
func renderTable(title string, rows [][]string) string {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case col == 1:
				return lipgloss.NewStyle().Width(11).PaddingRight(1).Align(lipgloss.Right)
			default:
				return lipgloss.NewStyle().Padding(0, 1)
			}
		}).
		Rows(rows...)

	return lipgloss.JoinVertical(lipgloss.Center, title, t.Render())
}
