# DataValidator Library
  DVL is a versatile tool for validating and transforming data into JSON format. Whether you're handling API requests, database queries, or working with structured data, DataValidator simplifies the process.

## Features

- Seamless Data Validation: Ensure your data meets the required criteria and format.
- Data Transformation: Convert validated data into structured JSON for easy processing.
- Versatile Use Cases: Suitable for API requests, database queries, and working with various data structures.
- Developer-Friendly: Simple and intuitive functions for data validation and transformation.
- Customizable: Tailor validation to your specific needs, choosing which data fields to validate.

### Kinds of Data DVL validates

-   Name: Item or Individual name in string format
-   email: Validate email's to its specific syntax
-   Phone number : Phone number support for Nigeria, USA, England and Germany formats
-   Dates and Times: Ensure that dates and times are in the correct format and fall within acceptable range and time zone.
-   Numeric Values : Validate numeric data, such as ages, prices, or quantities, are within acceptable ranges and have appropriate data types (e.g., integers or decimals).
-   Geographical Coordinates: Validate latitude and longitude coordinates for accuracy and within acceptable ranges.

## Installation
      go get github.com/philip21/datavalidator

## Usage

  ### Validate and transform your data with ease 
  #### Validate email and convert to JSON

```go
import "github.com/philip21/datavalidator/dvl"

func main() {
    email := "john@example.com"
    err := dvl.ValidateEmail(email)
    if err != nil {
	    	log.Println("Error:", err)
	    	return 
  	}
}
```
#### Validate email, 
 
