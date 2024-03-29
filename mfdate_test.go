package date_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/r0busta/go-date"
)

func TestMfDate_YearMonthEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		arg  date.MfDate
		want bool
	}{
		{
			date: date.NewMfDateFromTime(time.Time{}),
			arg:  date.NewMfDateFromTime(time.Time{}),
			want: true,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: true,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 9, 1),
			want: false,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 7, 1),
			want: false,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2020, 8, 31),
			want: false,
		},
		{
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2022, 8, 31),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.YearMonthEqual(tt.arg); got != tt.want {
				t.Errorf("date.MfDate.YearMonthEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_DayBeforeOrEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		arg  date.MfDate
		want bool
	}{
		{
			name: "the same day",
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: true,
		},
		{
			name: "one day before",
			date: date.NewMfDateFromDate(2021, 7, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
		{
			name: "one month before",
			date: date.NewMfDateFromDate(2021, 7, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
		{
			name: "one year before",
			date: date.NewMfDateFromDate(2020, 8, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
		{
			name: "one day after",
			date: date.NewMfDateFromDate(2021, 9, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.DayBeforeOrEquals(tt.arg); got != tt.want {
				t.Errorf("date.MfDate.DayBeforeOrEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_DayBefore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		arg  date.MfDate
		want bool
	}{
		{
			name: "one day after",
			date: date.NewMfDateFromDate(2021, 9, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: false,
		},
		{
			name: "the same day",
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: false,
		},
		{
			name: "one day before and an earlier month",
			date: date.NewMfDateFromDate(2021, 7, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
		{
			name: "one month before",
			date: date.NewMfDateFromDate(2021, 7, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: true,
		},
		{
			name: "one year before",
			date: date.NewMfDateFromDate(2020, 8, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.DayBefore(tt.arg); got != tt.want {
				t.Errorf("date.MfDate.DayBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_Equals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		arg  date.MfDate
		want bool
	}{
		{
			name: "the same day",
			date: date.NewMfDateFromDate(2021, 8, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: true,
		},
		{
			name: "one day after",
			date: date.NewMfDateFromDate(2021, 9, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: false,
		},
		{
			name: "one day before and an earlier month",
			date: date.NewMfDateFromDate(2021, 7, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: false,
		},
		{
			name: "one month before",
			date: date.NewMfDateFromDate(2021, 7, 31),
			arg:  date.NewMfDateFromDate(2021, 8, 31),
			want: false,
		},
		{
			name: "one year before",
			date: date.NewMfDateFromDate(2020, 8, 1),
			arg:  date.NewMfDateFromDate(2021, 8, 1),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.Equals(tt.arg); got != tt.want {
				t.Errorf("date.MfDate.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_AddMonthResetDay(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		arg  int
		want date.MfDate
	}{
		{
			date: date.NewMfDateFromDate(2020, 8, 1),
			arg:  0,
			want: date.NewMfDateFromDate(2020, 8, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 8, 1),
			arg:  1,
			want: date.NewMfDateFromDate(2020, 9, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 12, 31),
			arg:  1,
			want: date.NewMfDateFromDate(2021, 1, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 10, 31),
			arg:  1,
			want: date.NewMfDateFromDate(2020, 11, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 10, 31),
			arg:  -1,
			want: date.NewMfDateFromDate(2020, 9, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 1, 31),
			arg:  -12,
			want: date.NewMfDateFromDate(2019, 1, 1),
		},
		{
			date: date.NewMfDateFromDate(2020, 1, 31),
			arg:  -13,
			want: date.NewMfDateFromDate(2018, 12, 1),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.AddMonthResetDay(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MfDate.AddMonthResetDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_IsZero(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		want bool
	}{
		{
			want: true,
		},
		{
			date: date.NewMfDateFromTime(time.Time{}),
			want: true,
		},
		{
			date: date.NewMfDate(),
			want: false,
		},
		{
			date: date.NewMfDateFromDate(2021, 10, 2),
			want: false,
		},
		{
			date: date.NewMfDateFromTime(time.Now()),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.IsZero(); got != tt.want {
				t.Errorf("MfDate.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		date date.MfDate
		want string
	}{
		{
			name: "NewMfDate creates date from the current time",
			date: date.NewMfDate(),
			want: date.NewMfDateFromTime(time.Now()).String(),
		},
		{
			date: date.NewMfDateFromDate(2021, 10, 4),
			want: "Oct 4, 2021",
		},
		{
			date: date.NewMfDateFromTime(time.Time{}),
			want: "Jan 1, 0001",
		},
		{
			date: date.NewMfDateFromTime(time.Date(2021, 10, 4, 0, 0, 0, 0, time.UTC)),
			want: "Oct 4, 2021",
		},
		{
			date: date.NewMfDateFromTime(time.Date(2021, 10, 4, 23, 59, 59, 0, time.UTC)),
			want: "Oct 4, 2021",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.String(); got != tt.want {
				t.Errorf("MfDate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMfDate_Format(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		date   date.MfDate
		layout string
		want   string
	}{
		{
			date:   date.NewMfDateFromDate(2021, 10, 4),
			layout: time.RFC822,
			want:   "04 Oct 21 00:00 UTC",
		},
		{
			date:   date.NewMfDateFromDate(2021, 10, 4),
			layout: time.RFC3339Nano,
			want:   "2021-10-04T00:00:00Z",
		},
		{
			date:   date.NewMfDateFromTime(time.Time{}),
			layout: "2006-01-02T15:04:05Z",
			want:   "0001-01-01T00:00:00Z",
		},
		{
			date:   date.NewMfDateFromTime(time.Date(2021, 10, 4, 0, 0, 0, 0, time.UTC)),
			layout: "2006-01-02T15:04:05Z",
			want:   "2021-10-04T00:00:00Z",
		},
		{
			date:   date.NewMfDateFromTime(time.Date(2021, 10, 4, 23, 59, 59, 0, time.UTC)),
			layout: "2006-01-02T15:04:05Z",
			want:   "2021-10-04T00:00:00Z",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.date.Format(tt.layout); got != tt.want {
				t.Errorf("MfDate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
