package date

import (
	"time"
)

const mfDateStringLayout = "Jan 2, 2006"

// Martin Fowler's time point implementation https://martinfowler.com/eaaDev/TimePoint.html.
type MfDate struct {
	base *time.Time
}

func NewMfDate() MfDate {
	t := time.Now()

	return MfDate{
		base: &t,
	}
}

func NewMfDateFromTime(t time.Time) MfDate {
	d := MfDate{}
	d.initialize(t)

	return d
}

func NewMfDateFromDate(year int, month int, day int) MfDate {
	return NewMfDateFromTime(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func (d *MfDate) initialize(t time.Time) {
	d.base = trimToDays(t)
}

func trimToDays(t time.Time) *time.Time {
	res := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())

	return &res
}

func (d MfDate) IsZero() bool {
	return d.base == nil || d.base.IsZero()
}

func (d MfDate) Year() int {
	return d.base.Year()
}

func (d MfDate) Month() int {
	return int(d.base.Month())
}

func (d MfDate) Day() int {
	return d.base.Day()
}

func (d MfDate) Time() time.Time {
	return *d.base
}

func (d MfDate) After(arg MfDate) bool {
	return d.base.After(arg.Time())
}

func (d MfDate) Before(arg MfDate) bool {
	return d.base.Before(arg.Time())
}

func (d MfDate) Equals(arg MfDate) bool {
	return d.base.Equal(arg.Time())
}

func (d MfDate) AddMonthResetDay(mm int) MfDate {
	t := time.Date(d.Year(), time.Month(d.Month()+mm), 1, 0, 0, 0, 0, time.UTC)

	return NewMfDateFromTime(t)
}

func (d MfDate) YearMonthEqual(arg MfDate) bool {
	return d.Year() == arg.Year() && d.Month() == arg.Month()
}

func (d MfDate) DayBeforeOrEquals(arg MfDate) bool {
	if d.YearMonthEqual(arg) {
		return true
	}

	return d.DayBefore(arg)
}

func (d MfDate) DayBefore(arg MfDate) bool {
	if d.Year() < arg.Year() {
		return true
	}

	return d.Year() == arg.Year() && d.Month() < arg.Month()
}

func (d MfDate) String() string {
	return d.base.Format(mfDateStringLayout)
}
