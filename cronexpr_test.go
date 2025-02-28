/*!
 * Copyright 2013 Raymond Hill
 *
 * Project: github.com/gorhill/cronexpr
 * File: cronexpr_test.go
 * Version: 1.0
 * License: pick the one which suits you best:
 *   GPL v3 see <https://www.gnu.org/licenses/gpl.html>
 *   APL v2 see <http://www.apache.org/licenses/LICENSE-2.0>
 *
 */

package cronexpr_test

/******************************************************************************/

import (
	"testing"
	"time"

	"github.com/n4n0lix/cronexpr"
)

/******************************************************************************/

type crontimes struct {
	from string
	next string
}

type crontest struct {
	expr   string
	layout string
	times  []crontimes
}

var crontests = []crontest{
	// Seconds
	{
		"* * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2013-01-01 00:00:01"},
			{"2013-01-01 00:00:59", "2013-01-01 00:01:00"},
			{"2013-01-01 00:59:59", "2013-01-01 01:00:00"},
			{"2013-01-01 23:59:59", "2013-01-02 00:00:00"},
			{"2013-02-28 23:59:59", "2013-03-01 00:00:00"},
			{"2016-02-28 23:59:59", "2016-02-29 00:00:00"},
			{"2012-12-31 23:59:59", "2013-01-01 00:00:00"},
		},
	},

	// every 5 Second
	{
		"*/5 * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2013-01-01 00:00:05"},
			{"2013-01-01 00:00:59", "2013-01-01 00:01:00"},
			{"2013-01-01 00:59:59", "2013-01-01 01:00:00"},
			{"2013-01-01 23:59:59", "2013-01-02 00:00:00"},
			{"2013-02-28 23:59:59", "2013-03-01 00:00:00"},
			{"2016-02-28 23:59:59", "2016-02-29 00:00:00"},
			{"2012-12-31 23:59:59", "2013-01-01 00:00:00"},
		},
	},

	// Minutes
	{
		"* * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2013-01-01 00:01:00"},
			{"2013-01-01 00:00:59", "2013-01-01 00:01:00"},
			{"2013-01-01 00:59:00", "2013-01-01 01:00:00"},
			{"2013-01-01 23:59:00", "2013-01-02 00:00:00"},
			{"2013-02-28 23:59:00", "2013-03-01 00:00:00"},
			{"2016-02-28 23:59:00", "2016-02-29 00:00:00"},
			{"2012-12-31 23:59:00", "2013-01-01 00:00:00"},
		},
	},

	// Minutes with interval
	{
		"17-43/5 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2013-01-01 00:17:00"},
			{"2013-01-01 00:16:59", "2013-01-01 00:17:00"},
			{"2013-01-01 00:30:00", "2013-01-01 00:32:00"},
			{"2013-01-01 00:50:00", "2013-01-01 01:17:00"},
			{"2013-01-01 23:50:00", "2013-01-02 00:17:00"},
			{"2013-02-28 23:50:00", "2013-03-01 00:17:00"},
			{"2016-02-28 23:50:00", "2016-02-29 00:17:00"},
			{"2012-12-31 23:50:00", "2013-01-01 00:17:00"},
		},
	},

	// Minutes interval, list
	{
		"15-30/4,55 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2013-01-01 00:15:00"},
			{"2013-01-01 00:16:00", "2013-01-01 00:19:00"},
			{"2013-01-01 00:30:00", "2013-01-01 00:55:00"},
			{"2013-01-01 00:55:00", "2013-01-01 01:15:00"},
			{"2013-01-01 23:55:00", "2013-01-02 00:15:00"},
			{"2013-02-28 23:55:00", "2013-03-01 00:15:00"},
			{"2016-02-28 23:55:00", "2016-02-29 00:15:00"},
			{"2012-12-31 23:54:00", "2012-12-31 23:55:00"},
			{"2012-12-31 23:55:00", "2013-01-01 00:15:00"},
		},
	},

	// Days of week
	{
		"0 0 * * MON",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Mon 2013-01-07 00:00"},
			{"2013-01-28 00:00:00", "Mon 2013-02-04 00:00"},
			{"2013-12-30 00:30:00", "Mon 2014-01-06 00:00"},
		},
	},
	{
		"0 0 * * friday",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Fri 2013-01-04 00:00"},
			{"2013-01-28 00:00:00", "Fri 2013-02-01 00:00"},
			{"2013-12-30 00:30:00", "Fri 2014-01-03 00:00"},
		},
	},
	{
		"0 0 * * 6,7",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Sat 2013-01-05 00:00"},
			{"2013-01-28 00:00:00", "Sat 2013-02-02 00:00"},
			{"2013-12-30 00:30:00", "Sat 2014-01-04 00:00"},
		},
	},

	// Specific days of week
	{
		"0 0 * * 6#5",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-09-02 00:00:00", "Sat 2013-11-30 00:00"},
		},
	},

	// Work day of month
	{
		"0 0 14W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-03-31 00:00:00", "Mon 2013-04-15 00:00"},
			{"2013-08-31 00:00:00", "Fri 2013-09-13 00:00"},
		},
	},

	// Work day of month -- end of month
	{
		"0 0 30W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-03-02 00:00:00", "Fri 2013-03-29 00:00"},
			{"2013-06-02 00:00:00", "Fri 2013-06-28 00:00"},
			{"2013-09-02 00:00:00", "Mon 2013-09-30 00:00"},
			{"2013-11-02 00:00:00", "Fri 2013-11-29 00:00"},
		},
	},

	// Last day of month
	{
		"0 0 L * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-09-02 00:00:00", "Mon 2013-09-30 00:00"},
			{"2014-01-01 00:00:00", "Fri 2014-01-31 00:00"},
			{"2014-02-01 00:00:00", "Fri 2014-02-28 00:00"},
			{"2016-02-15 00:00:00", "Mon 2016-02-29 00:00"},
		},
	},

	// Last work day of month
	{
		"0 0 LW * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-09-02 00:00:00", "Mon 2013-09-30 00:00"},
			{"2013-11-02 00:00:00", "Fri 2013-11-29 00:00"},
			{"2014-08-15 00:00:00", "Fri 2014-08-29 00:00"},
		},
	},

	// TODO: more tests
}

