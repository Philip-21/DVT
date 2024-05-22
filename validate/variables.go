package validate

// initialize default instance of the Valid structs,to be made available for package level, when called by respective functions
var (
	defaultValidTimezone = NewDateTimeValidate()
	defaultValidPhone    = NewValidPhone()
)
