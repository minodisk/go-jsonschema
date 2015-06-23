package combine

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/minodisk/go-jsonschema/tools/encoding"

	"gopkg.in/yaml.v2"
)

type Options struct {
	Input    string
	Meta     string
	Output   string
	Encoding encoding.Encoding
}

func Run(o Options) (err error) {
	b, err := Combine(o.Input, o.Meta, o.Encoding)
	if err != nil {
		return err
	}

	if o.Output == "" {
		fmt.Println(string(b))
		return nil
	}

	return ioutil.WriteFile(o.Output, b, 0644)
}

func Combine(input string, meta string, enc encoding.Encoding) (combined []byte, err error) {
	d, err := filepath.Abs(input)
	if err != nil {
		return nil, err
	}

	meta, err = filepath.Abs(meta)
	if err != nil {
		return nil, err
	}

	files := map[string][]byte{}
	if err := readFiles(d, &files); err != nil {
		return nil, err
	}
	b, ok := files[meta]
	if !ok {
		return nil, fmt.Errorf("meta file is required: %s", meta)
	}
	delete(files, meta)

	var data struct {
		Schema      string `yaml:"$schema"`
		Type        interface{}
		Title       string
		Description string
		Links       []map[string]string
		Definitions map[string]interface{}
		Properties  map[string]map[string]interface{}
	}

	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	data.Definitions = make(map[string]interface{})
	data.Properties = make(map[string]map[string]interface{})
	for filename, b := range files {
		var d map[string]interface{}
		if err := yaml.Unmarshal(b, &d); err != nil {
			return nil, err
		}
		i, ok := d["id"]
		var id string
		if ok {
			id = i.(string)
		} else {
			id = strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
		}
		data.Definitions[id] = d
		data.Properties[id] = make(map[string]interface{})
		data.Properties[id]["$ref"] = fmt.Sprintf("#/definitions/%s", id)
	}

	combined, err = yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return combined, nil
}

func readFiles(dirname string, files *map[string][]byte) error {
	fs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, f := range fs {
		name := filepath.Join(dirname, f.Name())
		if f.IsDir() {
			readFiles(name, files)
			continue
		}
		b, err := ioutil.ReadFile(name)
		if err != nil {
			return err
		}
		(*files)[name] = b
	}
	return nil
}