var cronbackwardtests = []crontest{
	// Seconds
	{
		"* * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:01", "2013-01-01 00:00:00"},
			{"2013-01-01 00:01:00", "2013-01-01 00:00:59"},
			{"2013-01-01 01:00:00", "2013-01-01 00:59:59"},
			{"2013-01-02 00:00:00", "2013-01-01 23:59:59"},
			{"2013-03-01 00:00:00", "2013-02-28 23:59:59"},
			{"2016-02-29 00:00:00", "2016-02-28 23:59:59"},
			{"2013-01-01 00:00:00", "2012-12-31 23:59:59"},
		},
	},

	// every 5 Second
	{
		"*/5 * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:06", "2013-01-01 00:00:05"},
			{"2013-01-01 00:01:00", "2013-01-01 00:00:55"},
			{"2013-01-01 01:00:00", "2013-01-01 00:59:55"},
			{"2013-01-02 00:00:00", "2013-01-01 23:59:55"},
			{"2013-03-01 00:00:00", "2013-02-28 23:59:55"},
			{"2016-02-29 00:00:00", "2016-02-28 23:59:55"},
			{"2013-01-01 00:00:00", "2012-12-31 23:59:55"},
		},
	},

	// Minutes
	{
		"* * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:58", "2013-01-01 00:00:00"},
			{"2013-01-01 00:01:00", "2013-01-01 00:00:00"},
			{"2013-01-01 01:00:00", "2013-01-01 00:59:00"},
			{"2013-01-02 00:00:00", "2013-01-01 23:59:00"},
			{"2013-03-01 00:00:00", "2013-02-28 23:59:00"},
			{"2016-02-29 00:00:00", "2016-02-28 23:59:00"},
			{"2013-01-01 00:00:00", "2012-12-31 23:59:00"},
		},
	},

	// // Minutes with interval
	{
		"17-43/5 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:17:01", "2013-01-01 00:17:00"},
			{"2013-01-01 00:33:00", "2013-01-01 00:32:00"},
			{"2013-01-01 01:00:00", "2013-01-01 00:42:00"},
			{"2013-01-02 00:01:00", "2013-01-01 23:42:00"},
			{"2013-03-01 00:01:00", "2013-02-28 23:42:00"},
			{"2016-02-29 00:01:00", "2016-02-28 23:42:00"},
			{"2013-01-01 00:01:00", "2012-12-31 23:42:00"},
		},
	},

	// Minutes interval, list
	{
		"15-30/4,55 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:16:00", "2013-01-01 00:15:00"},
			{"2013-01-01 00:18:59", "2013-01-01 00:15:00"},
			{"2013-01-01 00:19:00", "2013-01-01 00:15:00"},
			{"2013-01-01 00:56:00", "2013-01-01 00:55:00"},
			{"2013-01-01 01:15:00", "2013-01-01 00:55:00"},
			{"2013-01-02 00:15:00", "2013-01-01 23:55:00"},
			{"2013-03-01 00:15:00", "2013-02-28 23:55:00"},
			{"2016-02-29 00:15:00", "2016-02-28 23:55:00"},
			{"2012-12-31 23:54:00", "2012-12-31 23:27:00"},
			{"2013-01-01 00:15:00", "2012-12-31 23:55:00"},
		},
	},

	// Hour interval
	{
		"* 9-19/3 * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-01-01 00:10:00", "2016-12-31 18:59:00"},
			{"2017-02-01 00:10:00", "2017-01-31 18:59:00"},
			{"2017-02-12 00:10:00", "2017-02-11 18:59:00"},
			{"2017-02-12 19:10:00", "2017-02-12 18:59:00"},
			{"2017-02-12 12:15:00", "2017-02-12 12:14:00"},
			{"2017-02-12 13:00:00", "2017-02-12 12:59:00"},
			{"2017-02-12 11:00:00", "2017-02-12 09:59:00"},
		},
	},

	// Hour interval, list
	{
		"5 12-21/3,23 * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-01-01 00:10:00", "2016-12-31 23:05:00"},
			{"2017-02-01 00:10:00", "2017-01-31 23:05:00"},
			{"2017-02-12 00:10:00", "2017-02-11 23:05:00"},
			{"2017-02-12 19:10:00", "2017-02-12 18:05:00"},
			{"2017-02-12 12:15:00", "2017-02-12 12:05:00"},
			{"2017-02-12 22:00:00", "2017-02-12 21:05:00"},
		},
	},

	// Day interval
	{
		"5 10-17 12-25/4 * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-01-01 00:10:00", "2016-12-24 17:05:00"},
			{"2017-02-01 10:10:00", "2017-01-24 17:05:00"},
			{"2017-02-27 13:10:00", "2017-02-24 17:05:00"},
			{"2017-02-23 13:10:00", "2017-02-20 17:05:00"},
			{"2017-02-11 13:10:00", "2017-01-24 17:05:00"},
		},
	},

	// Day interval, list
	{
		"* * 12-15,20-22 * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-01-01 00:20:00", "2016-12-22 23:59:00"},
			{"2017-02-01 10:30:00", "2017-01-22 23:59:00"},
			{"2017-02-27 13:40:00", "2017-02-22 23:59:00"},
			{"2017-02-17 16:10:00", "2017-02-15 23:59:00"},
			{"2017-02-11 13:10:00", "2017-01-22 23:59:00"},
		},
	},

	// Month
	{
		"5 10 1 4-6 *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-01-01 00:10:00", "2016-06-01 10:05:00"},
			{"2017-07-01 10:01:00", "2017-06-01 10:05:00"},
			{"2017-06-03 00:10:00", "2017-06-01 10:05:00"},
		},
	},

	// Month
	{
		"0 0 0 12 * * 2017-2020",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2017-12-11 00:10:00", "2017-11-12 00:00:00"},
			{"2023-01-11 00:10:00", "2020-12-12 00:00:00"},
			{"2021-01-11 00:10:00", "2020-12-12 00:00:00"},
		},
	},

	// Days of week
	{
		"0 0 * * MON",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-10 00:00:00", "Mon 2013-01-07 00:00"},
			{"2017-08-07 00:00:00", "Mon 2017-07-31 00:00"},
			{"2017-01-01 00:30:00", "Mon 2016-12-26 00:00"},
		},
	},
	{
		"0 0 * * friday",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-08-14 00:00:00", "Fri 2017-08-11 00:00"},
			{"2017-08-02 00:00:00", "Fri 2017-07-28 00:00"},
			{"2018-01-02 00:30:00", "Fri 2017-12-29 00:00"},
		},
	},
	{
		"0 0 * * 6,7",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-09-04 00:00:00", "Sun 2017-09-03 00:00"},
			{"2017-08-02 00:00:00", "Sun 2017-07-30 00:00"},
			{"2018-01-03 00:30:00", "Sun 2017-12-31 00:00"},
		},
	},

	// // Specific days of week
	{
		"0 0 * * 6#5",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-03-03 00:00:00", "Sat 2016-12-31 00:00"},
		},
	},

	// // Work day of month
	{
		"0 0 18W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-12-02 00:00:00", "Fri 2017-11-17 00:00"},
			{"2017-10-12 00:00:00", "Mon 2017-09-18 00:00"},
			{"2017-08-30 00:00:00", "Fri 2017-08-18 00:00"},
			{"2017-06-21 00:00:00", "Mon 2017-06-19 00:00"},
		},
	},

	// // Work day of month -- end of month
	{
		"0 0 30W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-03-02 00:00:00", "Mon 2017-01-30 00:00"},
			{"2017-06-02 00:00:00", "Tue 2017-05-30 00:00"},
			{"2017-08-02 00:00:00", "Mon 2017-07-31 00:00"},
			{"2017-11-02 00:00:00", "Mon 2017-10-30 00:00"},
		},
	},

	// // Last day of month
	{
		"0 0 L * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2017-01-02 00:00:00", "Sat 2016-12-31 00:00"},
			{"2017-02-01 00:00:00", "Tue 2017-01-31 00:00"},
			{"2017-03-01 00:00:00", "Tue 2017-02-28 00:00"},
			{"2016-03-15 00:00:00", "Mon 2016-02-29 00:00"},
		},
	},

	// // Last work day of month
	{
		"0 0 LW * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2016-03-02 00:00:00", "Mon 2016-02-29 00:00"},
			{"2017-11-02 00:00:00", "Tue 2017-10-31 00:00"},
			{"2017-08-15 00:00:00", "Mon 2017-07-31 00:00"},
		},
	},
}

