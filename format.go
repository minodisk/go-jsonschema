package jsonschema

type Format uint8

const (
	FormatDateTime Format = iota
	FormatEmal
	FormatHostname
	FormatIPV4
	FormatIPV6
	FormatURI
)
