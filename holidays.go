package holidays

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/wlbr/feiertage"
)

const dateForm = "01-02"

type Holidays struct {
	// You can find ISO3166-2 codes on Wikipedia, for e.g. Austria:
	// https://en.wikipedia.org/wiki/ISO_3166-2:AT
	country  string // ISO3166-1 country code
	name     string // name of subdivision
	timezone string
	workdays []string
	Holidays []Holiday
}

type Holiday struct {
	Name    string
	Date    feiertage.Feiertag
	Yearday int
}

func CheckIsBusinessDay(h_date time.Time, holidays Holidays) bool {
	// Check to see if the given date is either at the weekend or a public
	// holiday. If they are, return the next date that isn't.
	wd := false
	day := h_date.Weekday().String()
	for _, d := range holidays.workdays {
		if day == d {
			wd = true
			break
		}
	}
	// check against the list of public holidays
	for holiday := range holidays.Holidays {
		if h_date.YearDay() == holidays.Holidays[holiday].Date.YearDay() {
			wd = false
		}
	}
	return wd
}

func GetFirstBusinessDay(h_date time.Time, holidays Holidays) time.Time {
	// Short cut if it's today!
	if CheckIsBusinessDay(h_date, holidays) {
		return h_date
	}
	first := h_date.AddDate(0, 0, 1)
	for CheckIsBusinessDay(first, holidays) == false {
		first = first.AddDate(0, 0, 1)
	}
	return first
}

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

func GetHolidays() Holidays {
	return GetHolidaysByYear(time.Now().Year())
}

func MonthList(months string) []int {
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

func parseDate(h_date string) time.Time {
	parsed, err := time.Parse(dateForm, h_date)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
