package core

import "github.com/charmbracelet/lipgloss"

var (
	docStyle     = lipgloss.NewStyle().Margin(1, 2)
	headerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true)
	promptStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("40")).Bold(true)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
)

type licenseInfo struct {
	Key  string
	Name string
}

var availableLicenses = []licenseInfo{
	{Key: "mit", Name: "MIT License"},
}
