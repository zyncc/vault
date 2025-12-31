package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/zyncc/vault/db"
)

var (
	cellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	tableStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("238"))
)

func renderTable(rows []db.PasswordStore) string {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			return cellStyle
		}).
		Headers(
			"DOMAIN",
			"EMAIL",
			"PASSWORD",
			"CREATED AT",
			"UPDATED AT",
		)

	for _, r := range rows {
		year, month, day := r.CreatedAt.Date()
		humanReadableCreatedAt := fmt.Sprintf("%v %v %v", day, month.String(), year)

		uyear, umonth, uday := r.UpdatedAt.Date()
		humanReadableUpdatedAt := fmt.Sprintf("%v %v %v", uday, umonth.String(), uyear)
		t.Row(
			r.Domain,
			r.Email,
			r.Password,
			humanReadableCreatedAt,
			humanReadableUpdatedAt,
		)
	}

	return tableStyle.Render(t.String())
}
