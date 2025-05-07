package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockCustomerUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			reqBody: map[string]interface{}{
				"name":    "Test Customer",
				"email":   "test@example.com",
				"address": "123 Main St",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					CreateCustomer(gomock.Any()).
					DoAndReturn(func(customer *model.Customer) error {
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Missing Name Field",
			reqBody: map[string]interface{}{
				"email":   "test@example.com",
				"address": "123 Main St",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
			// Error message from validator can be complex, checking for status is often enough
			// expectedBody: gin.H{"error": "Key: 'Customer.Name' Error:Field validation for 'Name' failed on the 'required' tag"},
		},
		{
			name: "Missing Email Field",
			reqBody: map[string]interface{}{
				"name":    "Test Customer",
				"address": "123 Main St",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Address Field",
			reqBody: map[string]interface{}{
				"name":  "Test Customer",
				"email": "test@example.com",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Email Format",
			reqBody: map[string]interface{}{
				"name":    "Test Customer",
				"email":   "not-an-email",
				"address": "123 Main St",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Usecase Error",
			reqBody: map[string]interface{}{
				"name":    "Test Customer",
				"email":   "test@example.com",
				"address": "123 Main St",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					CreateCustomer(gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.POST("/customers", handler.CreateCustomer)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check response body if provided
			if tt.expectedBody != nil {
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)

				// For validation errors, the exact message from Gin can be complex.
				// We might only check if an "error" key exists for BadRequest.
				if tt.expectedStatus == http.StatusBadRequest {
					assert.Contains(t, gotBody, "error", "Error key missing for BadRequest")
				} else {
					expectedBodyBytes, _ := json.Marshal(tt.expectedBody)
					var expectedBodyMap map[string]interface{}
					json.Unmarshal(expectedBodyBytes, &expectedBodyMap)

					for k, v := range expectedBodyMap {
						assert.Equal(t, v, gotBody[k])
					}
				}
			} else if tt.expectedStatus == http.StatusBadRequest {
				// If it's a bad request and no specific body is expected,
				// ensure there's some error message.
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)
				assert.Contains(t, gotBody, "error", "Error key missing for BadRequest without specific body")
				assert.NotEmpty(t, gotBody["error"], "Error message should not be empty for BadRequest")
			}
		})
	}
}

func TestGetCustomer(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	customer := &model.Customer{
		ID:    uuid.New().String(),
		Name:  "Test Customer",
		Email: "test@example.com",
	}

	// Test cases
	tests := []struct {
		name           string
		customerID     string
		mockSetup      func(*mock.MockCustomerUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			customerID: customer.ID,
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					GetCustomer(customer.ID).
					Return(customer, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "Not Found",
			customerID: "non-existent-id",
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					GetCustomer("non-existent-id").
					Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: gin.H{
				"error": "Customer not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/customers/:id", handler.GetCustomer)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/customers/"+tt.customerID, nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check response body if provided
			if tt.expectedBody != nil {
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)

				expectedBodyBytes, _ := json.Marshal(tt.expectedBody)
				var expectedBodyMap map[string]interface{}
				json.Unmarshal(expectedBodyBytes, &expectedBodyMap)

				for k, v := range expectedBodyMap {
					assert.Equal(t, v, gotBody[k])
				}
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		customerID     string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockCustomerUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			customerID: "customer-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Customer",
				"email":   "updated@example.com",
				"address": "456 New Ave",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					UpdateCustomer(gomock.Eq("customer-id"), gomock.Any()).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Customer updated successfully",
			},
		},
		{
			name:       "Validation Error - Missing Name",
			customerID: "customer-id",
			reqBody: map[string]interface{}{
				"email":   "updated@example.com",
				"address": "456 New Ave",
			},
			mockSetup:      func(mockUsecase *mock.MockCustomerUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "Validation Error - Invalid Email",
			customerID: "customer-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Customer",
				"email":   "invalid-email",
				"address": "456 New Ave",
			},
			mockSetup:      func(mockUsecase *mock.MockCustomerUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "Usecase Error",
			customerID: "customer-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Customer",
				"email":   "updated@example.com",
				"address": "456 New Ave",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					UpdateCustomer(gomock.Eq("customer-id"), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.PUT("/customers/:id", handler.UpdateCustomer)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPut, "/customers/"+tt.customerID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check response body
			if tt.expectedBody != nil {
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)
				if tt.expectedStatus == http.StatusBadRequest {
					assert.Contains(t, gotBody, "error", "Error key missing for BadRequest")
				} else {
					expectedBodyBytes, _ := json.Marshal(tt.expectedBody)
					var expectedBodyMap map[string]interface{}
					json.Unmarshal(expectedBodyBytes, &expectedBodyMap)
					for k, v := range expectedBodyMap {
						assert.Equal(t, v, gotBody[k])
					}
				}
			} else if tt.expectedStatus == http.StatusBadRequest {
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)
				assert.Contains(t, gotBody, "error", "Error key missing for BadRequest without specific body")
				assert.NotEmpty(t, gotBody["error"], "Error message should not be empty for BadRequest")
			}
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		customerID     string
		mockSetup      func(*mock.MockCustomerUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			customerID: "customer-id",
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					DeleteCustomer(gomock.Eq("customer-id")).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Customer deleted successfully",
			},
		},
		{
			name:       "Usecase Error",
			customerID: "customer-id",
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					DeleteCustomer(gomock.Eq("customer-id")).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.DELETE("/customers/:id", handler.DeleteCustomer)

			// Create request
			req, _ := http.NewRequest(http.MethodDelete, "/customers/"+tt.customerID, nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check response body
			var gotBody map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &gotBody)

			expectedBodyBytes, _ := json.Marshal(tt.expectedBody)
			var expectedBodyMap map[string]interface{}
			json.Unmarshal(expectedBodyBytes, &expectedBodyMap)

			for k, v := range expectedBodyMap {
				assert.Equal(t, v, gotBody[k])
			}
		})
	}
}

func TestGetAllCustomers(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	customers := []*model.Customer{
		{
			ID:    uuid.New().String(),
			Name:  "Customer 1",
			Email: "customer1@example.com",
		},
		{
			ID:    uuid.New().String(),
			Name:  "Customer 2",
			Email: "customer2@example.com",
		},
	}

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(*mock.MockCustomerUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					GetAllCustomers().
					Return(customers, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase Error",
			mockSetup: func(mockUsecase *mock.MockCustomerUsecase) {
				mockUsecase.EXPECT().
					GetAllCustomers().
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "Failed to retrieve customers",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/customers", handler.GetAllCustomers)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check response body if provided
			if tt.expectedBody != nil {
				var gotBody map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &gotBody)

				expectedBodyBytes, _ := json.Marshal(tt.expectedBody)
				var expectedBodyMap map[string]interface{}
				json.Unmarshal(expectedBodyBytes, &expectedBodyMap)

				for k, v := range expectedBodyMap {
					assert.Equal(t, v, gotBody[k])
				}
			}
		})
	}
}
