package main

import (
	"log"

	"github.com/philip21/datavalidator/dvl"
	"github.com/philip21/datavalidator/validate"
)

func main() {
	// Define a sample RequestData with only Name, Email, and a Nigerian Phone number

	dat := validate.RequestData{
		Email: "johndoe@example.com",
	}
	num := "+234-8166959918"
	nameemail := "hilary@gmail.com"
	err := dvl.ValidateEmail(nameemail)
	if err != nil {
		log.Println("Error:", err)
		return 
	}
	st, err := dvl.ValidateEmailToString(nameemail)
	if err != nil {
		log.Println("Error:", err)
		return 
	} else {
		log.Println("Data is Valid", st)
	}
	jdata, err := dvl.ValidateEmailtoBytes(nameemail)
	if err != nil {
		log.Println("Error:", err)
		return
	} else {
		log.Println("Data is Valid", jdata)
	}
	str, err := dvl.ValidateEmailToString(dat.Email)
	if err != nil {
		log.Println("Error:", err)
		return
	} else {
		log.Println("Data is Valid", str)
	}
	nh, err := validate.ValidatePhoneToString(num)
	if err != nil {
		log.Println(err)
	}
	log.Println(nh)

}
