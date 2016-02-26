# gotimetools
Small collection of tools for working with time in Go.

## Parsing and formatting time in Go is confusing
Go's `time` package enables developers to parse date-time strings into `Time` objects
and to convert `Time` objects into formatted date-time strings. Unfortunately,
the authors' choice of specifying the layout for these strings is confusing.
Rather than using some type of formatting code (like most languages), Go uses
a layout-by-example technique. This would be fine except the example formatting
schema is confusing and unintuitive.

**gotimetools** aims to provide a simple work around for this
by providing a translation layer from a PHP-esk formatting schema to the
Go formatting schema.

## LayoutToGo
`LayoutToGo` is the core of **gotimetools** and is the main reason to import the project.
`LayoutToGo` implements the translation from a human layout string to the Go layout string.

#### Signature
```go
func LayoutToGo(layout string) (string)
```

#### Example
```go
import (
    "fmt"
    "time"
    "enigmafox/gotimetools"
)

func main() {
    // Time-string to be parsed.
    myTime := "Feb 3, 2013 at 7:54pm (PST)"

    // Format to provide to LayoutToGo.
    myFormat := "%M %j, %Y at %g:%i%a (%T)"

    // Parse the time.
    t, _ := time.Parse(gotimetools.LayoutToGo(myFormat), myTime)
    fmt.Println(t)

    // Get the time string back.
    myTime = t.Format(gotimetools.LayoutToGo(myFormat))
    fmt.Println(myTime)
}
```

#### Flags

| Flag | Description | Example |
| ---- | ----------- | ------- |
| | **_Dates_** |
| `%D` | A textual representation of a day, three letters | *Mon*, *Tue*, or *Wed* |
| `%d` | Day of the month, 2 digits with leading zeros | *02* |
| `%F` | A full textual representation of a month | *January* |
| `%j` | Day of the month without leading zeros | *2* |
| `%l` | A full textual representation of the day of the week | *Monday* |
| `%M` | A textual representation of a day, three letters | *Jan* |
| `%m` | Numeric representation of a month, with leading zeros | *01* |
| `%n` | Numeric representation of a month, without leading zeros | *1* |
| `%Y` | A full numeric representation of a year, 4 digits | *2006* |
| `%y` | A two digit representation of a year | *06* |
| | **_Times_** |
| `%A` | Uppercase Ante meridiem and Post meridiem | *AM* |
| `%a` | Lowercase Ante meridiem and Post meridiem | *am* |
| `%g` | 12-hour format of an hour without leading zeros | *3* |
| `%H` | 24-hour format of an hour with leading zeros | *13* |
| `%h` | 12-hour format of an hour with leading zeros | *03* |
| `%i` | Minutes with leading zeros | *04* |
| `%O` | Difference to Greenwich time (GMT) in hours | *+0700* |
| `%P` | Difference to Greenwich time (GMT) with colon between hours and minutes | *+07:00* |
| `%s` | Seconds with leading zeros | *05* |
| `%T` | Timezone abbreviation | *GMT* |
| `%u` | Nano-seconds with leading decimal | *.000000000* |
| `%v` | Milli-seconds with leading decimal | *.000* |

## Other methods
There are a few other methods included in the the package.

### DaysInMonth and DaysInMonthT
`DaysInMonth` and `DaysInMonthT` return the number of days in a given month.

#### Signatures
```go
func DaysInMonth(month, year int) (int)
func DaysInMonthT(t time.Time) (int)
```

#### Usage
```go
import (
    "fmt"
    "time"
    "enigmafox/gotimetools"
)

func main() {
    // Will print out "29" for February 2004
    fmt.Println(gotimetools.DaysInMonth(2, 2004))

    // Will print out "31" for 2000-10-01.
    myTime := time.Date(2000, 10, 1, 1, 0, 0, 0, time.UTC)
    fmt.Println(gotimetools.DaysInMonthT(myTime))
}
```

### IsLeapYear and IsLeapYearT
`IsLeapYear` and `IsLeapYearT` return whether or not a year is a leap year.

#### Signatures
```go
func IsLeapYear(year int) (int)
func IsLeapYearT(t time.Time) (int)
```

#### Usage
```go
import (
    "fmt"
    "time"
    "enigmafox/gotimetools"
)

func main() {
    // Will print out false.
    fmt.Println(gotimetools.IsLeapYear(1990))

    // Will print out true.
    myTime := time.Date(2000, 10, 1, 1, 0, 0, 0, time.UTC)
    fmt.Println(gotimetools.IsLeapYearT(myTime))
}
```
