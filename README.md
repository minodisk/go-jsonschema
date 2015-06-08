# go-json-schema

JSON Schema parser in Go

## Installation

```bash
go get github.com/minodisk/go-json-schema
```

## Usage

```go
import parser "github.com/minodisk/go-json-schema/parser"

func main() {
  parser.Parse()
}
```

## Support

- [json-schema-core](http://json-schema.org/latest/json-schema-core.html)
  - [ ] ID string

- [json-schema-validation](http://json-schema.org/latest/json-schema-validation.html)
  - 5. Validation keywords sorted by instance types
    - 5.1. Validation keywords for numeric instances (number and integer)
      - [ ] MultipleOf
      - [ ] Maximum
      - [ ] ExclusiveMaximum
      - [ ] Minimum
      - [ ] ExclusiveMinimum
    - 5.2. Validation keywords for strings
      - [ ] MaxLength
      - [ ] MinLength
      - [ ] Pattern
    - 5.3. Validation keywords for arrays
      - [ ] AdditionalItems
      - [ ] Items
      - [ ] MaxItems
      - [ ] MinItems
      - [ ] UniqueItems
    - 5.4. Validation keywords for objects
      - [ ] MaxProperties
      - [ ] MinProperties
      - [ ] Required
      - [ ] AdditionalProperties
      - [ ] Properties
      - [ ] PatternProperties
      - [ ] Dependencies
    - 5.5. Validation keywords for any instance type
      - [ ] Enum
      - [ ] Type
      - [ ] AllOf
      - [ ] AnyOf
      - [ ] OneOf
      - [ ] Not
      - [ ] Definitions
  - 6. Metadata keywords
    - [ ] Title
    - [ ] Description
    - [ ] Default
  - 7. Semantic validation with "format"
    - [ ] Format

- [json-schema-hypermedia](http://json-schema.org/latest/json-schema-hypermedia.html)
  - [ ] Links     []Link
  - [ ] Media     Media
  - [ ] ReadOnly  bool
  - [ ] PathStart string

## Why in Go

We need simple JSON Schema parser in Go.
