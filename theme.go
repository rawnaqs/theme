package theme

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	Gold      = "#C9933A"
	GoldLight = "#E8B86D"
	GoldDark  = "#8B6020"
	GoldDim   = "#3A2E18"

	// Backgrounds
	BgBase     = "#0f0f0f"
	BgRaised   = "#1A1A1A"
	BgElevated = "#242424"
	BgOverlay  = "#2E2E2E"

	// Borders
	BorderSubtle = "#2A2A2A"
	Border       = "#3A3A3A"
	BorderGold   = "#3A2E18"

	// Text
	TextPrimary   = "#E8B86D"
	TextSecondary = "#9A7840"
	TextTertiary  = "#5A4A28"

	// Semantic
	Success = "#4A7C59"
	Warning = "#C9933A"
	Error   = "#8B3A3A"
	Info    = "#3A6B7C"

	// Typography
	FontDisplay = "Cormorant Garamond"
	FontMono    = "IBM Plex Mono"

	// Radius
	RadiusSm = 2
	RadiusMd = 4
	RadiusLg = 8
)

var (
	Primary = lipgloss.NewStyle().Foreground(lipgloss.Color(GoldLight))
	Muted   = lipgloss.NewStyle().Foreground(lipgloss.Color(GoldDark))
	Bold    = lipgloss.NewStyle().Foreground(lipgloss.Color(GoldLight)).Bold(true)
	Inverse = lipgloss.NewStyle().Foreground(lipgloss.Color(BgBase)).Background(lipgloss.Color(Gold))

	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Success)).Bold(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(Error)).Bold(true)
	WarningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Gold)).Bold(true)

	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(GoldDim)).
		Padding(0, 1)

	PanelGold = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(Gold)).
			Padding(0, 1)
)
