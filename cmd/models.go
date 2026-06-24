package main

import (
	"db-viewer/internal/types"

	textinput "charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type State int

const (
	// connection layer
	Auth State = iota
	ConnList
	ConnForm
	Connecting

	// browse layer
	SchemaBrowser
	TableView

	// query layer
	QueryEditor
	QueryResult

	// system state
	ErrorState
	Quit
)

type AppStateModel struct {
	state State

	// auth
	email    textinput.Model
	password textinput.Model
	authErrMsg string
	focus    int // 0=email, 1=password

	connections  []types.Connection
	selectedConn int
	currentConn  *types.Connection

	tables        []string
	tableCursor   int
	selectedTable string

	queryInput  string
	queryResult string

	cursor  int
	err     error
	loading bool
	width   int
	height  int
}

func initAppStateModel() AppStateModel {
	email := textinput.New()
	email.Placeholder = "Usernname"
	email.Focus()
	email.SetWidth(30)

	password := textinput.New()
	password.Placeholder = "Password"
	password.EchoMode = textinput.EchoPassword
	password.EchoCharacter = '•'
	password.SetWidth(30)
	return AppStateModel{
		state: Auth,
		email: email,
		password: password,
	}
}

func (m AppStateModel) Init() tea.Cmd {
	return nil
}
