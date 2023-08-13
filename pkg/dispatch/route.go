package dispatch

import (
	"errors"
	"fmt"
	"strings"
)

type Segment struct {
	Name        string
	IsParameter bool
	Childs      []*Segment
	Route       *Route
}

type Route struct {
	Name     string
	Pattern  string
	CallBack any
}

var Root Segment

func (parent *Segment) Lookup(Name string) (*Segment, error) {
	for _, k := range parent.Childs {
		if k.Name == Name {
			return k, nil
		}
	}

	return nil, errors.New("404")
}

func NewRoute(r *Route) {
	r.Compile()
}

func NewSegment(parent *Segment, Name string) *Segment {
	seg, err := parent.Lookup(Name)
	if err == nil {
		return seg
	}

	new := &Segment{Name: Name}
	parent.Childs = append(parent.Childs, new)

	return new
}

func (r *Route) Compile() {
	sections := strings.Split(r.Pattern, "/")[1:]

	root := &Root
	for _, section := range sections {
		segment := NewSegment(root, section)

		root = segment
	}
	root.Route = r
}

func (seg *Segment) Draw(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("\t")
	}

	routeLabel := ""
	if seg.Route != nil {
		routeLabel = "( Route: " + seg.Route.Name + ")"
	}

	fmt.Println("+\"" + seg.Name + "\"" + routeLabel)
	for _, k := range seg.Childs {
		k.Draw(indent + 1)
	}

}

func Parse(url string) (*Route, error) {
	sections := strings.Split(""+url, "/")[1:]

	root := &Root
	for _, section := range sections {
		segment, err := root.Lookup(section)
		if err == nil {
			root = segment
		} else {
			return nil, errors.New("404")
		}
	}

	return root.Route, nil
}
