package gojs

type ReqAs struct {
	R  string
	As string
}

type ReqNode struct {
	Contents []byte
	Reqs     []ReqAs
	Map      *map[string]*ReqNode
	Included bool
}

func NewReqNode(c []byte, r []ReqAs, m *map[string]*ReqNode) *ReqNode {
	res := &ReqNode{
		Contents: c,
		Reqs:     r,
		Map:      m,
		Included: false,
	}
	return res
}
