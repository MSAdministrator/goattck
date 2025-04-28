// Contains utilities for the menu compoonent
package menu

import (
	"image/color"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	lipglosstable "github.com/charmbracelet/lipgloss/table"
	"github.com/lucasb-eyer/go-colorful"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Foreground(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}

func getFormattedTable(headers []string, rows [][]string) string {
	return lipglosstable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).Align(lipgloss.Center)
			default:
				return lipgloss.NewStyle().Padding(0, 1).Width(width / len(headers)).Foreground(lipgloss.Color("245"))
			}
		}).
		Headers(headers...).
		Rows(rows...).
		String()
}
