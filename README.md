# rawnaqs/theme

Single source of truth for design tokens.

## Files

| File | Purpose |
|------|---------|
| `tokens/tokens.json` | Source of truth |
| `theme.go` | Go constants + lipgloss styles |
| `theme.css` | Web tokens (shadcn compatible) |

## Go Usage

```go
import "github.com/rawnaqs/theme"

// Raw colors
style := lipgloss.NewStyle().
    Foreground(lipgloss.Color(theme.Gold)).
    Bold(true)

// Pre-built styles
fmt.Println(theme.Primary.Render("Hello"))
fmt.Println(theme.SuccessStyle.Render("✓ Success"))
fmt.Println(theme.PanelGold.Render("Card content"))
```

## Web Usage (shadcn)

Copy `theme.css` to your project's globals.css, or import directly:

```css
@import "./path/to/theme.css";
```

Then use shadcn semantic classes:

```tsx
<div className="bg-background text-foreground">
  <button className="bg-primary text-primary-foreground">Click</button>
  <span className="text-muted-foreground">Muted text</span>
</div>
```

## Updating

Edit `tokens/tokens.json`, then manually sync to:
- `theme.go` (Go constants)
- `css.css` (CSS variables)
