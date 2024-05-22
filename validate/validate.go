package validate

import (
	"fmt"
	"regexp"
)

type ValidPhone struct {
	Countrycode map[string]*regexp.Regexp
}

type ValidDateTime struct {
	Regexps map[string]*regexp.Regexp
}

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
func (v *ValidDateTime) ValidateDateTime(timestamp string) (any, bool, error) {
	for _, re := range v.Regexps {
		if re.MatchString(timestamp) {
			return timestamp, true, nil
		}

	}
	return timestamp, false, fmt.Errorf("invalid datetime format: %s", timestamp)
}

// NewValidPhone initializes the ValidPhone struct
func NewValidPhone() *ValidPhone {
	countryCodeNGA := `^\+234-\d{10}$`
	countryCodeUSA := `^\+1-\d{10}$`
	countryCodeGER := `^\+44-\d{10}$`
	countryCodeENG := `\+49-\d{10}$`
	countryCodeIND := `\+91-\d{10}$`
	return &ValidPhone{
		Countrycode: map[string]*regexp.Regexp{
			"+234": regexp.MustCompile(countryCodeNGA),
			"+1":   regexp.MustCompile(countryCodeUSA),
			"+44":  regexp.MustCompile(countryCodeENG),
			"+49":  regexp.MustCompile(countryCodeGER),
			"+91":  regexp.MustCompile(countryCodeIND),
		},
	}
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns the validated Phone and a bool value
func (v *ValidPhone) ValidPhone(phone any) (any, bool, error) {
	for _, re := range v.Countrycode {
		phonestr, ok := phone.(string)
		if !ok {
			return nil, false, fmt.Errorf("unable to convert to string")
		}
		if len(phonestr) < 10 || phonestr[0] != '+' {
			return nil, false, fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
		}
		if len(phonestr) < 10 || phonestr[0] != '+' {
			return nil, false, fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
		}
		if re.MatchString(phonestr) {
			return phonestr, true, nil
		}
	}
	phonestr, ok := phone.(string)
	if !ok {
		return nil, false, fmt.Errorf("unable to convert to string")
	}
	return nil, false, fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
}
