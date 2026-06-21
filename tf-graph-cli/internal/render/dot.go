package render

import (
	"fmt"
	"strings"

	"github.com/YaraZuhd/tf-graph/internal/graph"
)

// ToDOT renders a Graph as a Graphviz DOT-format string.
func ToDOT(g *graph.Graph) string {
	var b strings.Builder

	b.WriteString("digraph tf_graph {\n")
	b.WriteString("  rankdir=LR;\n")
	b.WriteString("  bgcolor=\"#1e1e1e\";\n")
	b.WriteString("  pad=\"0.3\";\n")
	b.WriteString("  node [shape=box, style=\"rounded,filled\", fontname=\"Helvetica\", fontcolor=\"#ffffff\", color=\"#666666\", fillcolor=\"#2d2d2d\", margin=\"0.2,0.15\"];\n")
	b.WriteString("  edge [color=\"#888888\", fontcolor=\"#cccccc\", fontname=\"Helvetica\"];\n\n")

	for _, n := range g.Nodes {
		// Build the label with a real DOT newline escape (\n) without
		// letting Go's %q re-escape the backslash. DOT label strings
		// just need to be wrapped in double quotes manually.
		label := fmt.Sprintf("%s\\n(%s)", n.Name, n.Type)
		b.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\"];\n", n.Address, label))
	}

	b.WriteString("\n")

	for _, e := range g.Edges {
		// Edge direction: dependency points FROM the resource TO what it depends on
		b.WriteString(fmt.Sprintf("  %q -> %q;\n", e.From, e.To))
	}

	b.WriteString("}\n")

	return b.String()
}
