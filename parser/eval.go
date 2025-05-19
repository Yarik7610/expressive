package parser

func Eval(nodes []Node) float64 {
	if len(nodes) == 0 {
		panic("eval: no nodes provided")
	}

	root := nodes[0]
	return root.Eval()
}
