package jsonschema_test

import (
	"encoding/json"
	"testing"

	"github.com/minodisk/go-jsonschema"
)

func TestResolve(t *testing.T) {
	src := `"/albums/{(#/definitions/album/definitions/album_id)}/photos/{(#/definitions/photo/definitions/photo_id)}"`
	h := jsonschema.HRef{}
	if err := json.Unmarshal([]byte(src), &h); err != nil {
		t.Fatal(err)
	}

	a := jsonschema.Example{}
	if err := json.Unmarshal([]byte("12345"), &a); err != nil {
		t.Fatal(err)
	}
	p := jsonschema.Example{}
	if err := json.Unmarshal([]byte("67890"), &p); err != nil {
		t.Fatal(err)
	}

	m := map[string]*jsonschema.Schema{
		"#/definitions/album/definitions/album_id": &jsonschema.Schema{Example: a},
		"#/definitions/photo/definitions/photo_id": &jsonschema.Schema{Example: p},
	}

	if err := h.Resolve(&m); err != nil {
		t.Fatal(err)
	}

	func() {
		e := "/albums/{(#/definitions/album/definitions/album_id)}/photos/{(#/definitions/photo/definitions/photo_id)}"
		a := h.String()
		if a != e {
			t.Errorf("HRef#String() expected \n%s, but actual \n%s", e, a)
		}
	}()
	func() {
		e := "/albums/12345/photos/67890"
		a := h.ExampleString()
		if a != e {
			t.Errorf("HRef#ExampleString() expected %s, but actual %s", e, a)
		}
	}()
}
