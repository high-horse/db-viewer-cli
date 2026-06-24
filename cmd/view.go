package main

import (
	// "db-viewer/internal/helpers"
	"fmt"

	tea "charm.land/bubbletea/v2"

	lipgloss "charm.land/lipgloss/v2"
)

func (m AppStateModel) View() tea.View {

	switch m.state {

	case Auth:
		return tea.NewView(
			"Login\n\n" +
				m.email.View() + "\n" +
				m.password.View() + "\n\n" +
				textDanger.Render(m.authErrMsg) + "\n\n" +
				"Tab to switch • Enter to login",
		)

	case ConnList:
		return tea.NewView(renderConnections(m))

	case SchemaBrowser:
		return tea.NewView(renderSchema(m))

	case TableView:
		return tea.NewView(renderQuery(m))

	case ErrorState:
		return tea.NewView(renderError(m))

	default:
		return tea.NewView("loading...")
	}
}



func renderConnections(m AppStateModel) string {
	s := "Connections\n\n"

	for i, c := range m.connections {

		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s", cursor, c)

		s += lipgloss.NewStyle().
			PaddingLeft(1).
			Render(line) + "\n"
	}

	return box.Width(m.width-4).Height(m.height-4).Render(s)
}

func renderSchema(m AppStateModel) string {
	s := "Tables\n\n"

	for i, t := range m.tables {

		cursor := " "
		if i == m.tableCursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, t)
	}

	return box.Render(s)
}

func renderQuery(m AppStateModel) string {
	return box.Render(
		"SQL Editor\n\n" +
			m.queryInput +
			"\n\nResult:\n" +
			m.queryResult +
			"\n\nPress ENTER to run",
	)
}

func renderError(m AppStateModel) string {
	if m.err == nil {
		return "no error"
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Render(m.err.Error())
}