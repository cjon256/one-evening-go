package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

// Add a NewDateRange constructor that returns a new DateRange and an error.
// Return an error if the provided start date or end date is empty. You can use the IsZero() method from time.Time.
// Return an error if the provided end date happens before the start date. You can use the Before() method from time.Time.
// Rework the main function to use the constructor.

func NewDateRange(start, end time.Time) (DateRange, error) {
	if time.Time.IsZero(start) || time.Time.IsZero(end) {
		return DateRange{}, errors.New("Start or End is empty")
	}
	if end.Before(start) {
		return DateRange{}, errors.New("End is before Start")
	}
	return DateRange{
		Start: start,
		End:   end,
	}, nil
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	lifetime := DateRange{
		Start: time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
