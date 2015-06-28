# go-jsonschema

JSON Schema parser and validator in Go.

- [Document Generator](tools/doc/README.md)
- [Source Combiner](tools/combine/README.md)
- Struct Generator

## Installation

```bash
go get github.com/minodisk/go-jsonschema
```

## Usage

```go
import "github.com/minodisk/go-jsonschema"

func main() {
  jsonString := `{
}`
  jsonschema.New(jsonString)
}
```

## Support

### [json-schema-core](http://json-schema.org/latest/json-schema-core.html)

- [ ] ID

### [json-schema-validation](http://json-schema.org/latest/json-schema-validation.html)

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

### [json-schema-hypermedia](http://json-schema.org/latest/json-schema-hypermedia.html)

- [ ] Links
- [ ] Media
- [ ] ReadOnly
- [ ] PathStart
