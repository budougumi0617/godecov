package godecov

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTotalsArray_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		subject   string
		in        string
		want      TotalsArray
		wantError bool
		err       error
	}{
		{
			subject: "simple",
			in:      "{\"totals\":[6,22,18,4,0,\"81.81818\",0,0,0,1,0,0,null]}",
			want: TotalsArray{
				Files:         6,
				Lines:         22,
				Hits:          18,
				Partials:      4,
				Missed:        0,
				CoverageRatio: "81.81818",
				Sessions:      1,
				Messages:      0,
				N:             0,
				Branches:      0,
				Methods:       0,
			},
			wantError: false,
			// err: fmt.Errorf(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {

			var got struct {
				Totals TotalsArray `json:"totals"`
			}
			err := json.Unmarshal([]byte(tt.in), &got)

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}
			if tt.wantError && !reflect.DeepEqual(err, tt.err) {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}

			if got.Totals != tt.want {
				t.Errorf("want %#v, but got %#v", tt.want, got.Totals)
			}
		})
	}
}

func TestTotalsArray_MarshalJSON(t *testing.T) {
	type In struct {
		Totals TotalsArray `json:"totals"`
	}
	tests := []struct {
		subject   string
		in        In
		want      string
		wantError bool
		err       error
	}{
		{
			subject: "simple",
			in: In{
				Totals: TotalsArray{
					Files:         6,
					Lines:         22,
					Hits:          18,
					Partials:      4,
					Missed:        0,
					CoverageRatio: "81.81818",
					Sessions:      1,
					Messages:      0,
					N:             0,
					Branches:      0,
					Methods:       0,
				},
			},
			want:      "{\"totals\":[6,22,18,4,0,\"81.81818\",0,0,0,1,0,0,null]}",
			wantError: false,
			// err: fmt.Errorf(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {

			got, err := json.Marshal(&tt.in)

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}
			if tt.wantError && !reflect.DeepEqual(err, tt.err) {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}

			if string(got) != tt.want {
				t.Errorf("want \n%s\nbut got \n%s", tt.want, string(got))
			}
		})
	}
}