func TestExpressions(t *testing.T) {
	for _, test := range crontests {
		for _, times := range test.times {
			from, _ := time.Parse("2006-01-02 15:04:05", times.from)
			expr, err := cronexpr.Parse(test.expr)
			if err != nil {
				t.Errorf(`cronexpr.Parse("%s") returned "%s"`, test.expr, err.Error())
			}
			next := expr.Next(from)
			nextstr := next.Format(test.layout)
			if nextstr != times.next {
				t.Errorf(`("%s").Next("%s") = "%s", got "%s"`, test.expr, times.from, times.next, nextstr)
			}
		}
	}
}

func TestBackwardExpressions(t *testing.T) {
	for _, test := range cronbackwardtests {
		for _, times := range test.times {
			from, _ := time.Parse("2006-01-02 15:04:05", times.from)
			expr, err := cronexpr.Parse(test.expr)
			if err != nil {
				t.Errorf(`cronexpr.Parse("%s") returned "%s"`, test.expr, err.Error())
			}
			last := expr.Last(from)
			laststr := last.Format(test.layout)
			if laststr != times.next {
				t.Errorf(`("%s").Last("%s") = "%s", got "%s"`, test.expr, times.from, times.next, laststr)
			}
		}
	}
}

/******************************************************************************/

func TestZero(t *testing.T) {
	from, _ := time.Parse("2006-01-02", "2013-08-31")
	next := cronexpr.MustParse("* * * * * 1980").Next(from)
	if next.IsZero() == false {
		t.Error(`("* * * * * 1980").Next("2013-08-31").IsZero() returned 'false', expected 'true'`)
	}

	next = cronexpr.MustParse("* * * * * 2050").Next(from)
	if next.IsZero() == true {
		t.Error(`("* * * * * 2050").Next("2013-08-31").IsZero() returned 'true', expected 'false'`)
	}

	next = cronexpr.MustParse("* * * * * 2099").Next(time.Time{})
	if next.IsZero() == false {
		t.Error(`("* * * * * 2014").Next(time.Time{}).IsZero() returned 'true', expected 'false'`)
	}
}

