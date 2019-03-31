package version

import "testing"

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
