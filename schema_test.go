package jsonschema_test

// func TestSchemaOrBool(t *testing.T) {
// 	var json string
// 	var s jsonschema.Schema
// 	var err error
//
// 	json = `{
// 		"additionalItems": {
// 			"id": "foo"
// 		}
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !s.AdditionalItems.IsSchema {
// 		t.Errorf("should be schema")
// 	}
// 	if a := s.AdditionalItems.Schema.ID; a != "foo" {
// 		t.Errorf("id should be expected foo, but actual %s", a)
// 	}
//
// 	json = `{
// 		"additionalItems": true
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s.AdditionalItems.IsSchema {
// 		t.Errorf("shouldn't be Schema")
// 	}
// 	if !s.AdditionalItems.Bool {
// 		t.Errorf("should be true")
// 	}
//
// 	json = `{
// 		"additionalItems": false
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s.AdditionalItems.IsSchema {
// 		t.Errorf("when bool, shouldn't be schema")
// 	}
// 	if s.AdditionalItems.Bool {
// 		t.Errorf("should be false")
// 	}
// }
//
// func TestSchemaOrSchemas(t *testing.T) {
// 	var json string
// 	var s jsonschema.Schema
// 	var err error
//
// 	json = `{
// 		"items": {
// 			"id": "foo"
// 		}
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !s.Items.IsSchema {
// 		t.Fatal("should be Schema")
// 	}
// 	if a := s.Items.Schema.ID; a != "foo" {
// 		t.Errorf("id is expected foo, but actual %s", a)
// 	}
//
// 	json = `{
// 		"items": [
// 			{
// 				"id": "foo"
// 			},
// 			{
// 				"id": "bar"
// 			}
// 		]
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s.Items.IsSchema {
// 		t.Fatal("shouldn't be Schema")
// 	}
// 	if a := s.Items.Schemas[0].ID; a != "foo" {
// 		t.Errorf("id is expected foo, but actual %s", a)
// 	}
// 	if a := s.Items.Schemas[1].ID; a != "bar" {
// 		t.Errorf("id is expected foo, but actual %s", a)
// 	}
// }
//
// func TestSchemaOrStrings(t *testing.T) {
// 	var json string
// 	var s jsonschema.Schema
// 	var err error
//
// 	json = `{
// 		"dependencies": {
// 			"id": "foo"
// 		}
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !s.Dependencies.IsSchema {
// 		t.Fatal("should be Schema")
// 	}
// 	if a := s.Dependencies.Schema.ID; a != "foo" {
// 		t.Errorf("id is expected foo, but actual %s", a)
// 	}
//
// 	json = `{
// 		"dependencies": [
// 			"foo",
// 			"bar"
// 		]
// 	}`
// 	s, err = jsonschema.Parse(bytes.NewBufferString(json))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if s.Dependencies.IsSchema {
// 		t.Fatal("shouldn't be Schema")
// 	}
// 	if a := s.Dependencies.Strings[0]; a != "foo" {
// 		t.Errorf("expected foo, but actual %s", a)
// 	}
// 	if a := s.Dependencies.Strings[1]; a != "bar" {
// 		t.Errorf("expected foo, but actual %s", a)
// 	}
// }

// func TestStrings(t *testing.T) {
// 	json := `{
// 		"properties": {
// 			"foo": {
// 				"type": "number"
// 			},
// 			"bar": {
// 				"type": ["number", "boolean"]
// 			}
// 		}
// 	}`
// 	buf := bytes.NewBufferString(json)
// 	schema, err := jsonschema.Parse(buf)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if a := schema.Properties["foo"].Type[0]; a != "number" {
// 		t.Errorf("foo.Type[0] is'nt number but %+v", a)
// 	}
// 	if a := schema.Properties["bar"].Type[0]; a != "number" {
// 		t.Errorf("foo.Type[0] is'nt number but %+v", a)
// 	}
// 	if a := schema.Properties["bar"].Type[1]; a != "boolean" {
// 		t.Errorf("foo.Type[0] is'nt boolean but %+v", a)
// 	}
// }

// func TestParse(t *testing.T) {
// 	json := `{
// 		"definitions": {
// 			"foo": {
// 				"type": ["object"],
// 				"id": "foo",
// 				"title": "This is Foo schema"
// 			}
// 		},
// 		"properties": {
// 			"foo": {
// 				"$ref": "#/definitions/foo"
// 			}
// 		}
// 	}`
// 	buf := bytes.NewBufferString(json)
// 	jsonschema.Parse(buf)
// 	buf = bytes.NewBufferString(json)
// 	jsonschema.Parse(buf)
//
// 	// btpl, err := ioutil.ReadFile("schema.tpl.md")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// tpl := template.Must(template.New("mytemplate").Parse(string(btpl)))
// 	// jsonschema.Doc(os.Stdout, &r, tpl)
// }
