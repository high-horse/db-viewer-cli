package main

import (
	lipgloss "charm.land/lipgloss/v2"
)


var box = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(1, 2).
	// Width(40).
	Align(lipgloss.Center)


var textDanger = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))

