package dispatch

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Parameter struct {
	Name  string
	Value string
}

type Link struct {
	Pattern      string
	HasParameter bool
	Parameters   map[string]*Parameter
}

type Route struct {
	Name     string
	Link     *Link
	Template string
}

func (l *Link) Extend() {
	var address string
	l.Parameters = make(map[string]*Parameter)

	segments := strings.Split(l.Pattern, "/")

	for _, segment := range segments {
		if regexp.MustCompile(`\{ files in (.*) \}`).MatchString(segment) {

			fmt.Sscanf(segment, "{ files in %s }", &address)
			fmt.Println(address)

			if address != "" {
				l.HasParameter = true
				l.Parameters[strconv.Itoa(len(l.Parameters))] = &Parameter{Name: address}
			}
		}
	}
}
