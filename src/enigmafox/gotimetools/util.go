/*============================================================================*/
/*----------------------------------------------------------------------------*/
/*                                                                            */
/*               gotimetools is released under the MIT License                */
/*                  https://github.com/enigmafox/gotimetools                  */
/*                                                                            */
/*----------------------------------------------------------------------------*/
/*============================================================================*/

package gotimetools

import (
	"time"
)

// DaysInMonth returns the number of days in a given month-year combo.
// Returns -1 in the event of an unknown month.
func DaysInMonth(month, year int) (int) {
	// 30 days have September, April, June, and November.
	// All the rest have 31, except for February, which has
	// 28 days each year, and 29 each leap year.
	switch month {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		if month < 1 || month > 12 {
			return -1
		}
		return 31
	}
}

// DaysInMonthT returns the number of days in the month of the provided Time.
func DaysInMonthT(t time.Time) (int) {
	return DaysInMonth(int(t.Month()), t.Year())
}

// IsLeapYear returns whether or not a provided year is a leap year.
func IsLeapYear(year int) (bool) {
	// There is a leap year every year whose number is perfectly divisible
	// by four - except for years which are both divisible by 100 and
	// not divisible by 400, ex: 1600 and 2000 are leap years, but 1700 is not.
	if year % 4 == 0 {
		if year % 100 == 0 {
			return year % 400 == 0
		}
		return true
	}
	return false
}

// IsLeapYearT returns whether or not a provided Time is a leap year.
func IsLeapYearT(t time.Time) (bool) {
	return IsLeapYear(t.Year())
}
