package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

		switch m.step {
		case chooseLicense:
			cmd = m.handleLicenseChoice(msg)
		case enterYear:
			m, cmd = m.handleYearInput(msg)
		case enterName:
			m, cmd = m.handleNameInput(msg)
		case confirmOverwrite:
			cmd = m.handleOverwrite(msg)
		}
	}
	return m, cmd
}

func (m model) View() string {
	var s strings.Builder

	switch m.step {
	case chooseLicense:
		s.WriteString(headerStyle.Render("Choose a license:") + "\n\n")
		for i, license := range availableLicenses {
			cursor := " "
			if m.cursor == i {
				cursor = cursorStyle.Render(">")
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, license.Name))
		}
		s.WriteString(promptStyle.Render("\nUse ↑/↓ to navigate, enter to select."))

	case enterYear:
		s.WriteString(headerStyle.Render("Enter the copyright year:") + "\n\n")
		s.WriteString(m.yearInput.View())
		s.WriteString(promptStyle.Render("\nPress enter to confirm."))

	case enterName:
		s.WriteString(headerStyle.Render("Enter the copyright holder's full name:") + "\n\n")
		s.WriteString(m.nameInput.View())
		s.WriteString(promptStyle.Render("\nPress enter to confirm."))

	case confirmOverwrite:
		s.WriteString(errorStyle.Render("WARNING: 'LICENSE' file already exists."))
		s.WriteString("\n\nAre you sure you want to overwrite it? (y/n)")

	case done:
		s.WriteString(m.finalMessage)
	}

	return docStyle.Render(s.String())
}

func (m *model) handleLicenseChoice(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(availableLicenses)-1 {
			m.cursor++
		}
	case "enter":
		m.selectedLicense = availableLicenses[m.cursor]
		m.step = enterYear
		m.yearInput.Focus()
	}
	return nil
}

func (m *model) handleYearInput(msg tea.KeyMsg) (model, tea.Cmd) {
	if msg.Type == tea.KeyEnter {
		m.step = enterName
		m.yearInput.Blur()
		m.nameInput.Focus()
		return *m, nil
	}
	var cmd tea.Cmd
	m.yearInput, cmd = m.yearInput.Update(msg)
	return *m, cmd
}

func (m *model) handleNameInput(msg tea.KeyMsg) (model, tea.Cmd) {
	if msg.Type == tea.KeyEnter && m.nameInput.Value() != "" {
		m.nameInput.Blur()
		if _, err := os.Stat("LICENSE"); err == nil {
			m.step = confirmOverwrite
		} else {
			return *m, m.createLicenseFile()
		}
		return *m, nil
	}
	var cmd tea.Cmd
	m.nameInput, cmd = m.nameInput.Update(msg)
	return *m, cmd
}

func (m *model) handleOverwrite(msg tea.KeyMsg) tea.Cmd {
	switch string(msg.Runes) {
	case "y", "Y":
		return m.createLicenseFile()
	case "n", "N":
		m.step = done
		m.finalMessage = errorStyle.Render("❌ Aborted license creation.")
	}
	return nil
}

func (m *model) createLicenseFile() tea.Cmd {
	m.step = done

	templatePath := fmt.Sprintf("templates/%s.txt", m.selectedLicense.Key)
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		m.finalMessage = errorStyle.Render(fmt.Sprintf("Error reading template file: %v", err))
		return tea.Quit
	}

	licenseText := string(templateBytes)
	licenseText = strings.ReplaceAll(licenseText, "[year]", m.yearInput.Value())
	licenseText = strings.ReplaceAll(licenseText, "[fullname]", m.nameInput.Value())

	err = os.WriteFile("LICENSE", []byte(licenseText), 0644)
	if err != nil {
		m.finalMessage = errorStyle.Render(fmt.Sprintf("Error writing file: %v", err))
	} else {
		m.finalMessage = successStyle.Render(fmt.Sprintf("✅ Successfully created LICENSE with the %s.", m.selectedLicense.Name))
	}

	return tea.Quit
}
