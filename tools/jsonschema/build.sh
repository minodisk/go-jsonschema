go generate ./...
go build
./jsonschema combine -m examples/yaml/_meta.yml -o examples/schema.json examples/yaml
./jsonschema generate -o examples/generated_codes examples/schema.json
