package graph

import "github.com/YaraZuhd/tf-graph/internal/parser"

// Node represents a single Terraform resource in the graph.
type Node struct {
	Address string // full address, e.g. "cloudflare_workers_kv_namespace.tf_graph_kv"
	Type    string // resource type, e.g. "cloudflare_workers_kv_namespace"
	Name    string // resource name, e.g. "tf_graph_kv"
}

// Edge represents a dependency: From depends on To.
type Edge struct {
	From string // address of the dependent resource
	To   string // address of the resource it depends on
}

// Graph is the full set of nodes and edges built from a Terraform state.
type Graph struct {
	Nodes []Node
	Edges []Edge
}

// Build converts a parsed TerraformState into a Graph.
func Build(state *parser.TerraformState) *Graph {
	g := &Graph{}

	for _, r := range state.Values.RootModule.Resources {
		g.Nodes = append(g.Nodes, Node{
			Address: r.Address,
			Type:    r.Type,
			Name:    r.Name,
		})

		for _, dep := range r.DependsOn {
			g.Edges = append(g.Edges, Edge{
				From: r.Address,
				To:   dep,
			})
		}
	}

	return g
}
