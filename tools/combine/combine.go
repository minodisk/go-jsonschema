package combine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/minodisk/go-jsonschema/tools/encoding"
	"github.com/minodisk/go-watcher"

	"gopkg.in/yaml.v2"
)

type Options struct {
	Input    string
	Meta     string
	Output   string
	Encoding encoding.Encoding
	IsWatch  bool
}

func Run(o Options) (err error) {
	err = run(o)
	if err != nil {
		return err
	}
	if o.IsWatch {
		done := make(chan int)
		w := watcher.New()
		go func() {
			for {
				select {
				case <-w.Events:
					err = run(o)
					if err != nil {
						log.Println("Running error:", err)
					}
				case err := <-w.Errors:
					log.Println("Error:", err)
				}
			}
		}()
		err = w.Watch([]string{o.Input, o.Meta})
		if err != nil {
			return err
		}
		<-done
	}
	return nil
}

func run(o Options) (err error) {
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
	dir, err := filepath.Abs(input)
	if err != nil {
		return nil, err
	}

	meta, err = filepath.Abs(meta)
	if err != nil {
		return nil, err
	}

	files := map[string][]byte{}
	if err := readFiles(dir, &files); err != nil {
		return nil, err
	}
	_, ok := files[meta]
	if ok {
		delete(files, meta)
		// return nil, fmt.Errorf("meta file is required: %s", meta)
	}

	var schema struct {
		Schema      string                            `yaml:"$schema" json:"$schema"`
		Type        interface{}                       `json:"type"`
		Title       string                            `json:"title"`
		Description string                            `json:"description"`
		Links       []map[string]string               `json:"links"`
		Definitions map[string]interface{}            `json:"definitions"`
		Properties  map[string]map[string]interface{} `json:"properties"`
	}

	data, err := ioutil.ReadFile(meta)
	if err != nil {
		return nil, fmt.Errorf("meta file is required: %s", meta)
	}
	metaEnc, err := encoding.NewWithFilename(meta)
	if err != nil {
		return nil, err
	}
	switch metaEnc {
	default:
		return nil, fmt.Errorf("unsupported encoding '%s'", metaEnc)
	case encoding.JSON:
		if err := json.Unmarshal(data, &schema); err != nil {
			return nil, err
		}
	case encoding.YAML:
		if err := yaml.Unmarshal(data, &schema); err != nil {
			return nil, err
		}
	}

	schema.Definitions = make(map[string]interface{})
	schema.Properties = make(map[string]map[string]interface{})

	for filename, b := range files {
		var s interface{}
		if err := yaml.Unmarshal(b, &s); err != nil {
			return nil, err
		}
		subSchema, err := encoding.KeyValueMap(s)
		if err != nil {
			return nil, err
		}
		subSchemaMap, ok := subSchema.(map[string]interface{})
		if !ok {
			log.Printf("ignore %s: invalid schema", filename)
			continue
		}
		i, ok := subSchemaMap["id"]
		var id string
		if ok {
			id = i.(string)
		} else {
			id = strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
		}
		schema.Definitions[id] = subSchemaMap
		schema.Properties[id] = make(map[string]interface{})
		schema.Properties[id]["$ref"] = fmt.Sprintf("#/definitions/%s", id)
	}

	switch enc {
	default:
		return nil, fmt.Errorf("unsupported encoding '%s'", enc)
	case encoding.JSON:
		return json.MarshalIndent(schema, "", "  ")
	case encoding.YAML:
		return yaml.Marshal(schema)
	}
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
