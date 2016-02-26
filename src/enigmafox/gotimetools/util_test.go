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

func TestDaysInMonth(t *testing.T) {
	testCases := []testIII {
		{0, 2000, -1},
		{1, 2000, 31},
		{2, 2000, 29},
		{2, 2001, 28},
		{3, 2000, 31},
		{4, 2000, 30},
		{5, 2000, 31},
		{6, 2000, 30},
		{7, 2000, 31},
		{8, 2000, 31},
		{9, 2000, 30},
		{10, 2000, 31},
		{11, 2000, 30},
		{12, 2000, 31},
		{13, 2000, -1},
	}

	for _, tc := range testCases {
		actual := DaysInMonth(tc.i1, tc.i2)
		if actual != tc.i3 {
			t.Errorf("Expected %d for [%d, %d] but got %d", tc.i3, tc.i1, tc.i2, actual)
		}
	}
}

func TestDaysInMonthT(t *testing.T) {
	testCases := []testIII {
		{1, 2000, 31},
		{2, 2000, 29},
		{2, 2001, 28},
		{3, 2000, 31},
		{4, 2000, 30},
		{5, 2000, 31},
		{6, 2000, 30},
		{7, 2000, 31},
		{8, 2000, 31},
		{9, 2000, 30},
		{10, 2000, 31},
		{11, 2000, 30},
		{12, 2000, 31},
	}

	for _, tc := range testCases {
		actual := DaysInMonthT(time.Date(tc.i2, time.Month(tc.i1), 1, 1, 0, 0, 0, time.UTC))
		if actual != tc.i3 {
			t.Errorf("Expected %d for [%d, %d] but got %d", tc.i3, tc.i1, tc.i2, actual)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	testCases := []testIB {
		{2000, true},
		{1999, false},
		{1998, false},
		{1997, false},
		{1996, true},
		{1995, false},
		{1994, false},
		{1993, false},
		{1992, true},
		{1991, false},
		{1990, false},
		{1900, false},
		{1800, false},
		{1700, false},
		{1600, true},
	}

	for _, tc := range testCases {
		if IsLeapYear(tc.i) != tc.b {
			t.Errorf("Wrong outcome on %d", tc.i)
		}
	}
}

func TestIsLeapYearT(t *testing.T) {
	testCases := []testIB {
		{2000, true},
		{1999, false},
		{1998, false},
		{1997, false},
		{1996, true},
		{1995, false},
		{1994, false},
		{1993, false},
		{1992, true},
		{1991, false},
		{1990, false},
		{1900, false},
		{1800, false},
		{1700, false},
		{1600, true},
	}

	for _, tc := range testCases {
		if IsLeapYearT(time.Date(tc.i, 1, 1, 1, 0, 0, 0, time.UTC)) != tc.b {
			t.Errorf("Wrong outcome on %d", tc.i)
		}
	}
}
