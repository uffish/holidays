package germany

import (
	"github.com/uffish/holidays"
	"github.com/wlbr/feiertage"
	"time"
)

func GetHolidaysByYear(year int) holidays.Holidays {
	var hols []holidays.Holiday
	germanyHolidays := feiertage.Deutschland(year).Feiertage
	for _, h := range germanyHolidays {
		o := holidays.Holiday{h.Time, h.Text, false, false}
		hols = append(hols, o)
	}

	localHolidays := holidays.Holidays{
		Country:  "de",
		Name:     "Germany/Berlin",
		Workdays: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"},
		Holidays: hols,
	}
	return localHolidays
}

func GetHolidays() holidays.Holidays {
	return GetHolidaysByYear(time.Now().Year())
}
