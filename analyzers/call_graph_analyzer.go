package analyzers

import (
	"bytes"
	"fmt"
	"github.com/ericr/solanalyzer/emulator"
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
	"strings"
	"text/template"
)

const (
	// CallGraphCallInternal represents an internal function call.
	CallGraphCallInternal = iota
)

const dotTemplate = `digraph solanalyzer {
	{{printf "%s" .DotAttrs.Inline}}

	node [{{printf "%s" .NodeDotAttrs}}];
	edge [{{printf "%s" .EdgeDotAttrs}}];

{{- range .SubGraphs}}
  {{printf "subgraph cluster_%s {" .ID}}
  	{{printf "%s;" .DotAttrs.Inline}}

   {{- range .Nodes}}
  	{{printf "%q [ %s ];" .Function .DotAttrs}}
	{{- end}}

  {{- range .Edges}}
  	{{printf "%q -> %q [ %s ];" .From.Function .To.Function .DotAttrs}}
	{{- end}}

	  {{println "}" }}
	{{- end}}

{{- range .Edges}}
  {{printf "%q -> %q [ %s ];" .From.Function .To.Function .DotAttrs}}
{{- end}}
}`

// CallGraphAnalyzer is an analyzer that reports issues related to function
// calls.
type CallGraphAnalyzer struct{}

// CallGraph is the root of a call graph that only has subgraphs and edges.
type CallGraph struct {
	SubGraphs    []*CallGraphSubGraph
	SubGraphsMap map[string]*CallGraphSubGraph
	Edges        []*CallGraphEdge
	EdgesMap     map[string]*CallGraphEdge
	DotAttrs     *DotAttrs
	NodeDotAttrs *DotAttrs
	EdgeDotAttrs *DotAttrs
}

// CallGraphSubGraph is a subgraph of a call graph.
type CallGraphSubGraph struct {
	ID       string
	Nodes    []*CallGraphNode
	NodesMap map[string]*CallGraphNode
	Edges    []*CallGraphEdge
	EdgesMap map[string]*CallGraphEdge
	DotAttrs *DotAttrs
}

// CallGraphNode is a call graph node.
type CallGraphNode struct {
	ID       string
	Function *sources.Function
	DotAttrs *DotAttrs
}

// CallGraphEdge is a call graph edge.
type CallGraphEdge struct {
	From     *CallGraphNode
	To       *CallGraphNode
	SubType  int
	DotAttrs *DotAttrs
}

// DotAttrs is a map of attributes for dot templates.
type DotAttrs map[string]string

// NewCallGraph returns a new instance of CallGraph.
func NewCallGraph() *CallGraph {
	return &CallGraph{
		SubGraphs:    []*CallGraphSubGraph{},
		SubGraphsMap: map[string]*CallGraphSubGraph{},
		Edges:        []*CallGraphEdge{},
		EdgesMap:     map[string]*CallGraphEdge{},
		DotAttrs: &DotAttrs{
			"rankdir": "LR",
		},
		NodeDotAttrs: &DotAttrs{
			"fontname": "Courier New",
			"fontsize": "12px",
		},
		EdgeDotAttrs: &DotAttrs{},
	}
}

// NewCallGraphNode returns a new instance of CallGraphNode.
func NewCallGraphNode(id string, fn *sources.Function) *CallGraphNode {
	return &CallGraphNode{
		ID:       id,
		Function: fn,
		DotAttrs: &DotAttrs{},
	}
}

// NewCallGraphSubGraph returns a new instance of CallGraphSubGraph.
func NewCallGraphSubGraph(id string) *CallGraphSubGraph {
	return &CallGraphSubGraph{
		ID:       id,
		Nodes:    []*CallGraphNode{},
		NodesMap: map[string]*CallGraphNode{},
		Edges:    []*CallGraphEdge{},
		EdgesMap: map[string]*CallGraphEdge{},
		DotAttrs: &DotAttrs{
			"label": id,
		},
	}
}

// NewCallGraphEdge returns a new instance of CallGraphEdge.
func NewCallGraphEdge(from *CallGraphNode, to *CallGraphNode, subType int) *CallGraphEdge {
	return &CallGraphEdge{
		From:     from,
		To:       to,
		SubType:  subType,
		DotAttrs: &DotAttrs{},
	}
}

