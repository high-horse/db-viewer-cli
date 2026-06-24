package main

import (
	"db-viewer/internal/db"
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	_, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(initAppStateModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}