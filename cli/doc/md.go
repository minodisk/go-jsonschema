package doc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/fsnotify.v1"
	"gopkg.in/yaml.v2"

	"github.com/minodisk/jsonschema"
)

type Encoding string
type Engine string

const (
	JSONEncoding Encoding = "json"
	YAMLEncoding Encoding = "yaml"

	TextTemplateEngine Engine = "text/template"
)

var (
	extEncodingMap = map[string]Encoding{
		".json": JSONEncoding,
		".yml":  YAMLEncoding,
		".yaml": YAMLEncoding,
	}

	rHashIgnorance = regexp.MustCompile(`[:\/]`)
	funcMap        = template.FuncMap{
		"urlHash": urlHash,
	}
)

type Options struct {
	Input    string
	Encoding Encoding
	Template string
	Engine   Engine
	Output   string
	Format   string
	IsWatch  bool
}

func stringMap(input interface{}) (interface{}, error) {
	switch i := input.(type) {
	default:
		return i, nil
	case map[interface{}]interface{}:
		output := make(map[string]interface{})
		var oKey string
		for key, val := range i {
			switch k := key.(type) {
			default:
				return nil, fmt.Errorf("unsupported key type %T", k)
			case int:
				oKey = strconv.Itoa(k)
			case string:
				oKey = k
			}
			oVal, err := stringMap(val)
			if err != nil {
				return nil, err
			}
			output[oKey] = oVal
		}
		return output, nil
	case []interface{}:
		output := make([]interface{}, len(i))
		for index, val := range i {
			oVal, err := stringMap(val)
			if err != nil {
				return nil, err
			}
			output[index] = oVal
		}
		return output, nil
	}
}

func Generate(o Options) error {
	if err := generate(o); err != nil {
		return err
	}
	if o.IsWatch {
		return watch(o)
	}
	return nil
}

func watch(o Options) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	filenames := []string{o.Input, o.Template}
	dirs := make(map[string]bool)
	for i, filename := range filenames {
		filename = path.Clean(filename)
		filenames[i] = filename
		dir := path.Dir(filename)
		dirs[dir] = true
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
					if in(filenames, event.Name) {
						log.Printf("[watcher] detect modified: %s", event.Name)
						if err := generate(o); err != nil {
							log.Printf("fail to generate: %s", err)
						}
					}
				}
			case err := <-watcher.Errors:
				log.Printf("[watcher] error: %s", err)
			}
		}
	}()
	for dir, _ := range dirs {
		log.Printf("[watcher] watch dir: %s", dir)
		watcher.Add(dir)
	}

	<-done
	return nil
}

func in(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

func generate(o Options) error {
	b, err := ioutil.ReadFile(o.Input)
	if err != nil {
		return err
	}
	log.Printf("[doc] read schema file: %s", o.Input)

	if o.Encoding == "" {
		ext := path.Ext(o.Input)
		encoding := extEncodingMap[ext]
		if encoding == "" {
			encoding = JSONEncoding
		}
		o.Encoding = encoding
	}
	switch o.Encoding {
	default:
		return fmt.Errorf("unsupported encoding %s", o.Encoding)
	case JSONEncoding:
	case YAMLEncoding:
		var i interface{}
		if err := yaml.Unmarshal(b, &i); err != nil {
			return err
		}
		i, err := stringMap(i)
		if err != nil {
			return err
		}
		b, err = json.Marshal(i)
		if err != nil {
			return err
		}
	}

	schema, err := jsonschema.New(b)
	if err != nil {
		return err
	}

	b, err = ioutil.ReadFile(o.Template)
	if err != nil {
		return err
	}
	log.Printf("[doc] read template file: %s", o.Template)
	template := string(b)

	var buf bytes.Buffer
	switch o.Engine {
	case TextTemplateEngine:
		if err := Markdown(&buf, schema, "tmp", template); err != nil {
			return err
		}
		if err := ioutil.WriteFile(o.Output, buf.Bytes(), 0644); err != nil {
			return err
		}
		log.Printf("[doc] write document file: %s", o.Output)
	default:
		return fmt.Errorf("unsupported engine %s", string(o.Engine))
	}

	return nil
}

func Markdown(w io.Writer, s *jsonschema.Schema, name, text string) error {
	tmpl := template.Must(template.New(name).Funcs(funcMap).Parse(text))
	return tmpl.Execute(w, s)
}

func urlHash(str string) string {
	u := strings.Replace(string(str), " ", "-", -1)
	u = rHashIgnorance.ReplaceAllString(u, "")
	u = strings.ToLower(u)
	return fmt.Sprintf("#%s", u)
}