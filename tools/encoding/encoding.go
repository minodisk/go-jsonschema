package encoding

import (
	"fmt"
	"path/filepath"
)

type Encoding string

const (
	JSONEncoding Encoding = "json"
	YAMLEncoding Encoding = "yaml"
)

var (
	extEncodingMap = map[string]Encoding{
		".json": JSONEncoding,
		".yml":  YAMLEncoding,
		".yaml": YAMLEncoding,
	}
)

func NewWithFilename(filename string) (Encoding, error) {
	enc, ok := extEncodingMap[filepath.Ext(filename)]
	if !ok {
		return "", fmt.Errorf("no encoding '%s'", filename)
	}
	return enc, nil
}
