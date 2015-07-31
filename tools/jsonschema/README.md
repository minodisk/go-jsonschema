# Command-line tool

## Installation

```bash
go get github.com/minodisk/go-jsonschema/tools/jsonschema
```

## Usage

### Help

```bash
jsonschema help
```

### Generate implementation in Go

#### Struct

```bash
jsonschema struct schema.json
```

#### Routing

```bash
jsonschema router schema.json
```

#### Validator

```bash
jsonschema validator schema.json
```

#### Client

Generate client code to access to APIs.
You can make SDK of APIs easily.

```bash
jsonshcema client schema.json
```

You can find other options with `jsonschema client -h`.

#### Test

```bash
jsonschema test schema.json
```

#### Mock server

```bash
jsonschema mock schema.json
```

Then mock server will be started.
You can access http://localhost:3333 (in default).
You can find other options with `jsonschema mock -h`.

### Generate document in markdown or HTML (any format you like)

```bash
jsonschema doc -o schema.md schema.json
```

You can output other formats you like with template.

```bash
jsonschema doc -o schema.html -t schema.html.tmpl schema.json
```

### Combine schema files in YAML or JSON into one schema file in JSON

```bash
jsonschema combine -m meta.yml -o schema.json ./fixtures/yaml
```
