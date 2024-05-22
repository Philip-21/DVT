package json

import "github.com/philip21/dvt/transform"

// initialize default instance of Transform structs{},to be made available for package level, when called by respective functions
var (
	defaultTransformTimezone = transform.NewDateTimeTransform()
	defaultTransformPhone = transform.NewValidPhone()
)
