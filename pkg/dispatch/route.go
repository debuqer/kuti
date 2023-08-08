package dispatch

import (
	"regexp"
	"strings"
)

type Node struct {
	IsParameter bool
	Name        string
	P           *Node
	Childs      []*Node
	Route       *Route
}

type Route struct {
	Name     string
	Pattern  string
	CallBack any
}

func (r *Route) addToTree(offset int, root *Node) {
	sections := strings.Split(r.Pattern, "/")[offset:]

	if len(sections) > 0 {
		var seg *Node
		newPath := true
		for _, k := range root.Childs {
			if k.Name == sections[0] {
				seg = k
				newPath = false
				break
			}
		}

		if newPath {
			seg = &Node{
				IsParameter: regexp.MustCompile(`\{.*\}`).MatchString(sections[0]),
				Name:        sections[0],
				P:           root,
			}

			root.Childs = append(root.Childs, seg)
		}

		if len(sections) == 1 {
			seg.Route = r
		}

		if len(sections) > 1 {
			r.addToTree(offset+1, seg)
		}
	}
}

func Parse(url string, root *Node) *Node {
	sections := strings.Split(url, "/")

	cur := root
	for _, section := range sections {
		for _, k := range cur.Childs {
			if k.Name == section || k.IsParameter {
				cur = k
				break
			}
		}
	}

	return cur
}
