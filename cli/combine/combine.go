package combine

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	DefaultMetaFilename = "_meta.yml"
)

type Data struct {
	Schema      string `yaml:"$schema"`
	Type        interface{}
	Title       string
	Description string
	Links       []map[string]string
	Definitions map[string]interface{}
	Properties  map[string]map[string]interface{}
}

// func Combine(dir string, meta string, dest string) {
// }

func Combine(dirname string, meta string) ([]byte, error) {
	d, err := filepath.Abs(dirname)
	if err != nil {
		return nil, err
	}

	if meta == "" {
		meta = filepath.Join(d, DefaultMetaFilename)
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
	// var data map[string]interface{}

	// data := new(Data)
	var data Data
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

	out, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return out, nil
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
