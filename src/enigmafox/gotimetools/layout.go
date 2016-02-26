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
	"bytes"
)

// conversionMap holds a mapping from PHP style formatting to Go layouts.
var conversionMap = map[byte]string {
		//========== DATES ==========//
		'D': "Mon", // A textual representation of a day, three letters
		'd': "02",  // Day of the month, 2 digits with leading zeros
		'F': "January", // A full textual representation of a month, such as January or March
		'j': "2", // Day of the month without leading zeros
		'l': "Monday", // A full textual representation of the day of the week
		'M': "Jan", // A textual representation of a month, three letters
		'm': "01", // Numeric representation of a month, with leading zeros
		'n': "1", // Numeric representation of a month, without leading zeros
		'Y': "2006", // A full numeric representation of a year, 4 digits
		'y': "06", // A two digit representation of a year

		//========== TIMES ==========//
		'A': "PM", // Uppercase Ante meridiem and Post meridiem
		'a': "pm", // Lowercase Ante meridiem and Post meridiem
		'g': "3", // 12-hour format of an hour without leading zeros
		'H': "15", // 24-hour format of an hour with leading zeros
		'h': "03", // 12-hour format of an hour with leading zeros
		'i': "04", // Minutes with leading zeros
		'O': "-0700", // Difference to Greenwich time (GMT) in hours
		'P': "-07:00", // Difference to Greenwich time (GMT) with colon between hours and minutes
		's': "05", // Seconds with leading zeros
		'T': "MST", // Timezone abbreviation
		'u': ".000000000", // Nano-seconds with leading decimal
		'v': ".000", // Milli-seconds with leading decimal
	}

// LayoutToGo converts a date-time layout string in PHP-esk format
// into the unintuitive Go layout which can be used by time.Format()
// and time.Parse().
func LayoutToGo(layout string) (string) {
	var buf bytes.Buffer
	lenLayout := len(layout)

	// For each character...
	for i := 0; i < lenLayout; i++ {
		cur := layout[i]
		if cur == '%' && i < lenLayout - 1 {
			next := layout[i + 1]

			// Find the translation in the conversion map.
			// If none exists, just pass the character to the output.
			if out, exists := conversionMap[next]; exists {
				buf.WriteString(out)
				i++
			} else if next == '%' {
				buf.WriteString("%%")
				i++
			} else {
				buf.WriteByte(cur)
			}
		} else {
			buf.WriteByte(cur)
		}
	}

	return buf.String()
}
