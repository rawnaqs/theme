package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/*
Avoid creating this file with tokens.json those are for brand identiy and this file is for clis and respects the users terminal themes
*/

const (
	Gold      = "#C9933A"
	GoldLight = "#E8B86D"
	GoldDark  = "#A67830"
	GoldDim   = "#6B5530"

	BgBase     = "#0f0f0f"
	BgRaised   = "#1A1A1A"
	BgElevated = "#242424"
	BgOverlay  = "#2E2E2E"

	BorderSubtle = "#2A2A2A"
	Border       = "#3A3A3A"
	BorderGold   = "#5A4520"

	TextPrimary   = "#E8B86D"
	TextSecondary = "#A67830"
	TextTertiary  = "#6B5530"

	Success = "#4A7C59"
	Warning = "#C9933A"
	Error   = "#8B3A3A"
	Info    = "#3A6B7C"

	FontDisplay = "Cormorant Garamond"
	FontMono    = "IBM Plex Mono"

	RadiusSm = 2
	RadiusMd = 4
	RadiusLg = 8
)

var (
	Primary = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	Muted   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	Dim     = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	Bold    = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Bold(true)
	Italic  = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Italic(true)
	Inverse = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("15"))

	SuccessStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	ErrorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	WarningStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true)
	InfoStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	ProcessingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Italic(true)

	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Padding(0, 1)

	PanelGold = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("3")).
			Padding(0, 1)

	Tag = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("3")).
		PaddingLeft(1).PaddingRight(1)

	TagMuted = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Background(lipgloss.Color("0")).
			PaddingLeft(1).PaddingRight(1)
)

func CaptureOK(action, tags, duration string) string {
	return SuccessStyle.Render("✓") + " " +
		Primary.Render(action) + " " +
		Muted.Render("· "+tags) + " " +
		Dim.Render("· "+duration)
}

func CaptureQueued(jobType, id string) string {
	return ProcessingStyle.Render("⏳") + " " +
		Muted.Render("queued · "+jobType) + " " +
		Dim.Render("· id: "+id)
}

func Divider(width int) string {
	var line strings.Builder
	for range width {
		line.WriteString("─")
	}
	return Dim.Render(line.String())
}

func RenderTag(tag string) string {
	return Tag.Render(tag)
}
