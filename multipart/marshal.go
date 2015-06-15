package multipart

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

const (
	Boundary = "---BoundaryX"
	part     = `--{{.boundary}}
Content-Disposition: form-data; name="[{{.name}}]"

{{.content}}
`
)

var (
	partTmpl = template.Must(template.New("part").Parse(part))
)

func Marshal(d interface{}) (string, error) {
	var str string
	switch t := d.(type) {
	default:
		return "", fmt.Errorf("unsuported multipard data: %+v", t)
	case map[string]interface{}:
		for name, content := range t {
			switch t := content.(type) {
			default:
				return "", fmt.Errorf("unsupported multipart content: %+v", t)
			case string:
				b, err := marshalPart(map[string]interface{}{
					"boundary": Boundary,
					"name":     name,
					"content":  content,
				})
				if err != nil {
					return "", err
				}
				str = string(b)
			case []interface{}:
				n := fmt.Sprintf("%s[]", name)
				results := []string{}
				for _, c := range t {
					b, err := marshalPart(map[string]interface{}{
						"boundary": Boundary,
						"name":     n,
						"content":  c,
					})
					if err != nil {
						return "", err
					}
					results = append(results, string(b))
				}
				str = strings.Join(results, "")
			}
		}
		return fmt.Sprintf("%s\n--%s--", str, Boundary), nil
	}
}

func marshalPart(data interface{}) ([]byte, error) {
	// buf := bytes.NewBuffer()
	var buf bytes.Buffer
	if err := partTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
