package validate

import (
	"fmt"
	"regexp"
)

// validate time and returns the validated timezone and a bool value
func ValidateDateTime(timestamp string) (any, bool, error) {
	return defaultValidTimezone.ValidateDateTime(timestamp)
}

// redundant
func DateTime(timeStamp string) (any, bool, error) {
	gmtTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$` //"2023-07-10T16:22:41+00:00"
	watTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	cetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	eetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$` //"2023-07-10T16:22:41+02:00"
	cstTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$` //"2023-07-10T16:22:41-06:00"

	//compile the regular expression once
	gmtRegexp := regexp.MustCompile(gmtTimeStamp)
	watRegexp := regexp.MustCompile(watTimeStamp)
	cetRegexp := regexp.MustCompile(cetTimeStamp)
	eetRegexp := regexp.MustCompile(eetTimeStamp)
	cstRegexp := regexp.MustCompile(cstTimeStamp)

	switch {
	case gmtRegexp.MatchString(timeStamp):
		return timeStamp, true, nil
	case watRegexp.MatchString(timeStamp):
		return timeStamp, true, nil
	case cetRegexp.MatchString(timeStamp):
		return timeStamp, true, nil
	case eetRegexp.MatchString(timeStamp):
		return timeStamp, true, nil
	case cstRegexp.MatchString(timeStamp):
		return timeStamp, true, nil
	default:
		return nil, false, fmt.Errorf("invalid datetime format: %s", timeStamp)
	}

}