/******************************************************************************/

func TestNextN(t *testing.T) {
	expected := []string{
		"Sat, 30 Nov 2013 00:00:00",
		"Sat, 29 Mar 2014 00:00:00",
		"Sat, 31 May 2014 00:00:00",
		"Sat, 30 Aug 2014 00:00:00",
		"Sat, 29 Nov 2014 00:00:00",
	}
	from, _ := time.Parse("2006-01-02 15:04:05", "2013-09-02 08:44:30")
	result := cronexpr.MustParse("0 0 * * 6#5").NextN(from, uint(len(expected)))
	if len(result) != len(expected) {
		t.Errorf(`MustParse("0 0 * * 6#5").NextN("2013-09-02 08:44:30", 5):\n"`)
		t.Errorf(`  Expected %d returned time values but got %d instead`, len(expected), len(result))
	}
	for i, next := range result {
		nextStr := next.Format("Mon, 2 Jan 2006 15:04:15")
		if nextStr != expected[i] {
			t.Errorf(`MustParse("0 0 * * 6#5").NextN("2013-09-02 08:44:30", 5):\n"`)
			t.Errorf(`  result[%d]: expected "%s" but got "%s"`, i, expected[i], nextStr)
		}
	}
}

func TestLastN(t *testing.T) {
	expected := []string{
		"Sat, 29 Nov 2014 00:00:00",
		"Sat, 30 Aug 2014 00:00:00",
		"Sat, 31 May 2014 00:00:00",
		"Sat, 29 Mar 2014 00:00:00",
		"Sat, 30 Nov 2013 00:00:00",
	}
	layout := "2006-01-02 15:04:05"
	ts := "2015-01-02 08:44:30"
	expr := "0 0 * * 6#5"
	from, _ := time.Parse(layout, ts)
	result := cronexpr.MustParse(expr).
		LastN(from, uint(len(expected)))
	if len(result) != len(expected) {
		t.Errorf(`MustParse("%s").LastN("%s", 5):\n"`, expr, ts)
		t.Errorf(`  Expected %d returned time values but got %d instead`, len(expected), len(result))
	}
	for i, next := range result {
		nextStr := next.Format("Mon, 2 Jan 2006 15:04:15")
		if nextStr != expected[i] {
			t.Errorf(`MustParse("%s").LastN("%s", 5):\n"`, expr, ts)
			t.Errorf(`  result[%d]: expected "%s" but got "%s"`, i, expected[i], nextStr)
		}
	}
}

