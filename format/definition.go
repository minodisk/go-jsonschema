package format

type Definition struct {
	Name     string
	Example  string
	Validate func(v interface{}) bool
}

var (
	DateTime = Definition{
		Name:     "date-time",
		Example:  "2006-01-02 15:04:06",
		Validate: func(v interface{}) bool { return false },
	}
	Email = Definition{
		Name:     "email",
		Example:  "mail@example.com",
		Validate: func(v interface{}) bool { return false },
	}
	Hostname = Definition{
		Name:     "hostname",
		Example:  "example.com",
		Validate: func(v interface{}) bool { return false },
	}
	IPV4 = Definition{
		Name:     "ipv4",
		Example:  "0.0.0.0",
		Validate: func(v interface{}) bool { return false },
	}
	IPV6 = Definition{
		Name:     "ipv6",
		Example:  "0:0:0:0:0:0:0:0",
		Validate: func(v interface{}) bool { return false },
	}
	URI = Definition{
		Name:     "uri",
		Example:  "http://example.com",
		Validate: func(v interface{}) bool { return false },
	}
)
