package auth

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetAPIKeyErr(t *testing.T) {
	testheader := http.Header{}
	testheader.Add("Authorization", "")
	_, err := GetAPIKey(testheader)
	if err == nil {
		t.Fatalf("ERR")
	}
}

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		key   string
		err   bool
	}{
		"Test Empty Value": {
			input: func() http.Header {
				header := http.Header{}
				header.Add("Authorization", "something")
				return header
			}(),
			key: "",
			err: true,
		},
		"Test Auth Key": {
			input: func() http.Header {
				header := http.Header{}
				header.Add("Authorization", "ApiKey 123abc")
				return header
			}(),
			key: "123abc",
			err: false,
		},
		"Purposefully Failing Test": {
			input: func() http.Header {
				header := http.Header{}
				header.Add("Authorization", "APIKEY 123abc")
				return header
			}(),
			key: "123abc",
			err: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//t.Fatalf("%v", tc.input.Get("Authorization"))
			key, err := GetAPIKey(tc.input)
			if tc.err && err == nil {
				t.Fatalf("%v did not error when expected", name)
			} else if !tc.err && err != nil {
				t.Fatalf("%v has unexpected error %v", name, err)
			} else {
				keydiff := cmp.Diff(tc.key, key)
				if keydiff != "" {
					t.Fatalf("Key Mismatch: %v", keydiff)
				}
			}
		})
	}
}
