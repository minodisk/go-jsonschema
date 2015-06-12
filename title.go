package gojsa

import (
	"fmt"
	"regexp"
	"strings"
)

var rTitleInvalidMarks = regexp.MustCompile(`[:\/]`)

type Title string

func (t Title) GracefulString() string {
	return ""
}

func (t Title) LinkString() string {
	g := t.GracefulString()
	u := strings.Replace(string(g), " ", "-", -1)
	u = rTitleInvalidMarks.ReplaceAllString(u, "")
	u = strings.ToLower(u)
	return fmt.Sprintf("#%s", u)
}
