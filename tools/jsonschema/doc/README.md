# Document Generator

## Usage

```bash
jsonschema doc schema.json
```

## Feature

- Encodings: supports several encoding types
  - JSON
  - YAML
- Partial files: one file for one schema
- Example inference: infer `example` from `format` or `type`

## Inspired

- [r7kamura/jdoc](https://github.com/r7kamura/jdoc)

## Support

- [x] Title
- [x] Description
- [x] Properties
- [x] Link
  - [x] Request
    - [x] Queries
    - [x] Example
      - [x] Header
      - [x] Body
        - [x] JSON
        - [x] Multipart
          - [x] Single file
          - [ ] Multi-files as array
  - Response
    - Rel
      - [x] Instances
    - EncType
      - [x] Null
    - Example response
      - [x] Header
      - [x] Body