func TestNextN_every5min(t *testing.T) {
	expected := []string{
		"Mon, 2 Sep 2013 08:45:00",
		"Mon, 2 Sep 2013 08:50:00",
		"Mon, 2 Sep 2013 08:55:00",
		"Mon, 2 Sep 2013 09:00:00",
		"Mon, 2 Sep 2013 09:05:00",
	}
	from, _ := time.Parse("2006-01-02 15:04:05", "2013-09-02 08:44:32")
	result := cronexpr.MustParse("*/5 * * * *").NextN(from, uint(len(expected)))
	if len(result) != len(expected) {
		t.Errorf(`MustParse("*/5 * * * *").NextN("2013-09-02 08:44:30", 5):\n"`)
		t.Errorf(`  Expected %d returned time values but got %d instead`, len(expected), len(result))
	}
	for i, next := range result {
		nextStr := next.Format("Mon, 2 Jan 2006 15:04:05")
		if nextStr != expected[i] {
			t.Errorf(`MustParse("*/5 * * * *").NextN("2013-09-02 08:44:30", 5):\n"`)
			t.Errorf(`  result[%d]: expected "%s" but got "%s"`, i, expected[i], nextStr)
		}
	}
}

func TestLastN_every5min(t *testing.T) {
	expected := []string{
		"Mon, 2 Sep 2013 09:05:00",
		"Mon, 2 Sep 2013 09:00:00",
		"Mon, 2 Sep 2013 08:55:00",
		"Mon, 2 Sep 2013 08:50:00",
		"Mon, 2 Sep 2013 08:45:00",
	}
	layout := "2006-01-02 15:04:05"
	ts := "2013-09-02 09:08:32"
	cron := "*/5 * * * *"
	from, _ := time.Parse(layout, ts)
	result := cronexpr.MustParse(cron).
		LastN(from, uint(len(expected)))
	if len(result) != len(expected) {
		t.Errorf(`MustParse("%s").LastN("%s", 5):\n"`, cron, ts)
		t.Errorf(`  Expected %d returned time values but got %d instead`, len(expected), len(result))
	}
	for i, next := range result {
		nextStr := next.Format("Mon, 2 Jan 2006 15:04:05")
		if nextStr != expected[i] {
			t.Errorf(`MustParse("%s").LastN("%s", 5):\n"`, cron, ts)
			t.Errorf(`  result[%d]: expected "%s" but got "%s"`, i, expected[i], nextStr)
		}
	}
}

