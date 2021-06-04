package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetCreatedUserData(t *testing.T) {

	testCases := []struct {
		testName     string
		userID       string
		expectErr    bool
		expectedResp *Whole_body
	}{
		{
			testName:  "Get data of Existing user",
			userID:    "PQXBTF7",
			expectErr: false,
			expectedResp: &Whole_body{
				User: User{
					Name:            "THARUN Chunchu",
					Email:           "tharun@clevertap.com",
					Type:            "user",
					Id:              "PQXBTF7",
					Role:            "owner",
					Contact_methods: []Contact_methods{{Type: "email_contact_method_reference", Summary: "Default"}},
				},
			},
		},
		{
			testName:     "Get data of non-Existing user",
			userID:       "07686",
			expectErr:    true,
			expectedResp: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			authToken := os.Getenv("token")

			client := NewClient(authToken)

			user, err := client.GetUser(tc.userID)

			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_CreateUser(t *testing.T) {

	testCases := []struct {
		testName     string
		newItem      Whole_body
		expectedResp *Whole_body
		expectErr    bool
	}{
		{
			testName: "create new user",
			newItem: Whole_body{
				User: User{
					Name:  "tharun",
					Email: "tharun@gmail.com",
					Type:  "user",
					Role:  "admin",
				},
			},

			expectedResp: &Whole_body{
				User: User{Name: "tharun",
					Email:           "tharun@gmail.com",
					Type:            "user",
					Role:            "admin",
					Contact_methods: []Contact_methods{{Type: "email_contact_method_reference", Summary: "Default"}},
				},
			},

			expectErr: false,
		},
		{
			testName: "create existing user",
			newItem: Whole_body{
				User: User{
					Name:  "tharun",
					Email: "tharun@clevertap.com",
					Type:  "user",
					Role:  "admin",
				},
			},

			expectedResp: nil,

			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			authToken := os.Getenv("token")
			client := NewClient(authToken)
			re, err := client.CreateUser(tc.newItem)

			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			tc.expectedResp.User.Id = re.User.Id
			user, err := client.GetUser(re.User.Id)
			client.DeleteUser(re.User.Id)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedResp, user)

		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName     string
		updatedUser  Whole_body
		expectedResp *Whole_body
		userID       string
		expectErr    bool
	}{
		{
			testName: "Update existing user",
			updatedUser: Whole_body{
				User: User{
					Name:  "Tharunafterupdate",
					Type:  "user",
					Email: "tharun@123test.com",
					Role:  "admin",
				},
			},
			expectedResp: &Whole_body{
				User: User{
					Name:            "Tharunafterupdate",
					Type:            "user",
					Email:           "tharun@123test.com",
					Role:            "admin",
					Contact_methods: []Contact_methods{{Type: "email_contact_method_reference", Summary: "Default"}},
					Id:              "PMQFYGT",
				},
			},
			expectErr: false,
			userID:    "PMQFYGT",
		},
		{
			testName: "Update  non existing user",
			updatedUser: Whole_body{
				User: User{
					Name:  "Tharunafterupdate",
					Type:  "user",
					Email: "tharun@clevertap.com",
					Role:  "owner",
				},
			},
			expectedResp: &Whole_body{
				User: User{
					Name:            "Tharunafterupdate",
					Type:            "user",
					Email:           "tharun@clevertap.com",
					Role:            "owner",
					Contact_methods: []Contact_methods{{Type: "email_contact_method_reference", Summary: "Default"}},
				},
			},
			expectErr: true,
			userID:    "fvsfmvosm",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			authToken := os.Getenv("token")
			client := NewClient(authToken)
			_, err := client.UpdateUser(tc.updatedUser, tc.userID)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			tc.expectedResp.User.Id = tc.userID
			user, err := client.GetUser(tc.userID)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	testCases := []struct {
		testName  string
		new_user  Whole_body
		expectErr bool
		userID    string
	}{
		{
			testName: "Delete user",
			new_user: Whole_body{
				User: User{
					Name:  "Tharundeletetest",
					Type:  "user",
					Email: "tharundelete@test.com",
					Role:  "admin",
				},
			},
			expectErr: false,
		},
		{
			testName: "Delete non existing user",
			new_user: Whole_body{
				User: User{
					Name:  "Tharundeletetest",
					Type:  "user",
					Email: "tharundelete@test.com",
					Role:  "admin",
				},
			},

			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			authToken := os.Getenv("token")
			client := NewClient(authToken)
			tc.userID = "Random"
			if tc.testName == "Delete user" {
				re, _ := client.CreateUser(tc.new_user)
				tc.userID = re.User.Id
			}

			err := client.DeleteUser(tc.userID)

			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			_, err = client.GetUser(tc.userID)
			assert.Error(t, err)
		})
	}
}
