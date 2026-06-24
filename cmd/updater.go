package main

import (
	tea "charm.land/bubbletea/v2"
)

func (m AppStateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}

	switch m.state {
	case Auth:
		{
			var cmd tea.Cmd
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "tab":
					if m.email.Focused() {
						m.email.Blur()
						m.password.Focus()
					} else {
						m.password.Blur()
						m.email.Focus()
					}

				case "enter":
					ok, msg := validatePassword(m.email.Value(), m.password.Value())
					if !ok {
						m.authErrMsg = msg
					} else {
						m.state = ConnList
					}
				}

			}

			m.email, cmd = m.email.Update(msg)
			m.password, _ = m.password.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func validatePassword(uname, passwd string) (bool, string) {
	if len(uname) == 0 || len(passwd) == 0 {
		return false, "username and password cannot be empty"
	}

	if uname != "root" || passwd != "root" {
		return false, "invalid credetials"
	}

	return true, ""
}
