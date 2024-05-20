package transform

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type TransformDateTime struct {
	Regexps map[string]*regexp.Regexp
}

// NewDateTimeTransform Initialize the  TransformDateTime struct
func NewDateTimeTransform() *TransformDateTime {
	return &TransformDateTime{
		Regexps: map[string]*regexp.Regexp{
			"gmt": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$`), // "2023-07-10T16:22:41+00:00"
			"wat": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"cet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"eet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$`), // "2023-07-10T16:22:41+02:00"
			"cst": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$`), // "2023-07-10T16:22:41-06:00"
		},
	}
}

// validates a datetime to a []byte type , this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func (t *TransformDateTime) DateTimeToBytes(timestamp string) ([]byte, error) {
	for _, re := range t.Regexps {
		if re.MatchString(timestamp) {
			jsonBytes, err := json.Marshal(timestamp)
			if err != nil {
				return nil, err
			}
			return jsonBytes, nil
		}
	}
	return nil, fmt.Errorf("invalid datetime format: %s", timestamp)
}

// validate's and converts datetime to a string JSON object.
func (t *TransformDateTime) DateTimeToString(timestamp string) (string, error) {
	for _, re := range t.Regexps {
		if re.MatchString(timestamp) {
			jsonBytes, err := json.Marshal(timestamp)
			if err != nil {
				return "", err
			}
			return string(jsonBytes), nil
		}
	}
	return "", fmt.Errorf("invalid datetime format: %s", timestamp)
}
