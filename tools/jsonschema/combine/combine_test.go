package combine_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/minodisk/go-jsonschema/tools/combine"
	"github.com/minodisk/go-jsonschema/tools/encoding"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCombine(t *testing.T) {
	a, err := combine.Combine("../../fixtures/yaml", "../../fixtures/yaml/_meta.yml", encoding.YAML)
	if err != nil {
		t.Fatal(err)
	}

	e, err := ioutil.ReadFile("../../fixtures/schema.yml")
	if err != nil {
		t.Fatal(err)
	}

	expected := string(e)
	actual := string(a)
	if expected != actual {
		t.Fail()
	}
}
