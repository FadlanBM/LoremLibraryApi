package helper

import (
	"fmt"
	"time"
)

func FormatDate(date time.Time) string {
	loc := time.FixedZone("WIB", 7*60*60)

	date = date.In(loc)

	months := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	day := date.Day()
	month := date.Month()
	year := date.Year()

	formattedDate := fmt.Sprintf("%d %s %d", day, months[month-1], year)

	return formattedDate
}
