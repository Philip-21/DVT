package json

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns  a JSON string object
func PhoneToString(phone any) (string, error) {
	return defaultTransformPhone.PhoneToString(phone)
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// to a []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func PhoneToBytes(phone any) ([]byte, error) {
	return defaultTransformPhone.PhoneToBytes(phone)
}
