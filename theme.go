package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/*
This file is for CLIs only — respects the user's terminal theme via ANSI colors.
The hex constants below are for brand identity (CSS, web, Python) — never use them in styles here.
All styles use ANSI color codes so Catppuccin, Dracula, Nord, Gruvbox etc. all look correct.
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

// ── Base styles ───────────────────────────────────────────────────────────────
var (
	Primary = lipgloss.NewStyle().Foreground(lipgloss.Color("15")) // bright white
	Muted   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))  // white
	Dim     = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))  // bright black
	Bold    = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Bold(true)
	Italic  = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Italic(true)
	Inverse = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("15"))
)

// ── Status styles ─────────────────────────────────────────────────────────────
var (
	SuccessStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)   // green
	ErrorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)   // red
	WarningStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true)   // yellow
	InfoStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))              // cyan
	ProcessingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Italic(true) // yellow italic
)

// ── Panel styles ──────────────────────────────────────────────────────────────
var (
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Padding(0, 1)

	PanelAccent = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("3")).
			Padding(0, 1)
)

// ── Tag styles ────────────────────────────────────────────────────────────────
var (
	// Tag — yellow, used for text note tags
	Tag = lipgloss.NewStyle().
		Foreground(lipgloss.Color("3")).
		Background(lipgloss.Color("0")).
		PaddingLeft(1).PaddingRight(1)

	// TagArticle — blue, used for article note tags
	TagArticle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4")).
			Background(lipgloss.Color("0")).
			PaddingLeft(1).PaddingRight(1)

	// TagImage — magenta, used for image note tags
	TagImage = lipgloss.NewStyle().
			Foreground(lipgloss.Color("5")).
			Background(lipgloss.Color("0")).
			PaddingLeft(1).PaddingRight(1)

	// TagMuted — dim, used for secondary/metadata tags
	TagMuted = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Background(lipgloss.Color("0")).
			PaddingLeft(1).PaddingRight(1)
)

// ── Type badge styles ─────────────────────────────────────────────────────────
var (
	// TypeText — green badge for text captures
	TypeText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Background(lipgloss.Color("0")).
			Bold(true).
			PaddingLeft(1).PaddingRight(1)

	// TypeArticle — blue badge for article captures
	TypeArticle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4")).
			Background(lipgloss.Color("0")).
			Bold(true).
			PaddingLeft(1).PaddingRight(1)

	// TypeImage — magenta badge for image captures
	TypeImage = lipgloss.NewStyle().
			Foreground(lipgloss.Color("5")).
			Background(lipgloss.Color("0")).
			Bold(true).
			PaddingLeft(1).PaddingRight(1)
)

// ── Search result styles ──────────────────────────────────────────────────────
var (
	// SearchTitle — bright, bold — most prominent element per result
	SearchTitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true)

	// SearchScore — dim, right-aligned — present but not dominant
	SearchScore = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))

	// SearchDate — muted
	SearchDate = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))

	// SearchExcerpt — italic, muted, with left border
	SearchExcerpt = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")).
			Italic(true).
			BorderLeft(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("8")).
			PaddingLeft(1)
)

// ── Helper functions ──────────────────────────────────────────────────────────

// CaptureOK renders: ✓ saved · #tag1 #tag2 · 3ms
func CaptureOK(action, tags, duration string) string {
	return SuccessStyle.Render("✓") + " " +
		Primary.Render(action) + " " +
		Muted.Render("· "+tags) + " " +
		Dim.Render("· "+duration)
}

// CaptureQueued renders: ⏳ queued · type · id: xyz
func CaptureQueued(jobType, id string) string {
	return ProcessingStyle.Render("⏳") + " " +
		Muted.Render("queued · "+jobType) + " " +
		Dim.Render("· id: "+id)
}

// Divider renders a full-width horizontal rule
func Divider(width int) string {
	var b strings.Builder
	for range width {
		b.WriteString("─")
	}
	return Dim.Render(b.String())
}

// RenderTag renders a tag badge — type determines color
// noteType: "text" | "article" | "image" | "" (muted)
func RenderTag(tag, noteType string) string {
	switch noteType {
	case "article":
		return TagArticle.Render(tag)
	case "image":
		return TagImage.Render(tag)
	case "text":
		return Tag.Render(tag)
	default:
		return TagMuted.Render(tag)
	}
}

// RenderTypeBadge renders a capture type badge
func RenderTypeBadge(noteType string) string {
	switch noteType {
	case "text":
		return TypeText.Render(noteType)
	case "article":
		return TypeArticle.Render(noteType)
	case "image":
		return TypeImage.Render(noteType)
	default:
		return TagMuted.Render(noteType)
	}
}
