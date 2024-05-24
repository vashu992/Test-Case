package pwg_test

import (
	"propwg/model"
	pwg "propwg/service"
	"reflect"
	"testing"
)

func TestGenerateAuthSession(t *testing.T) {
	// Define test cases
	tests := []struct {
		name       string
		authConfig pwg.AuthConfig
		want       model.AuthSession
	}{
		{
			name: validCase,
			authConfig: pwg.AuthConfig{
				ClecId:   "testClecId",
				UserName: "testUserName",
				Token:    "testToken",
				Pin:      "testPin",
			},
			want: model.AuthSession{
				AuthClec: model.AuthClec{
					ID: "testClecId",
					AuthAgentUser: model.AuthAgentUser{
						UserName: "testUserName",
						Token:    "testToken",
						Pin:      "testPin",
					},
				},
			},
		},
		// Add more test cases as needed
	}

	// Iterate over test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function under test
			got := pwg.GenerateAuthSession(tc.authConfig)

			// Compare the actual result with the expected result
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GenerateAuthSession() = %v, want %v", got, tc.want)
			}
		})
	}
}
