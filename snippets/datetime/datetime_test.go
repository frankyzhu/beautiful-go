package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestToTime(t *testing.T) {
	type args struct {
		unixMS int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "ToUTCTime",
			args: args{
				unixMS: 946684800000,
			},
			want: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUTCTime(tt.args.unixMS); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUTCTime() = %v, want %v, nano %v", got, tt.want, tt.want.UnixNano())
			}
		})
	}
}
