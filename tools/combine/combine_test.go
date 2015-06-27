package combine_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aryann/difflib"
	"github.com/minodisk/go-jsonschema/tools/combine"
	"github.com/minodisk/go-jsonschema/tools/encoding"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCombine(t *testing.T) {
	a, err := combine.Combine("../../fixtures/schemata/yaml", "../../fixtures/schemata/yaml/_meta.yml", encoding.JSON)
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
		diff := difflib.Diff(strings.Split(expected, "\n"), strings.Split(actual, "\n"))
		for _, d := range diff {
			fmt.Println(d)
		}
		t.Fail()
	}
}