// Issue: https://github.com/gorhill/cronexpr/issues/16
func TestInterval_Interval60Issue(t *testing.T) {
	_, err := cronexpr.Parse("*/60 * * * * *")
	if err == nil {
		t.Errorf("parsing with interval 60 should return err")
	}

	_, err = cronexpr.Parse("*/61 * * * * *")
	if err == nil {
		t.Errorf("parsing with interval 61 should return err")
	}

	_, err = cronexpr.Parse("2/60 * * * * *")
	if err == nil {
		t.Errorf("parsing with interval 60 should return err")
	}

	_, err = cronexpr.Parse("2-20/61 * * * * *")
	if err == nil {
		t.Errorf("parsing with interval 60 should return err")
	}
}

/******************************************************************************/

var benchmarkExpressions = []string{
	"* * * * *",
	"@hourly",
	"@weekly",
	"@yearly",
	"30 3 15W 3/3 *",
	"30 0 0 1-31/5 Oct-Dec * 2000,2006,2008,2013-2015",
	"0 0 0 * Feb-Nov/2 thu#3 2000-2050",
}
var benchmarkExpressionsLen = len(benchmarkExpressions)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cronexpr.MustParse(benchmarkExpressions[i%benchmarkExpressionsLen])
	}
}

func BenchmarkNext(b *testing.B) {
	exprs := make([]*cronexpr.Expression, benchmarkExpressionsLen)
	for i := 0; i < benchmarkExpressionsLen; i++ {
		exprs[i] = cronexpr.MustParse(benchmarkExpressions[i])
	}
	from := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr := exprs[i%benchmarkExpressionsLen]
		next := expr.Next(from)
		next = expr.Next(next)
		next = expr.Next(next)
		next = expr.Next(next)
		next = expr.Next(next)
	}
}
