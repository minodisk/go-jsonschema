package doc_test

// func TestMarkdown(t *testing.T) {
// 	// y, err := ioutil.ReadFile("../../fixtures/schema.yml")
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// s, err := jsonschema.NewYAML(y)
// 	j, err := ioutil.ReadFile("../../fixtures/schema.json")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	s, err := jsonschema.NewJSON(j)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	tmpl, err := ioutil.ReadFile("../../fixtures/schema.md.tmpl")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	buf := bytes.NewBuffer(nil)
// 	if err := doc.Markdown(buf, s, "md", string(tmpl)); err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// fmt.Println(string(buf.Bytes()))
//
// 	ioutil.WriteFile("../../fixtures/schema.go.md", buf.Bytes(), 0644)
// }
