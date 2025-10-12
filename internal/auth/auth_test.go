package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr string
	}{
		"simple":      {input: http.Header{"Authorization": {"ApiKey testString"}}, want: "testString", wantErr: ""},
		"wrongKey":    {input: http.Header{"Authorization": {"WrongKey testString"}}, want: "", wantErr: "malformed authorization header"},
		"emptyHeader": {input: http.Header{}, want: "", wantErr: "no authorization header included"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			auth, err := GetAPIKey(tc.input)

			if err != nil && tc.wantErr != err.Error() {
				t.Errorf("GetAPIKey() error = %v, wantedErr = %v", err, tc.wantErr)
			}
			if tc.want != auth {
				t.Errorf("GetAPIKey() got = %v, expected = %v", auth, tc.want)
			}
		})
	}
}
