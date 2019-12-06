package viewshow

import (
	"reflect"
	"testing"
)

func Test_getAllCinema(t *testing.T) {
	tests := []struct {
		name    string
		want    *Cinema
		wantErr bool
	}{
		{
			name: "Must include 1|TP",
			want: &Cinema{
				Name: "台北信義威秀影城",
				ID:   "1|TP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cinemas, err := getAllCinema()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllCinema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, c := range cinemas {
				if reflect.DeepEqual(c, tt.want) {
					return
				}

			}
			t.Errorf("getAllCinema() can not get %v", tt.want)
		})
	}
}
