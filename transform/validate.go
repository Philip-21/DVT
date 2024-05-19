package dvt

// func ValidateDateTime(timeStamp string) error {
// 	var gmtTimeStamp string
// 	var watTimeStamp string
// 	var cetTimeStamp string
// 	var eetTimeStamp string
// 	var cstTimeStamp string

// 	if gmtTimeStamp == `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$` {
// 		regEXP := regexp.MustCompile(gmtTimeStamp)
// 		match := regEXP.MatchString(timeStamp)
// 		if !match {
// 			return fmt.Errorf("invalid timestamp format for %s", timeStamp)
// 		}

// 	}
// 	return nil
// }

// ValidateJSONData validates the RequestData structure

// func ValidateJSONData(data RequestData) ([]byte, error) {
// 	validateFields := []string{"Name", "Email", "Phone", "DateTime"}

// 	// Create a map for the validated data
// 	validatedData := make(map[string]interface{})

// 	// Loop through the fields to validate
// 	for _, field := range validateFields {
// 		switch field {

// 		case "Name":
// 			if data.Name == "" {
// 				return nil, errors.New("name field is required")
// 			}
// 			validatedData["Name"] = data.Name
// 		case "Email":
// 			if !ValidateEmail(data.Email) {
// 				return errors.New("invalid email format")
// 			}
// 			validatedData["Email"] = data.Email
// 		case "Phone":
// 			err := data.ValidatePhone()
// 			if err != nil {
// 				return nil, err
// 			}
// 			validatedData["Phone"] = data.Phone
// 			//  case "DateTime":
// 			// 	 if data.DateTime == "" {
// 			// 		 return nil, errors.New("DateTime field is required")
// 			// 	 }
// 			// 	 err := ValidateDateTime(data.DateTime)
// 			// 	 if err != nil {
// 			// 		 return nil, err
// 			// 	 }
// 			// 	 validatedData["DateTime"] = data.DateTime
// 		}
// 	}

// 	// Marshal the validated data to JSON
// 	jsonData, err := json.Marshal(validatedData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jsonData, nil

// }
