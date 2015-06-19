package combine_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/aryann/difflib"
	"github.com/minodisk/go-jsonschema/cli/combine"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCombine(t *testing.T) {
	a, err := combine.Combine("../../fixtures/schemata", "")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile("../../fixtures/schema.yml")
	if err != nil {
		t.Fatal(err)
	}
	var tmp interface{}
	if err := yaml.Unmarshal(b, &tmp); err != nil {
		t.Fatal(err)
	}
	e, err := yaml.Marshal(tmp)
	if err != nil {
		t.Fatal(err)
	}

	diff := difflib.Diff(strings.Split(string(e), "\n"), strings.Split(string(a), "\n"))
	for _, d := range diff {
		fmt.Println(d)
	}
}
