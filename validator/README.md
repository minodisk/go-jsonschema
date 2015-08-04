# JSON Schema Validator

Information about validation, see “5. Validation keywords sorted by instance types” in [json-schema-validation](http://json-schema.org/latest/json-schema-validation.html).

## Installation

```bash
go get github.com/minodisk/go-jsonschema/validator
```

## How to generate validation code

You can generate Go code with `jsonschema generate schema.json`.
See [jsonschema client tool](../tools/jsonschema/README.md).

## Supports

- 5.1. Validation keywords for numeric instances (number and integer)
  - [x] MultipleOf
  - [x] Maximum
  - [x] ExclusiveMaximum
  - [x] Minimum
  - [x] ExclusiveMinimum
- 5.2. Validation keywords for strings
  - [x] MaxLength
  - [x] MinLength
  - [x] Pattern
- 5.3. Validation keywords for arrays
  - [x] AdditionalItems
  - [ ] Items
  - [x] MaxItems
  - [x] MinItems
  - [x] UniqueItems
- 5.4. Validation keywords for objects
  - [x] MaxProperties
  - [x] MinProperties
  - [x] Required
  - ~~AdditionalProperties~~ Won't support[^dynamic]
  - [ ] Properties
  - ~~PatternProperties~~ Won't support[^dynamic]
  - [x] Dependencies
- 5.5. Validation keywords for any instance type
  - [ ] Enum
  - [ ] Type
  - [ ] AllOf
  - [ ] AnyOf
  - [ ] OneOf
  - [ ] Not
  - [ ] Definitions

[^dynamic]: Dynamic properties aren't allowed in static typing language like Go.
