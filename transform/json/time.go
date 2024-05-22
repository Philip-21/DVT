package json

// validates a datetime to a []byte type , this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func DateTimeToBytes(timeStamp string) ([]byte, error) {
	return defaultTransformTimezone.DateTimeToBytes(timeStamp)
}

// validate's and converts datetime to a string JSON object.
func DateTimeToString(timeStamp string) (string, error) {
	return defaultTransformTimezone.DateTimeToString(timeStamp)
}
