package validate

import (
	"fmt"
	"regexp"
)

type ValidDateTime struct {
	Regexps map[string]*regexp.Regexp
}

var defaultTransform = NewDateTimeValidate()

// NewDateTimeValidate Initializes the  ValidDateTime struct
func NewDateTimeValidate() *ValidDateTime {
	return &ValidDateTime{
		Regexps: map[string]*regexp.Regexp{
			"gmt": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$`), // "2023-07-10T16:22:41+00:00"
			"wat": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"cet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"eet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$`), // "2023-07-10T16:22:41+02:00"
			"cst": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$`), // "2023-07-10T16:22:41-06:00"
		},
	}
}

// validate timezone and returns the validated timezone and a bool value
func (t *ValidDateTime) ValidateDateTime(timestamp string) (any, bool, error) {
	for _, re := range t.Regexps {
		if re.MatchString(timestamp) {
			return timestamp, true, nil
		}

	}
	return timestamp, false, fmt.Errorf("invalid datetime format: %s", timestamp)
}
