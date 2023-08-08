package dispatch

import (
	"regexp"
	"strings"
)

type Segment struct {
	IsParameter bool
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
		var seg *Segment
		newPath := true
		for _, k := range root.Childs {
			if k.Name == sections[0] {
				seg = k
				newPath = false
				break
			}
		}

		if len(sections) >= 1 {
			if newPath {
				seg = &Segment{
					IsParameter: regexp.MustCompile(`\{.*\}`).MatchString(sections[0]),
					Name:        sections[0],
					P:           root,
				}

				root.Childs = append(root.Childs, seg)
			}

			r.addToTree(offset+1, seg)
		} else {
			seg.Route = r
		}
	}
}

func Parse(url string, root *Segment) *Segment {
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
