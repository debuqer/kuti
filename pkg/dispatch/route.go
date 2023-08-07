package dispatch

import (
	"fmt"
	"regexp"
	"strings"
)

type Link struct {
	Pattern string
}

type Route struct {
	Name     string
	Link     *Link
	Template string
}

func (l *Link) Extend() (format, address string) {
	segments := strings.Split(l.Pattern, "/")

	for _, segment := range segments {
		if regexp.MustCompile(`\{ (.*) files in (.*)\ }`).MatchString(segment) {

			fmt.Sscanf(segment, "{ %s files in %s }", &format, &address)
		}
	}

	return
}
