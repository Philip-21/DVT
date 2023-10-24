# DataValidator Library

DVL is a versatile tool for validating and transforming data into JSON format. Whether you're handling API requests, database queries, or working with structured data, DataValidator simplifies the process.

## Features

-   Seamless Data Validation: Ensure your data meets the required criteria and format.
-   Data Transformation: Convert validated data into structured JSON for easy processing.
-   Versatile Use Cases: Suitable for API requests, database queries, and working with various data structures.
-   Developer-Friendly: Simple and intuitive functions for data validation and transformation.
-   Customizable: Tailor validation to your specific needs, choosing which data fields to validate.

### Kinds of Data DVL validates

-   Name: Item or Individual name
-   Email: Validate email's to its specific syntax
-   Phone number : Phone number support for Nigeria, USA, England, Germany and India formats
-   Dates and Times: Ensure that dates and times are in the correct format and fall within acceptable range and time zone.
    timezone fomats supported Include [GMT, WAT, CET, EET, CST]
-   Numeric Values : Validate numeric data, such as ages, prices, or quantities, are within acceptable ranges and have appropriate data types (e.g., integers or decimals).
-   Geographical Coordinates: Validate latitude and longitude coordinates for accuracy and within acceptable ranges.

## Installation

      go get github.com/philip21/datavalidator

## Usage

### Validate and transform your data with ease

#### Validate's Data and converts to JSON

```go
import "github.com/philip21/datavalidator/dvl"

func main() {
    email := "philip@example.com"
    err := dvl.ValidateEmail(email)
    if err != nil {
        log.Println("Error:", err)
          return
  	}
 }
```

#### Validate's Data and returns the converted JSON data to string Format for specific use

```go
import "github.com/philip21/datavalidator/dvl"

func main() {
    mobile := "+234-0000000000"
    jsonstring, err := dvl.ValidatePhoneToString(mobile)
    if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the jsonstring according to your needs
}
```

#### Validate's Data and returns the converted JSON data to raw []byte for specific use

```go
import "github.com/philip21/datavalidator/dvl"

func main() {
    mobile := "+44-0000000000"
    jsonByte, err := dvl.ValidatePhoneToBytes(mobile)
    if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the jsonByte according to your needs
}
```
#### Validate's Data with struct types 
```go
import "github.com/philip21/datavalidator/dvl"
type RequestData struct{
  Date string
}

func main() {
   dat := RequestData {
    Date : "2023-07-10T16:22:41+01:00"
 }
  jsonString, err := dvl.ValidateDateTimeToString(dat.Date)
  if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the jsonByte according to your needs
  
}
```
## Licenses
