package core

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type step int

const (
	chooseLicense step = iota
	enterYear
	enterName
	confirmOverwrite
	done
)

type model struct {
	step            step
	cursor          int
	selectedLicense licenseInfo
	yearInput       textinput.Model
	nameInput       textinput.Model
	history         []string
	finalMessage    string
}

func InitialModel() model {
	tiYear := textinput.New()
	tiYear.Prompt = "▶ "
	tiYear.Placeholder = "Year"
	tiYear.CharLimit = 4
	tiYear.Width = 10
	tiYear.SetValue(fmt.Sprintf("%d", time.Now().Year()))

	tiName := textinput.New()
	tiName.Prompt = "▶ "
	tiName.Placeholder = "Your Name"
	tiName.CharLimit = 156
	tiName.Width = 50

	return model{
		step:      chooseLicense,
		history:   []string{},
		yearInput: tiYear,
		nameInput: tiName,
	}
}
