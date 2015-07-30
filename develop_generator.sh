go-bindata -pkg generator -o tools/generator/schema.go.tmpl.go fixtures/schema.go.tmpl && \
  go build -o bin/jsonschema tools/cli/main.go && \
  ./bin/jsonschema generate -o fixtures/schema.go fixtures/schema.json
