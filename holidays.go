package holidays

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/wlbr/feiertage"
)

const dateForm = "01-02"

// Holidays defines a set of public holidays for a given territory.
type Holidays struct {
	// You can find ISO3166-2 codes on Wikipedia, for e.g. Austria:
	// https://en.wikipedia.org/wiki/ISO_3166-2:AT
	country  string // ISO3166-1 country code
	name     string // name of subdivision
	timezone string
	workdays []string
	Holidays []Holiday
}

// Holiday defines an individual public holiday.
type Holiday struct {
	Name    string
	Date    feiertage.Feiertag
	Yearday int
}

// CheckIsBusinessDay determines whether a given date is either at the weekend
// or a public holiday.
func CheckIsBusinessDay(hDate time.Time, holidays Holidays) bool {
	// Check to see if the given date is either at the weekend or a public
	// holiday. If they are, return the next date that isn't.
	wd := false
	day := hDate.Weekday().String()
	for _, d := range holidays.workdays {
		if day == d {
			wd = true
			break
		}
	}
	// check against the list of public holidays
	for holiday := range holidays.Holidays {
		if hDate.YearDay() == holidays.Holidays[holiday].Date.YearDay() {
			wd = false
		}
	}
	return wd
}

// GetFirstBusinessDay returns either the current date or, if that is not a
// working day, the next day that is.
func GetFirstBusinessDay(hDate time.Time, holidays Holidays) time.Time {
	// Short cut if it's today!
	if CheckIsBusinessDay(hDate, holidays) {
		return h_date
	}
	first := h_date.AddDate(0, 0, 1)
	for CheckIsBusinessDay(first, holidays) == false {
		first = first.AddDate(0, 0, 1)
	}
	return first
}

// GetHolidaysByYear returns a list of holidays for a given year.
func GetHolidaysByYear(year int) Holidays {
	var hols []Holiday
	austriaHolidays := feiertage.Österreich(year).Feiertage
	for _, h := range austriaHolidays {
		o := Holiday{h.Text, h, h.YearDay()}
		hols = append(hols, o)
	}

	localHolidays := Holidays{
		country:  "at",
		name:     "Austria/Vienna",
		workdays: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"},
		Holidays: hols,
	}
	return localHolidays
}

// GetHolidays returns a list of holidays for the current year.
func GetHolidays() Holidays {
	return GetHolidaysByYear(time.Now().Year())
}

// monthList converts a comma-separated list of months into an array.
func monthList(months string) []int {
	// Take a comma-separated list of months and convert them into an array.
	if months == "all" {
		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	}
	m := strings.Split(months, ",")
	var converted = []int{}
	for _, n := range m {
		monthno, _ := strconv.Atoi(n)
		converted = append(converted, monthno)
	}
	return converted
}

func parseDate(hDate string) time.Time {
	parsed, err := time.Parse(dateForm, hDate)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
