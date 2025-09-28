package core

import "github.com/charmbracelet/lipgloss"

var (
	appStyle              = lipgloss.NewStyle().Margin(1, 2)
	headerStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true).Padding(0, 0, 1, 0)
	questionStyle         = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	answeredQuestionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
	userAnswerStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("57"))
	promptStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	selectedItemStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	successStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("40")).Bold(true)
	errorStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	resultBoxStyle        = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("40")).
				Padding(1, 2)
)

type licenseInfo struct {
	Key  string
	Name string
}

var availableLicenses = []licenseInfo{
	{Key: "afl-3.0", Name: "Academic Free License v3.0"},
	{Key: "apache-2.0", Name: "Apache License 2.0"},
	{Key: "artistic-2.0", Name: "Artistic License 2.0"},
	{Key: "bsl-1.0", Name: "Boost Software License 1.0"},
	{Key: "bsd-2-clause", Name: "BSD 2-Clause \"Simplified\" License"},
	{Key: "bsd-3-clause", Name: "BSD 3-clause \"New\" or \"Revised\" License"},
	{Key: "bsd-3-clause-clear", Name: "BSD 3-Clause Clear License"},
	{Key: "bsd-4-clause", Name: "BSD 4-clause \"Original\" or \"Old\" License"},
	{Key: "bsd-zero-clause", Name: "BSD Zero-Clause License"},
	{Key: "cc-by-4.0", Name: "Creative Commons Attribution 4.0 International"},
	{Key: "cc-by-sa-4.0", Name: "Creative Commons Attribution-ShareAlike 4.0 International"},
	{Key: "cc-by-nc-4.0", Name: "Creative Commons Attribution-NonCommercial 4.0 International"},
	{Key: "cc0-1.0", Name: "Creative Commons Zero v1.0 Universal"},
	{Key: "wtfpl", Name: "Do What The F*ck You Want To Public License"},
	{Key: "ecl-2.0", Name: "Educational Community License v2.0"},
	{Key: "epl-1.0", Name: "Eclipse Public License 1.0"},
	{Key: "epl-2.0", Name: "Eclipse Public License 2.0"},
	{Key: "eupl-1.1", Name: "European Union Public License 1.1"},
	{Key: "agpl-3.0", Name: "GNU Affero General Public License v3.0"},
	{Key: "gpl-3.0", Name: "GNU General Public License v3.0"},
	{Key: "gpl-2.0", Name: "GNU General Public License v2.0"},
	{Key: "lgpl-3.0", Name: "GNU Lesser General Public License v3.0"},
	{Key: "mit", Name: "MIT License"},
}
