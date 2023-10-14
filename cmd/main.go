package main

import (
	"fmt"
	"philip-datavalidator/validate"
)

func main() {
	// Define a sample RequestData with only Name, Email, and a Nigerian Phone number
	data := validate.RequestData{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Phone: "234-9077421913",
	}

	// Validate the data
	err := validate.ValidateJSONData(data)
	if err != nil {
		fmt.Println("Validation Error:", err)
	} else {
		fmt.Println("Data is valid")
	}
}
