package dispatch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Param struct {
	Key   string
	Value string
}

type Segment struct {
	Name        string
	IsParameter bool
	Childs      []*Segment
	Route       *Route
}

type Route struct {
	Methods  []string
	Name     string
	Pattern  string
	CallBack any
}

type RouteOptions struct {
	Route  *Route
	Params map[string]string
}

var Router Segment

func (parent *Segment) Match(Name string) (*Segment, *Param, error) {
	for _, k := range parent.Childs {
		if regexp.MustCompile(`\:(.*)`).MatchString(k.Name) {
			pname := regexp.MustCompile(`\:(.*)`).ReplaceAllString(k.Name, "$1")

			return k, &Param{Key: pname, Value: Name}, nil
		} else if k.Name == Name {
			return k, nil, nil
		}
	}

	return nil, nil, errors.New("404")
}

func (parent *Segment) Lookup(Name string) (*Segment, error) {
	for _, k := range parent.Childs {
		if k.Name == Name {
			return k, nil
		}
	}

	return nil, errors.New("404")
}

func POST(r *Route) {
	r.Methods = append(r.Methods, "POST")
	r.Compile()
}

func GET(r *Route) {
	r.Methods = append(r.Methods, "GET")
	r.Compile()
}

func PUT(r *Route) {
	r.Methods = append(r.Methods, "PUT")
	r.Compile()
}

func DELETE(r *Route) {
	r.Methods = append(r.Methods, "DELETE")
	r.Compile()
}

func HEAD(r *Route) {
	r.Methods = append(r.Methods, "HEAD")
	r.Compile()
}

func Method(methods []string, r *Route) {
	r.Methods = methods
	r.Compile()
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

	root := &Router
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

func Parse(url string) (*RouteOptions, error) {
	if url == "" {
		url = "/"
	}
	sections := strings.Split(url, "/")[1:]

	params := make(map[string]string)

	root := &Router
	for _, section := range sections {
		segment, param, err := root.Match(section)
		if err == nil {
			if param != nil {
				params[param.Key] = param.Value
			}
			root = segment
		} else {
			return nil, errors.New("404")
		}
	}

	return &RouteOptions{Route: root.Route, Params: params}, nil
}
