go generate ./...
go build
./jsonschema combine -m fixtures/yaml/_meta.yml -o fixtures/schema.json fixtures/yaml
./jsonschema generate -o fixtures/schema.go fixtures/schema.json
