package date_test

import (
	"testing"

	"github.com/r0busta/go-date"
)

func TestParsePeriodFromStringSlice(t *testing.T) {
	t.Parallel()

	type args struct {
		period []string
	}

	tests := []struct {
		name    string
		args    args
		want    date.Range
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				period: []string{"2021-8-31", "2021-9-2"},
			},
			want: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 31),
				date.NewMfDateFromDate(2021, 9, 2),
			),
			wantErr: false,
		},
		{
			name: "same dates",
			args: args{
				period: []string{"2021-8-31", "2021-8-31"},
			},
			want: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 31),
				date.NewMfDateFromDate(2021, 8, 31),
			),
			wantErr: false,
		},
		{
			name: "incorrect args",
			args: args{
				period: []string{"2021-8-31"},
			},
			want:    date.Range{},
			wantErr: true,
		},
		{
			name: "no args",
			args: args{
				period: []string{},
			},
			want:    date.Range{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := date.ParsePeriodStringSlice(tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePeriodFromStringSlice() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !tt.wantErr && !got.Equals(tt.want) {
				t.Errorf("ParsePeriodFromStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Includes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		period date.Range
		arg    date.MfDate
		want   bool
	}{
		{
			name:   "empty",
			period: date.NewEmptyRange(),
			arg:    date.NewMfDate(),
			want:   false,
		},
		{
			name: "one day",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 8, 30),
			),
			arg:  date.NewMfDateFromDate(2021, 8, 30),
			want: true,
		},
		{
			name: "two days, include start",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 9, 1),
			),
			arg:  date.NewMfDateFromDate(2021, 8, 30),
			want: true,
		},
		{
			name: "two days, include end",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 9, 1),
			),
			arg:  date.NewMfDateFromDate(2021, 9, 1),
			want: true,
		},
		{
			name: "three days",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 9, 2),
			),
			arg:  date.NewMfDateFromDate(2021, 9, 1),
			want: true,
		},
		{
			name: "outside period, before start",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 9, 2),
			),
			arg:  date.NewMfDateFromDate(2021, 8, 29),
			want: false,
		},
		{
			name: "outside period, after start",
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 8, 30),
				date.NewMfDateFromDate(2021, 9, 2),
			),
			arg:  date.NewMfDateFromDate(2021, 9, 3),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.period.Includes(tt.arg); got != tt.want {
				t.Errorf("Range.Includes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_IsEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		period date.Range
		want   bool
	}{
		{
			want: true,
		},
		{
			period: date.Range{},
			want:   true,
		},
		{
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 10, 2),
				date.NewMfDateFromDate(2021, 10, 1),
			),
			want: true,
		},
		{
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 10, 2),
				date.NewMfDateFromDate(2021, 10, 2),
			),
			want: false,
		},
		{
			period: date.NewRangeFromDates(
				date.NewMfDateFromDate(2021, 10, 2),
				date.NewMfDateFromDate(2021, 10, 3),
			),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.period.IsEmpty(); got != tt.want {
				t.Errorf("Range.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
