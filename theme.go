package theme

import "strings"

import "github.com/charmbracelet/lipgloss"

const (
	Gold      = "#C9933A"
	GoldLight = "#E8B86D"
	GoldDark  = "#8B6020"
	GoldDim   = "#3A2E18"

	BgBase     = "#0f0f0f"
	BgRaised   = "#1A1A1A"
	BgElevated = "#242424"
	BgOverlay  = "#2E2E2E"

	BorderSubtle = "#2A2A2A"
	Border       = "#3A3A3A"
	BorderGold   = "#3A2E18"

	TextPrimary   = "#E8B86D"
	TextSecondary = "#8A6B35"
	TextTertiary  = "#6A5530"

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
	ColorGold = lipgloss.AdaptiveColor{
		Dark:  "#C9933A",
		Light: "#7A5010",
	}
	ColorGoldLight = lipgloss.AdaptiveColor{
		Dark:  "#E8B86D",
		Light: "#5A3A10",
	}
	ColorGoldDark = lipgloss.AdaptiveColor{
		Dark:  "#8B6020",
		Light: "#8B6020",
	}
	ColorGoldDim = lipgloss.AdaptiveColor{
		Dark:  "#3A2E18",
		Light: "#C0A060",
	}
	ColorSuccess = lipgloss.AdaptiveColor{
		Dark:  "#4A7C59",
		Light: "#2A5C39",
	}
	ColorError = lipgloss.AdaptiveColor{
		Dark:  "#8B3A3A",
		Light: "#6B1A1A",
	}
	ColorWarning = lipgloss.AdaptiveColor{
		Dark:  "#C9933A",
		Light: "#7A5010",
	}
	ColorInfo = lipgloss.AdaptiveColor{
		Dark:  "#3A6B7C",
		Light: "#1A4B5C",
	}
	ColorBg = lipgloss.AdaptiveColor{
		Dark:  "#0f0f0f",
		Light: "#FFFFFF",
	}
	ColorBgRaised = lipgloss.AdaptiveColor{
		Dark:  "#1A1A1A",
		Light: "#F5F0E8",
	}
	ColorBorder = lipgloss.AdaptiveColor{
		Dark:  "#2A2A2A",
		Light: "#D0C8B8",
	}
	ColorBorderGold = lipgloss.AdaptiveColor{
		Dark:  "#3A2E18",
		Light: "#B09060",
	}
)

var (
	Primary = lipgloss.NewStyle().Foreground(ColorGoldLight)
	Muted   = lipgloss.NewStyle().Foreground(ColorGoldDark)
	Dim     = lipgloss.NewStyle().Foreground(ColorGoldDim)
	Bold    = lipgloss.NewStyle().Foreground(ColorGoldLight).Bold(true)
	Italic  = lipgloss.NewStyle().Foreground(ColorGoldDark).Italic(true)
	Inverse = lipgloss.NewStyle().
		Foreground(ColorBg).
		Background(ColorGold)

	SuccessStyle    = lipgloss.NewStyle().Foreground(ColorSuccess).Bold(true)
	ErrorStyle      = lipgloss.NewStyle().Foreground(ColorError).Bold(true)
	WarningStyle    = lipgloss.NewStyle().Foreground(ColorWarning).Bold(true)
	InfoStyle       = lipgloss.NewStyle().Foreground(ColorInfo)
	ProcessingStyle = lipgloss.NewStyle().Foreground(ColorGoldDark).Italic(true)

	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorGoldDim).
		Padding(0, 1)

	PanelGold = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorGold).
			Padding(0, 1)

	Tag = lipgloss.NewStyle().
		Foreground(ColorBg).
		Background(ColorGold).
		PaddingLeft(1).PaddingRight(1)

	TagMuted = lipgloss.NewStyle().
			Foreground(ColorGoldLight).
			Background(ColorBgRaised).
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
