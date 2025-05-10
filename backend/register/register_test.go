package register

import (
	"bytes"
	"encoding/json"
	"login_register/models"
	"login_register/register"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func reset() {
	models.User_Registers = make(map[string]models.User_Register)
}

func init() {
	gin.SetMode(gin.TestMode)
}

func TestRegister(t *testing.T) {
	r := gin.Default()
	r.POST("/register", register.Register)

	tests := []struct {
		name               string
		input              models.User_Register
		setupUserRegisters func()
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name: "Missing details",
			input: models.User_Register{
				Username: "",
				Password: "password123",
				Email:    "user@example.com",
			},
			setupUserRegisters: reset,
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Please fill all the details",
		},
		{
			name: "User already exists",
			input: models.User_Register{
				Username: "existingUser",
				Password: "password123",
				Email:    "newuser@example.com",
			},
			setupUserRegisters: func() {
				reset()
				models.User_Registers["existingUser"] = models.User_Register{
					Username: "existingUser",
					Password: "password123",
					Email:    "existing@example.com",
				}
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "User already exists",
		},
		{
			name: "Email already exists",
			input: models.User_Register{
				Username: "newUser",
				Password: "password123",
				Email:    "existing@example.com",
			},
			setupUserRegisters: func() {
				reset()
				models.User_Registers["existingUser"] = models.User_Register{
					Username: "existingUser",
					Password: "password123",
					Email:    "existing@example.com",
				}
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Email already exists",
		},
		{
			name: "Successful registration",
			input: models.User_Register{
				Username: "newUser",
				Password: "password123",
				Email:    "newuser@example.com",
			},
			setupUserRegisters: reset,
			expectedStatusCode: http.StatusOK,
			expectedMessage:    "User registered successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupUserRegisters()

			data, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("could not marshal input data: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(data))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var response map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("could not unmarshal response: %v", err)
			}

			assert.Equal(t, tt.expectedMessage, response["message"])
		})
	}
}
