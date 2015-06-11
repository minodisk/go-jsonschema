package gojsa

import (
	"fmt"
	"regexp"
	"strings"
)

var rTitleInvalidMarks = regexp.MustCompile(`[:\/]`)

type Title string

func (t Title) LinkString() string {
	u := strings.Replace(string(t), " ", "-", -1)
	u = rTitleInvalidMarks.ReplaceAllString(u, "")
	u = strings.ToLower(u)
	return fmt.Sprintf("#%s", u)
}