// SubGraph creates a new graph or returns an existing one.
func (cg *CallGraph) SubGraph(id string) *CallGraphSubGraph {
	if cg.SubGraphsMap[id] == nil {
		subGraph := NewCallGraphSubGraph(id)

		cg.SubGraphs = append(cg.SubGraphs, subGraph)
		cg.SubGraphsMap[id] = subGraph
	}

	return cg.SubGraphsMap[id]
}

// AddEdge adds a new edge to the graph.
func (cg *CallGraph) AddEdge(from *CallGraphNode, to *CallGraphNode, subType int) {
	id := edgeID(from, to)

	if cg.EdgesMap[id] == nil {
		edge := NewCallGraphEdge(from, to, subType)

		cg.Edges = append(cg.Edges, edge)
		cg.EdgesMap[id] = edge
	}
}

func (cg *CallGraph) String() string {
	var output bytes.Buffer

	tmpl := template.New("dot")

	if _, err := tmpl.Parse(dotTemplate); err != nil {
		logrus.Errorf("Call graph analysis produced error: %s", err)
		return ""
	}

	if err := tmpl.Execute(&output, cg); err != nil {
		logrus.Errorf("Call graph analysis produced error: %s", err)
		return ""
	}

	return output.String()
}

// Node creates a new node or returns an existing one.
func (sg *CallGraphSubGraph) Node(fn *sources.Function) *CallGraphNode {
	id := sg.nodeID(fn)

	if sg.NodesMap[id] == nil {
		node := NewCallGraphNode(id, fn)

		sg.Nodes = append(sg.Nodes, node)
		sg.NodesMap[id] = node
	}

	return sg.NodesMap[id]
}

// AddEdge adds a new edge to a subgraph.
func (sg *CallGraphSubGraph) AddEdge(from *CallGraphNode, to *CallGraphNode, subType int) {
	id := edgeID(from, to)

	if sg.EdgesMap[id] == nil {
		edge := NewCallGraphEdge(from, to, subType)

		sg.Edges = append(sg.Edges, edge)
		sg.EdgesMap[id] = edge
	}
}

func (sg *CallGraphSubGraph) nodeID(fn *sources.Function) string {
	return fmt.Sprintf("%s/%s", sg.ID, fn)
}

// AsProps returns a string of attributes in property list format.
func (da *DotAttrs) AsProps() []string {
	props := []string{}

	for key, value := range *da {
		props = append(props, fmt.Sprintf("%s=%q", key, value))
	}

	return props
}

// Inline returns a string of attributes in in-line format.
func (da *DotAttrs) Inline() string {
	if len(da.AsProps()) > 0 {
		return strings.Join(da.AsProps(), ";")
	}
	return ""
}

func (da *DotAttrs) String() string {
	return strings.Join(da.AsProps(), " ")
}

// Name returns the name of the analyzer.
func (cga *CallGraphAnalyzer) Name() string {
	return "call graph"
}

// ID returns the unique ID of the analyzer.
func (cga *CallGraphAnalyzer) ID() string {
	return "call-graph"
}

// Execute runs the analyzer on a given source.
func (cga *CallGraphAnalyzer) Execute(source *sources.Source) ([]*Issue, error) {
	graph := NewCallGraph()
	issues := []*Issue{}

	em := emulator.New(source)
	em.OnEvent("function_definition", func(event *emulator.Event) {
		graph.SubGraph(event.Contract.Identifier).Node(event.Function)
	})
	em.OnEvent("function_call", func(event *emulator.Event) {
		// For internal calls
		subGraph := graph.SubGraph(event.Contract.Identifier)
		from := subGraph.Node(event.CallerFunction)
		to := subGraph.Node(event.Function)

		subGraph.AddEdge(from, to, CallGraphCallInternal)
	})
	em.Run()

	issues = append(issues, &Issue{
		Severity:   SeverityInfo,
		Title:      "Call Graph",
		Format:     "dot",
		Message:    graph.String(),
		analyzer:   cga,
		sourcePath: source.FilePath,
	})

	return issues, nil
}

func edgeID(from *CallGraphNode, to *CallGraphNode) string {
	return fmt.Sprintf("%s <-> %s", from.ID, to.ID)
}
