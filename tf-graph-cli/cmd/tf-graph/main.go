package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/YaraZuhd/tf-graph/internal/graph"
	"github.com/YaraZuhd/tf-graph/internal/parser"
	"github.com/YaraZuhd/tf-graph/internal/render"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "render":
		runRender(os.Args[2:])
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`tf-graph - visualize Terraform state as a dependency graph

Usage:
  tf-graph render --state <path-to-state.json> [--out <output-name>]

Flags:
  --state   Path to a terraform show -json output file (required)
  --out     Output file name without extension (default: "graph")

Example:
  terraform show -json > state.json
  tf-graph render --state state.json --out my-infra`)
}

func runRender(args []string) {
	var statePath, outName string
	outName = "graph"

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--state":
			if i+1 >= len(args) {
				fail("missing value for --state")
			}
			statePath = args[i+1]
			i++
		case "--out":
			if i+1 >= len(args) {
				fail("missing value for --out")
			}
			outName = args[i+1]
			i++
		}
	}

	if statePath == "" {
		fail("--state is required\n\nExample:\n  tf-graph render --state state.json")
	}

	state, err := parser.LoadState(statePath)
	if err != nil {
		fail(err.Error())
	}

	g := graph.Build(state)

	if len(g.Nodes) == 0 {
		fail("no resources found in state file — nothing to render")
	}

	dotContent := render.ToDOT(g)

	dotPath := outName + ".dot"
	if err := os.WriteFile(dotPath, []byte(dotContent), 0644); err != nil {
		fail(fmt.Sprintf("writing dot file: %v", err))
	}
	fmt.Printf("Wrote %s (%d nodes, %d edges)\n", dotPath, len(g.Nodes), len(g.Edges))

	pngPath := outName + ".png"
	if err := runGraphviz(dotPath, pngPath); err != nil {
		fmt.Fprintf(os.Stderr, "\nGraphviz rendering failed: %v\n", err)
		fmt.Fprintf(os.Stderr, "The .dot file was still created — you can render it manually with:\n  dot -Tpng %s -o %s\n", dotPath, pngPath)
		os.Exit(1)
	}

	absPath, _ := filepath.Abs(pngPath)
	fmt.Printf("Wrote %s\n", absPath)
}

func runGraphviz(dotPath, pngPath string) error {
	cmd := exec.Command("dot", "-Tpng", dotPath, "-o", pngPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, string(output))
	}
	return nil
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}
