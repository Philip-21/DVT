package dvl

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func ErrTime() string {
	errt :=
		`
	Ensure you have the right time stamp format [GMT, WAT, CET, EET, CST]
	XXXX-XX-XXTXX:XX:XX+00:00 GMT,
	XXXX-XX-XXTXX:XX:XX+01:00 WAT,
	XXXX-XX-XXTXX:XX:XX+01:00 CET,
	XXXX-XX-XXTXX:XX:XX+02:00 EET,
	XXXX-XX-XXTXX:XX:XX-06:00 CST,
	`
	return errt
}

var Errtime = ErrTime()

func ValidateDateTime(timeStamp string) error {
	gmtTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$` //"2023-07-10T16:22:41+00:00"
	watTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	cetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	eetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$` //"2023-07-10T16:22:41+02:00"
	cstTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$` //"2023-07-10T16:22:41-06:00"

	var regEXP *regexp.Regexp

	switch {
	case regexp.MustCompile(gmtTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(gmtTimeStamp)

	case regexp.MustCompile(watTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(watTimeStamp)

	case regexp.MustCompile(cetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cetTimeStamp)

	case regexp.MustCompile(eetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(eetTimeStamp)

	case regexp.MustCompile(cstTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cstTimeStamp)
	default:
		return fmt.Errorf("invalid DateTime format: %s\n", timeStamp, Errtime)
	}

	if !regEXP.MatchString(timeStamp) {
		return fmt.Errorf("invalid timestamp format for %s", timeStamp, Errtime)
	}

	return nil
}

func ValidateDateTimeToString(timeStamp string) (string, error) {
	gmtTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$` //"2023-07-10T16:22:41+00:00"
	watTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	cetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	eetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$` //"2023-07-10T16:22:41+02:00"
	cstTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$` //"2023-07-10T16:22:41-06:00"

	var regEXP *regexp.Regexp

	switch {
	case regexp.MustCompile(gmtTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(gmtTimeStamp)

	case regexp.MustCompile(watTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(watTimeStamp)

	case regexp.MustCompile(cetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cetTimeStamp)

	case regexp.MustCompile(eetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(eetTimeStamp)

	case regexp.MustCompile(cstTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cstTimeStamp)
	default:
		return "", fmt.Errorf("invalid DateTime format: %s\n", timeStamp, Errtime)
	}

	if !regEXP.MatchString(timeStamp) {
		return "", fmt.Errorf("invalid timestamp format for %s", timeStamp, Errtime)
	}
	jsonString, err := json.Marshal(timeStamp)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func ValidateDateTimeToBytes(timeStamp string) ([]byte, error) {
	gmtTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$` //"2023-07-10T16:22:41+00:00"
	watTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	cetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$` //"2023-07-10T16:22:41+01:00"
	eetTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$` //"2023-07-10T16:22:41+02:00"
	cstTimeStamp := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$` //"2023-07-10T16:22:41-06:00"

	var regEXP *regexp.Regexp

	switch {
	case regexp.MustCompile(gmtTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(gmtTimeStamp)

	case regexp.MustCompile(watTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(watTimeStamp)

	case regexp.MustCompile(cetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cetTimeStamp)

	case regexp.MustCompile(eetTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(eetTimeStamp)

	case regexp.MustCompile(cstTimeStamp).MatchString(timeStamp):
		regEXP = regexp.MustCompile(cstTimeStamp)
	default:
		return nil, fmt.Errorf("invalid DateTime format: %s\n", timeStamp, Errtime)
	}

	if !regEXP.MatchString(timeStamp) {
		return nil, fmt.Errorf("invalid timestamp format for %s", timeStamp, Errtime)
	}
	jsonString, err := json.Marshal(timeStamp)
	if err != nil {
		return nil, err
	}
	return jsonString, nil
}
