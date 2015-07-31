go build -o bin/jsonschema tools/cli/main.go && \
./bin/jsonschema generate -o fixtures/schema.go fixtures/schema.json
