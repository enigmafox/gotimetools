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
	"testing"
	"time"
)

func TestLayoutToGo(t *testing.T) {
	testCases := []testSS {
		{"2015-07-29 16:30:00", "%Y-%m-%d %H:%i:%s"},
		{"Thu 1-1-2009 6:00", "%D %n-%j-%Y %g:%i"},
		{"Sat Jul 23 02:16:57 05 +0400 %", "%D %M %d %h:%i:%s %y %O %"},
		{"2007-11-09 T 11:20.001 UTC", "%Y-%m-%d T %H:%i%v %T"},
		{"Feb 3, 2013 at 7:54pm (PST)", "%M %j, %Y at %g:%i%a (%T)"},
		{"Feb 3, 2013 at 7:54pm (PST) %%%%%", "%M %j, %Y at %g:%i%a (%T) %%%%%"},
	}

	for _, tc := range testCases {
		// parse the time.
		layout := LayoutToGo(tc.s2)
		parsed, err := time.Parse(layout, tc.s1)
		if err != nil {
			t.Errorf("Input: [%s] [%s] [%s]. Error: %s", tc.s1, tc.s2, layout, err.Error())
		} else {
			// format time
			formatted := parsed.Format(layout)
			if formatted != tc.s1 {
				t.Errorf("Incorrect format.  Input [%s] [%s] [%s] [%s]",  tc.s1, tc.s2, layout, formatted)
			}
		}
	}
}
