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
	// Extract the country code and phone number from the "Phone" field
	// The format of the "Phone" field should be +XX-XXXXXXXXXX

	if len(rd.Phone) < 10 || rd.Phone[0] != '+' {
		return errors.New("invalid phone field format, ensure you have + added at the beginning and the right ")
	}
	// // Check for the presence of the hyphen at the correct position
	if rd.Phone[2] != '-' || rd.Phone[3] != '-' || rd.Phone[4] != '-' {
		return errors.New("invalid phone field format, missing hyphen (-) after a country code")
	}
	countryCode := rd.Phone[0:4]
	phoneNumberNigeria := rd.Phone[5:]
	phoneNumberUSA := rd.Phone[3:]
	phoneNumberEngland_Germany := rd.Phone[4:]

	var pattern string

	switch countryCode {
	case "+234":
		pattern = `^\d{10}$` //Nigeria
		match, _ := regexp.MatchString(pattern, phoneNumberNigeria)
		if !match {
			return fmt.Errorf("invalid phone number format  %s entered", countryCode)
		}
	case "+1":
		pattern = `^\d{10}$` // USA
		match, _ := regexp.MatchString(pattern, phoneNumberUSA)
		if !match {
			return fmt.Errorf("invalid phone number format  %s entered", countryCode)
		}
	case "+44":
		pattern = `^\d{10}$` // England
		match, _ := regexp.MatchString(pattern, phoneNumberEngland_Germany)
		if !match {
			return fmt.Errorf("invalid phone number format  %s entered", countryCode)
		}
	case "+49":
		pattern = `^\d{10}$` // Germany
		match, _ := regexp.MatchString(pattern, phoneNumberEngland_Germany)
		if !match {
			return fmt.Errorf("invalid phone number format  %s entered", countryCode)
		}

	default:
		return fmt.Errorf("invalid country code: %s, valid country codes are +234, +44, +1, and +49", countryCode)

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
