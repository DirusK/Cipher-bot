package cipher

const delimiter = " - "

var (
	errInvalidKeyLength = "invalid key length"
	errInvalidHexText   = "invalid hex"
	errDecryptionError  = "decryption error"
)

type InvalidHexError struct {
	Message string
}

func (e InvalidHexError) Error() string {
	return errInvalidHexText + delimiter + e.Message
}

type DecryptionError struct {
	Message string
}

func (e DecryptionError) Error() string {
	return errDecryptionError + delimiter + e.Message
}
