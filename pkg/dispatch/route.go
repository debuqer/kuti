package dispatch

import (
	"strings"
)

type Segment struct {
	IsParameter bool
	IsTerminal  bool
	Name        string
	P           *Segment
	Childs      []*Segment
	Route       *Route
}

type Route struct {
	Name     string
	Pattern  string
	CallBack any
}

func (r *Route) addToTree(offset int, root *Segment) {
	sections := strings.Split(r.Pattern, "/")[offset:]

	if len(sections) > 0 {
		seg := Segment{
			IsParameter: false,
			Name:        sections[0],
			P:           root,
			IsTerminal:  (len(sections) == 1),
		}

		if seg.IsTerminal {
			seg.Route = r
		} else {
			root.Childs = append(root.Childs, &seg)
			r.addToTree(offset+1, &seg)
		}
	}
}
