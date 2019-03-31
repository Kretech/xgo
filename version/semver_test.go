package version

import (
	"testing"

	"github.com/Kretech/xgo/test"
)

func TestSemVer_String(t *testing.T) {
	type fields struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   `test`,
			fields: fields{`v1.2.3`},
			want:   `v1.2.3`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Parse(tt.fields.s)
			if got := v.String(); got != tt.want {
				t.Errorf("SemVer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSemVer_Next(t *testing.T) {
	as := test.A(t)

	v := Parse(`1.2.3`)

	as.Equal(v.NextMajor().NumberString(), `2.0.0`)
	as.Equal(v.NextMinor().NumberString(), `1.3.0`)
	as.Equal(v.NextPatch().NumberString(), `1.2.4`)
}
