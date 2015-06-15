package jsonschema

type Title string

func (t Title) DocString() string {
	return string(t)
}

// func (t Title) LinkString() string {
// 	g := t.GracefulString()
// 	u := strings.Replace(string(g), " ", "-", -1)
// 	u = rTitleInvalidMarks.ReplaceAllString(u, "")
// 	u = strings.ToLower(u)
// 	return fmt.Sprintf("#%s", u)
// }
