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
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerCarCreate(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			reqBody: map[string]interface{}{
				"car_id":      "car123",
				"customer_id": "cust123",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					CreateCustomerCar(gomock.Any()).
					DoAndReturn(func(customerCar *model.CustomerCar) error {
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Missing CarID Field",
			reqBody: map[string]interface{}{
				"customer_id": "cust123",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing CustomerID Field",
			reqBody: map[string]interface{}{
				"car_id": "car123",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Usecase Error",
			reqBody: map[string]interface{}{
				"car_id":      "car123",
				"customer_id": "cust123",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					CreateCustomerCar(gomock.Any()).
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
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.POST("/customer-cars", handler.Create)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/customer-cars", bytes.NewBuffer(body))
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

func TestCustomerCarGetByID(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	customerCar := &model.CustomerCar{
		ID:         "cc123",
		CarID:      "car123",
		CustomerID: "cust123",
	}

	// Test cases
	tests := []struct {
		name           string
		customerCarID  string
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:          "Success",
			customerCarID: customerCar.ID,
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCar(customerCar.ID).
					Return(customerCar, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:          "Not Found",
			customerCarID: "non-existent-id",
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCar("non-existent-id").
					Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: gin.H{
				"error": "Customer car relationship not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/customer-cars/:id", handler.GetByID)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/customer-cars/"+tt.customerCarID, nil)

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

func TestCustomerCarGetAll(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: "car123", CustomerID: "cust123"},
		{ID: "cc456", CarID: "car456", CustomerID: "cust456"},
	}

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetAllCustomerCars().
					Return(customerCars, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase Error",
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetAllCustomerCars().
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "Failed to get customer car relationships",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/customer-cars", handler.GetAll)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/customer-cars", nil)

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

func TestCustomerCarGetByCustomerID(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	customerID := "cust123"
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: "car123", CustomerID: customerID},
		{ID: "cc456", CarID: "car456", CustomerID: customerID},
	}

	// Test cases
	tests := []struct {
		name           string
		customerID     string
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			customerID: customerID,
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCarsByCustomerID(customerID).
					Return(customerCars, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "Usecase Error",
			customerID: customerID,
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCarsByCustomerID(customerID).
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "Failed to get customer car relationships",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/customers/:customer_id/cars", handler.GetByCustomerID)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/customers/"+tt.customerID+"/cars", nil)

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

func TestCustomerCarGetByCarID(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	carID := "car123"
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: carID, CustomerID: "cust123"},
		{ID: "cc456", CarID: carID, CustomerID: "cust456"},
	}

	// Test cases
	tests := []struct {
		name           string
		carID          string
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:  "Success",
			carID: carID,
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCarsByCarID(carID).
					Return(customerCars, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:  "Usecase Error",
			carID: carID,
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					GetCustomerCarsByCarID(carID).
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error": "Failed to get customer car relationships",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/cars/:car_id/customers", handler.GetByCarID)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/cars/"+tt.carID+"/customers", nil)

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

func TestCustomerCarUpdate(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		customerCarID  string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:          "Success",
			customerCarID: "cc123",
			reqBody: map[string]interface{}{
				"car_id":      "car456",
				"customer_id": "cust456",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					UpdateCustomerCar(gomock.Eq("cc123"), gomock.Any()).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Customer car relationship updated successfully",
			},
		},
		{
			name:          "Validation Error - Missing CarID",
			customerCarID: "cc123",
			reqBody: map[string]interface{}{
				"customer_id": "cust456",
			},
			mockSetup:      func(mockUsecase *mock.MockCustomerCarUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:          "Validation Error - Missing CustomerID",
			customerCarID: "cc123",
			reqBody: map[string]interface{}{
				"car_id": "car456",
			},
			mockSetup:      func(mockUsecase *mock.MockCustomerCarUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:          "Usecase Error",
			customerCarID: "cc123",
			reqBody: map[string]interface{}{
				"car_id":      "car456",
				"customer_id": "cust456",
			},
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					UpdateCustomerCar(gomock.Eq("cc123"), gomock.Any()).
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
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.PUT("/customer-cars/:id", handler.Update)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPut, "/customer-cars/"+tt.customerCarID, bytes.NewBuffer(body))
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

func TestCustomerCarDelete(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		customerCarID  string
		mockSetup      func(*mock.MockCustomerCarUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:          "Success",
			customerCarID: "cc123",
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					DeleteCustomerCar(gomock.Eq("cc123")).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Customer car relationship deleted successfully",
			},
		},
		{
			name:          "Usecase Error",
			customerCarID: "cc123",
			mockSetup: func(mockUsecase *mock.MockCustomerCarUsecase) {
				mockUsecase.EXPECT().
					DeleteCustomerCar(gomock.Eq("cc123")).
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
			mockUsecase := mock.NewMockCustomerCarUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewCustomerCarHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.DELETE("/customer-cars/:id", handler.Delete)

			// Create request
			req, _ := http.NewRequest(http.MethodDelete, "/customer-cars/"+tt.customerCarID, nil)

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
