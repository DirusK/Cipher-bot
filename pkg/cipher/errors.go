package cipher

const delimiter = " - "

var (
	errInvalidKeyLength = "invalid key length"
	errInvalidHexText   = "invalid hex"
)

type InvalidHexError struct {
	Message string
}

func (e InvalidHexError) Error() string {
	return errInvalidHexText + delimiter + e.Message
}
