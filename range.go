package date

import (
	"fmt"
	"time"
)

const (
	periodLayout = "2006-1-2"
)

// Martin Fowler's date range implementation https://martinfowler.com/eaaDev/Range.html
type Range struct {
	start MfDate
	end   MfDate
}

func NewEmptyRange() Range {
	return NewRangeFromDates(
		NewMfDateFromDate(1970, 1, 3),
		NewMfDateFromDate(1970, 1, 2),
	)
}

func NewRangeFromDates(start MfDate, end MfDate) Range {
	r := Range{
		start: start,
		end:   end,
	}

	return r
}

func (r Range) IsEmpty() bool {
	if r.start.IsZero() && r.end.IsZero() {
		return true
	}

	return r.start.After(r.end)
}

func (r Range) Includes(arg MfDate) bool {
	return !arg.Before(r.start) && !arg.After(r.end)
}

func (r Range) Equals(arg Range) bool {
	return r.start.Equals(arg.start) && r.end.Equals(arg.end)
}

func (r Range) StartMonth() int {
	return int(r.start.getTime().Month())
}

func (r Range) EndMonth() int {
	return int(r.end.getTime().Month())
}

func (r Range) Start() MfDate {
	return r.start
}

func (r Range) End() MfDate {
	return r.end
}

func (r Range) String() string {
	return fmt.Sprintf("%s - %s", r.start, r.end)
}

func ParsePeriodString(from string, to string) (Range, error) {
	start, err := time.Parse(periodLayout, from)
	if err != nil {
		return Range{}, fmt.Errorf("error parsing `from` date: %w", err)
	}

	end, err := time.Parse(periodLayout, to)
	if err != nil {
		return Range{}, fmt.Errorf("error parsing `to` date: %w", err)
	}

	return NewRangeFromDates(NewMfDateFromTime(start), NewMfDateFromTime(end)), nil
}

func ParsePeriodStringSlice(period []string) (Range, error) {
	if len(period) != 2 {
		return Range{}, fmt.Errorf("expected `from` and `to` period dates")
	}

	return ParsePeriodString(period[0], period[1])
}
