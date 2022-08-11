package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"secrets_keeper/app/pkg/service"
	mock_service "secrets_keeper/app/pkg/service/mocks"
	"testing"
)

func TestHandler_SetMessage(t *testing.T) {
	type mockBehavior func(s *mock_service.MockKeeper, key, message string)

	testTable := []struct {
		name               string
		inputBody          string
		inputKey           string
		inputMessage       string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:         "OK",
			inputBody:    `{"key":"test", "message":"hello world!"}`,
			inputKey:     "test",
			inputMessage: "hello world!",
			mockBehavior: func(s *mock_service.MockKeeper, key, message string) {
				s.EXPECT().Set(key, message).Return(nil)
			},
			expectedStatusCode: 200,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			keeper := mock_service.NewMockKeeper(c)
			testCase.mockBehavior(keeper, testCase.inputKey, testCase.inputMessage)

			services := &service.Service{Keeper: keeper}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/", handler.SetMessage)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_GetMessage(t *testing.T) {
	type mockBehavior func(s *mock_service.MockKeeper, key string)

	testTable := []struct {
		name                string
		inputBody           string
		inputKey            string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `test`,
			inputKey:  "test",
			mockBehavior: func(s *mock_service.MockKeeper, key string) {
				s.EXPECT().Get(key).Return("hello world!", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "test string",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			keeper := mock_service.NewMockKeeper(c)
			testCase.mockBehavior(keeper, testCase.inputKey)

			services := &service.Service{Keeper: keeper}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/message/:key", handler.GetMessage)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/message/:key", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
