package validate

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns the validated Phone and a bool value
func ValidatePhone(phone any) (any, bool, error) {
	return defaultValidPhone.ValidPhone(phone)
}
