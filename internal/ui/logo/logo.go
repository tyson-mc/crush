// Package logo renders the BertBud wordmark in a stylized way.
package logo

import (
	"fmt"
	"image/color"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/crush/internal/ui/styles"
	"github.com/charmbracelet/x/ansi"
)

// letterform represents a letterform. It can be stretched horizontally by
// a given amount via the boolean argument.
type letterform func(bool) string

const diag = `╱`

// brandMark sits where upstream Crush prints "Charm™".
const brandMark = "✦"

// Opts are the options for rendering the BertBud title art.
type Opts struct {
	FieldColor   color.Color // diagonal lines
	TitleColorA  color.Color // left gradient ramp point
	TitleColorB  color.Color // right gradient ramp point
	CharmColor   color.Color // brand mark (✦) color
	VersionColor color.Color // version text color
	Width        int         // width of the rendered logo, used for truncation
	Hyper        bool        // ignored: BertBud has no Hypercrush variant

	// Ignored: the BertBud wordmark never stretches a letterform.
	Unstable bool
}

// Render renders the BertBud logo. Set the argument to true to render the narrow
// version, intended for use in a sidebar.
//
// The compact argument determines whether it renders compact for the sidebar
// or wider for the main pane.
func Render(base lipgloss.Style, version string, compact bool, o Opts) string {
	charm := brandMark

	fg := func(c color.Color, s string) string {
		return lipgloss.NewStyle().Foreground(c).Render(s)
	}

	// Title. The narrow version stacks BERT over BUD to fit the sidebar.
	const spacing = 1
	bertLetterforms := []letterform{LetterB, LetterE, LetterR, LetterT}
	budLetterforms := []letterform{LetterB, LetterU, LetterD}
	crushLetterforms := append(append([]letterform{}, bertLetterforms...), budLetterforms...)

	// Upstream stretches one random letterform; BERTBUD reads better without.
	crush := renderWord(spacing, -1, crushLetterforms...)
	if compact {
		crush = renderWord(spacing, -1, bertLetterforms...) + "\n" + renderWord(spacing, -1, budLetterforms...)
	}
	crushWidth := lipgloss.Width(crush)
	b := new(strings.Builder)
	for r := range strings.SplitSeq(crush, "\n") {
		fmt.Fprintln(b, styles.ApplyForegroundGrad(base, r, o.TitleColorA, o.TitleColorB))
	}
	crush = b.String()

	// Charm and version.
	metaRowGap := 1
	maxVersionWidth := crushWidth - lipgloss.Width(charm) - metaRowGap
	version = ansi.Truncate(version, maxVersionWidth, "…") // truncate version if too long.
	gap := max(0, crushWidth-lipgloss.Width(charm)-lipgloss.Width(version))
	metaRow := fg(o.CharmColor, charm) + strings.Repeat(" ", gap) + fg(o.VersionColor, version)

	// Join the meta row and big BertBud title.
	crush = strings.TrimSpace(metaRow + "\n" + crush)

	// Narrow version, stacked.
	if compact {
		field := fg(o.FieldColor, strings.Repeat(diag, crushWidth))
		return strings.Join([]string{field, field, crush, field, ""}, "\n")
	}

	fieldHeight := lipgloss.Height(crush)

	// Left field.
	const leftWidth = 6
	leftFieldRow := fg(o.FieldColor, strings.Repeat(diag, leftWidth))
	leftField := new(strings.Builder)
	for range fieldHeight {
		fmt.Fprintln(leftField, leftFieldRow)
	}

	// Right field.
	rightWidth := max(15, o.Width-crushWidth-leftWidth-2) // 2 for the gap.
	const stepDownAt = 0
	rightField := new(strings.Builder)
	for i := range fieldHeight {
		width := rightWidth
		if i >= stepDownAt {
			width = rightWidth - (i - stepDownAt)
		}
		fmt.Fprint(rightField, fg(o.FieldColor, strings.Repeat(diag, width)), "\n")
	}

	// Return the wide version.
	const hGap = " "
	logo := lipgloss.JoinHorizontal(lipgloss.Top, leftField.String(), hGap, crush, hGap, rightField.String())
	if o.Width > 0 {
		// Truncate the logo to the specified width.
		lines := strings.Split(logo, "\n")
		for i, line := range lines {
			lines[i] = ansi.Truncate(line, o.Width, "")
		}
		logo = strings.Join(lines, "\n")
	}
	return logo
}

// SmallRender renders a smaller version of the BertBud logo, suitable for
// smaller windows or sidebar usage.
func SmallRender(t *styles.Styles, width int, o Opts) string {
	name := "BertBud"
	charm := brandMark
	title := t.Logo.SmallCharm.Render(charm)
	title = fmt.Sprintf("%s %s", title, styles.ApplyBoldForegroundGrad(t.Logo.GradCanvas, name, t.Logo.SmallGradFromColor, t.Logo.SmallGradToColor))
	remainingWidth := width - lipgloss.Width(title) - 1 // 1 for the space after the name
	if remainingWidth > 0 {
		lines := strings.Repeat("╱", remainingWidth)
		title = fmt.Sprintf("%s %s", title, t.Logo.SmallDiagonals.Render(lines))
	}
	return title
}
