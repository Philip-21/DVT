# DataValidatorTransform Library

dvt is a versatile tool for validating data types. Whether you're handling API requests, database queries, or working with structured data, DataValidator simplifies the process.
It also has the option for transforming validated data into various formats

## Features

-   Seamless Data Validation: Ensure your data meets the required criteria and format.
-   Data Transformation: Convert validated data into structured formats (e.g json) for easy processing.
-   Versatile Use Cases: Suitable for API requests, database queries, and working with various data structures.
-   Developer-Friendly: Simple and intuitive functions for data validation and transformation.
-   Customizable: Tailor validation to your specific needs, choosing which data fields to validate.

### Kinds of Data DVT validates

-   Email: Validate email's to its specific syntax

-   Phone number : Phone number support for Nigeria, USA, England, Germany and India formats

-   Dates and Times: Ensure that dates and times are in the correct format and fall within acceptable range and time zone.
    timezone fomats supported Include [GMT, WAT, CET, EET, CST]

-   URLs: Ensure URLs are syntactically correct and optionally reachable.

-   IP Addresses: Validate both IPv4 and IPv6 addresses.

-   Currency Codes: Validate that currency codes comply with the ISO 4217 standard.

-   Hexadecimal Values: Ensure strings are valid hexadecimal representations.

-   Numeric Values : Validate numeric data, such as ages, prices, or quantities, are within acceptable ranges and have appropriate data types (e.g., integers or decimals).

-   Geographical Coordinates: Validate latitude and longitude coordinates for accuracy and within acceptable ranges.

## Installation

      go get github.com/philip21/datavalidator

## Usage

### Validate and transform your data with ease

#### Validate's Data and converts to JSON

```go
import "github.com/philip21/dvt/validate"

func main() {
    email := "philip@example.com"
    err := validate.ValidateEmail(email)
    if err != nil {
        log.Println("Error:", err)
          return
  	}
 }
```

#### Validate and transform data to JSON string Format for specific use

```go
import "github.com/philip21/dvt/transform/json"

func main() {
    mobile := "+234-0000000000"
    jsonstring, err := json.PhoneToString(mobile)
    if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the jsonstring accordingly
}
```

#### Validate and transform data to JSON raw []byte for specific use

```go
import "github.com/philip21/dvt/transform/json"

func main() {
    mobile := "+44-0000000000"
    jsonByte, err := json.PhoneToString(mobile)
    if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the jsonByte accordingly
}
```

#### Validate's Data with struct types

```go
import "github.com/philip21/dvt/validate"
type RequestData struct{
  Date string
}

func main() {
   dat := RequestData {
    Date : "2023-07-10T16:22:41+01:00"
 }
  date, boolValues, err := validate.DateTime(dat.Date)
  if err != nil {
        log.Println("Error:", err)
        return
      }
    //handle the variables accordingly

}
```

## License

See dvt MIT [License](https://github.com/Philip-21/dvt/blob/master/license.md)
