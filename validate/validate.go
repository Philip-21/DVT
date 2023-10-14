package validate

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func IsEmail(email string) bool {
	if !govalidator.IsEmail(email) {
		log.Println("Invalid email format")
		return false
	}
	return true
}

func (rd *RequestData) ValidatePhone() error {
	if len(rd.Phone) < 10 || rd.Phone[0] != '+' {
		return fmt.Errorf("invalid phone field format, ensure you have + added at the beginning of the phone number")
	}
	countryCodeNGA := rd.Phone[0:4]
	countryCodeUSA := rd.Phone[0:2]
	countryCodeENGGER := rd.Phone[0:3]
	var pattern string

	if countryCodeNGA == "+234" {
		// Handle Nigeria case
		pattern = `^\d{10}$`
		hyphenPos := -1
		for i := 0; i < len(rd.Phone); i++ {
			if rd.Phone[i] == '-' {
				hyphenPos = i
				break
			}
		}
		if hyphenPos == -1 {
			return fmt.Errorf("invalid phone format, missing - after %s", countryCodeNGA)
		}
		phoneNumberNigeria := rd.Phone[hyphenPos+1:]
		match, _ := regexp.MatchString(pattern, phoneNumberNigeria)
		if !match {
			return errors.New("invalid phone number format")
		}
	} else if countryCodeUSA == "+1" {
		pattern = `^\d{10}$` // USA
		hyphenPos := -1
		for i := 0; i < len(rd.Phone); i++ { //n.b i++ searches through the array to find -
			if rd.Phone[i] == '-' {
				hyphenPos = i
				break
			}
		}
		if hyphenPos == -1 {
			return fmt.Errorf("invalid phone format, missing - after %s", countryCodeUSA)
		}
		phoneNumberUSA := rd.Phone[hyphenPos+1:]
		match, _ := regexp.MatchString(pattern, phoneNumberUSA)
		if !match {
			return errors.New("invalid phone number format")
		}
	} else if countryCodeENGGER == "+44" {
		pattern = `^\d{10}$` // England
		hyphenPos := -1
		for i := 0; i < len(rd.Phone); i++ {
			if rd.Phone[i] == '-' { // '-' extract a single charac
				hyphenPos = i
				break
			}
		}
		if hyphenPos == -1 {
			return fmt.Errorf("invalid phone format, missing - after %s", countryCodeENGGER)
		}
		phoneNumberEngland := rd.Phone[hyphenPos+1:]
		match, _ := regexp.MatchString(pattern, phoneNumberEngland)
		if !match {
			return errors.New("invalid phone number format for")
		}

	} else if countryCodeENGGER == "+49" {
		pattern = `^\d{10}$` // Germany
		hyphenPos := -1
		for i := 0; i < len(rd.Phone); i++ {
			if rd.Phone[i] == '-' {
				hyphenPos = i
				break
			}
		}
		if hyphenPos == -1 {
			return fmt.Errorf("invalid phone format, missing - after %s", countryCodeENGGER)
		}
		phoneNumberGer := rd.Phone[hyphenPos+1:]
		match, _ := regexp.MatchString(pattern, phoneNumberGer)
		if !match {
			return fmt.Errorf("invalid phone number format")
		}
	} else {
		return fmt.Errorf("invalid country code:, valid country codes are +234, +44, +1, and +49")
	}

	return nil
}

// ValidateJSONData validates the RequestData structure
func ValidateJSONData(data RequestData) error {
	if data.Name == "" {
		return errors.New("name field is required")
	}

	if !IsEmail(data.Email) {
		return errors.New("invalid email format")
	}
	if err := data.ValidatePhone(); err != nil {
		return err
	}

	return nil
}