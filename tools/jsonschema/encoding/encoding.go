package encoding

import (
	"fmt"
	"path/filepath"
)

type Encoding string

const (
	JSON Encoding = "json"
	YAML Encoding = "yaml"
)

var (
	extEncodingMap = map[string]Encoding{
		".json": JSON,
		".yml":  YAML,
		".yaml": YAML,
	}
)

func NewWithFilename(filename string) (Encoding, error) {
	enc, ok := extEncodingMap[filepath.Ext(filename)]
	if !ok {
		return "", fmt.Errorf("no encoding '%s'", filename)
	}
	return enc, nil
}
