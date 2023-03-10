package hammy

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Stringy interface {
	~string
}

type AssertionMessage struct {
	Message      string
	IsSuccessful bool
}
