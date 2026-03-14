package graph

import (
	"fmt"
	"strings"

	"github.com/herfstvalt/ossref/internal/refs"
)

const (
	nodeWidth    = 200
	nodeHeight   = 40
	nodeRadius   = 8
	padding      = 60
	centerY      = 80
	refStartY    = 200
	refSpacingY  = 70
	labelOffsetY = 25
)

// colors
const (
	rootFill   = "#4338CA"
	rootText   = "#FFFFFF"
	refFill    = "#F3F4F6"
	refText    = "#1F2937"
	refStroke  = "#D1D5DB"
	lineStroke = "#9CA3AF"
	learnedCol = "#6B7280"
	bgColor    = "#FFFFFF"
)

// RenderSVG generates an SVG dependency graph.
func RenderSVG(projectName string, f *refs.File) string {
	// Group by project
	grouped := map[string][]refs.Reference{}
	order := []string{}
	for _, r := range f.References {
		if _, exists := grouped[r.Project]; !exists {
			order = append(order, r.Project)
		}
		grouped[r.Project] = append(grouped[r.Project], r)
	}

	count := len(order)
	if count == 0 {
		return ""
	}

	totalWidth := count*nodeWidth + (count-1)*padding
	svgWidth := totalWidth + padding*2
	svgHeight := refStartY + count*refSpacingY + padding

	// Calculate max learned lines per project for height
	maxEntries := 0
	for _, entries := range grouped {
		if len(entries) > maxEntries {
			maxEntries = len(entries)
		}
	}
	if maxEntries > 1 {
		svgHeight += maxEntries * 20
	}

	var b strings.Builder

	b.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" width="%d" height="%d">`, svgWidth, svgHeight, svgWidth, svgHeight))
	b.WriteString("\n")

	// Background
	b.WriteString(fmt.Sprintf(`  <rect width="%d" height="%d" fill="%s"/>`, svgWidth, svgHeight, bgColor))
	b.WriteString("\n")

	// Styles
	b.WriteString(`  <style>
    .root-text { font: bold 14px -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; fill: ` + rootText + `; }
    .ref-text { font: 13px -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; fill: ` + refText + `; }
    .learned-text { font: 11px -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; fill: ` + learnedCol + `; }
  </style>`)
	b.WriteString("\n")

	// Root node (centered at top)
	rootX := svgWidth/2 - nodeWidth/2
	b.WriteString(fmt.Sprintf(`  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s"/>`,
		rootX, centerY, nodeWidth, nodeHeight, nodeRadius, rootFill))
	b.WriteString("\n")

	rootName := shortName(projectName)
	b.WriteString(fmt.Sprintf(`  <text x="%d" y="%d" text-anchor="middle" class="root-text">%s</text>`,
		svgWidth/2, centerY+labelOffsetY, escape(rootName)))
	b.WriteString("\n")

	rootCenterX := svgWidth / 2
	rootBottomY := centerY + nodeHeight

	// Reference nodes
	for i, project := range order {
		entries := grouped[project]
		refX := padding + i*(nodeWidth+padding)
		refCenterX := refX + nodeWidth/2
		refTopY := refStartY

		// Connection line
		b.WriteString(fmt.Sprintf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="2"/>`,
			rootCenterX, rootBottomY, refCenterX, refTopY, lineStroke))
		b.WriteString("\n")

		// Node box — taller if multiple entries
		boxHeight := nodeHeight
		if len(entries) > 1 {
			boxHeight += (len(entries) - 1) * 18
		}

		b.WriteString(fmt.Sprintf(`  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s" stroke="%s" stroke-width="1"/>`,
			refX, refTopY, nodeWidth, boxHeight, nodeRadius, refFill, refStroke))
		b.WriteString("\n")

		name := shortName(project)
		b.WriteString(fmt.Sprintf(`  <text x="%d" y="%d" text-anchor="middle" class="ref-text">%s</text>`,
			refCenterX, refTopY+18, escape(name)))
		b.WriteString("\n")

		// Learned entries below the name
		for j, e := range entries {
			learned := truncate(e.Learned, 28)
			ly := refTopY + 34 + j*16
			b.WriteString(fmt.Sprintf(`  <text x="%d" y="%d" text-anchor="middle" class="learned-text">%s</text>`,
				refCenterX, ly, escape(learned)))
			b.WriteString("\n")
		}
	}

	b.WriteString("</svg>\n")
	return b.String()
}

func shortName(project string) string {
	parts := strings.Split(project, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return project
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-1] + "..."
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}
