# jsonschema

JSON Schema parser in Go

## Installation

```bash
go get github.com/minodisk/go-jsonschema
```

## Usage

```go
import "github.com/minodisk/go-jsonschema"

func main() {
  jsonString := `
{
}
`
  jsonschema.NewSchema(jsonString)
}
```

## Why in Go

We need simple JSON Schema parser in Go.
