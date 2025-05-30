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

func TestCreateSupplier(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockSupplierUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			reqBody: map[string]interface{}{
				"name":    "Test Supplier",
				"email":   "supplier@example.com",
				"address": "456 Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					CreateSupplier(gomock.Any()).
					DoAndReturn(func(supplier *model.Supplier) error {
						return nil
					})
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Missing Name Field",
			reqBody: map[string]interface{}{
				"email":   "supplier@example.com",
				"address": "456 Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Email Field",
			reqBody: map[string]interface{}{
				"name":    "Test Supplier",
				"address": "456 Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Address Field",
			reqBody: map[string]interface{}{
				"name":  "Test Supplier",
				"email": "supplier@example.com",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Email Format",
			reqBody: map[string]interface{}{
				"name":    "Test Supplier",
				"email":   "not-a-valid-email",
				"address": "456 Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				// No mock calls expected
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Usecase Error",
			reqBody: map[string]interface{}{
				"name":    "Test Supplier",
				"email":   "supplier@example.com",
				"address": "456 Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					CreateSupplier(gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"code": "internal_error",
				"message": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockSupplierUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewSupplierHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.POST("/suppliers", handler.CreateSupplier)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPost, "/suppliers", bytes.NewBuffer(body))
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

				if tt.expectedStatus == http.StatusBadRequest {
					assert.Contains(t, gotBody, "code", "Code key missing for BadRequest")
					assert.Contains(t, gotBody, "message", "Message key missing for BadRequest")
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
				assert.Contains(t, gotBody, "code", "Code key missing for BadRequest without specific body")
				assert.Contains(t, gotBody, "message", "Message key missing for BadRequest without specific body")
				assert.NotEmpty(t, gotBody["message"], "Error message should not be empty for BadRequest")
			}
		})
	}
}

func TestGetSupplier(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	supplier := &model.Supplier{
		ID:    uuid.New().String(),
		Name:  "Test Supplier",
		Email: "supplier@example.com",
	}

	// Test cases
	tests := []struct {
		name           string
		supplierID     string
		mockSetup      func(*mock.MockSupplierUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			supplierID: supplier.ID,
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					GetSupplier(supplier.ID).
					Return(supplier, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "Not Found",
			supplierID: "non-existent-id",
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					GetSupplier("non-existent-id").
					Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: gin.H{
				"code": "not_found",
				"message": "Supplier not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockSupplierUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewSupplierHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/suppliers/:id", handler.GetSupplier)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/suppliers/"+tt.supplierID, nil)

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

func TestUpdateSupplier(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		supplierID     string
		reqBody        map[string]interface{}
		mockSetup      func(*mock.MockSupplierUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			supplierID: "supplier-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Supplier",
				"email":   "updated@example.com",
				"address": "789 New Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					UpdateSupplier(gomock.Eq("supplier-id"), gomock.Any()).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Supplier updated successfully",
			},
		},
		{
			name:       "Validation Error - Missing Name",
			supplierID: "supplier-id",
			reqBody: map[string]interface{}{
				"email":   "updated@example.com",
				"address": "789 New Supplier Ave",
			},
			mockSetup:      func(mockUsecase *mock.MockSupplierUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "Validation Error - Invalid Email",
			supplierID: "supplier-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Supplier",
				"email":   "invalid-email-format",
				"address": "789 New Supplier Ave",
			},
			mockSetup:      func(mockUsecase *mock.MockSupplierUsecase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "Usecase Error",
			supplierID: "supplier-id",
			reqBody: map[string]interface{}{
				"name":    "Updated Supplier",
				"email":   "updated@example.com",
				"address": "789 New Supplier Ave",
			},
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					UpdateSupplier(gomock.Eq("supplier-id"), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"code": "internal_error",
				"message": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockSupplierUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewSupplierHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.PUT("/suppliers/:id", handler.UpdateSupplier)

			// Create request body
			body, _ := json.Marshal(tt.reqBody)
			req, _ := http.NewRequest(http.MethodPut, "/suppliers/"+tt.supplierID, bytes.NewBuffer(body))
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
					assert.Contains(t, gotBody, "code", "Code key missing for BadRequest")
					assert.Contains(t, gotBody, "message", "Message key missing for BadRequest")
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
				assert.Contains(t, gotBody, "code", "Code key missing for BadRequest without specific body")
				assert.Contains(t, gotBody, "message", "Message key missing for BadRequest without specific body")
				assert.NotEmpty(t, gotBody["message"], "Error message should not be empty for BadRequest")
			}
		})
	}
}

func TestDeleteSupplier(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		supplierID     string
		mockSetup      func(*mock.MockSupplierUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "Success",
			supplierID: "supplier-id",
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					DeleteSupplier(gomock.Eq("supplier-id")).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"message": "Supplier deleted successfully",
			},
		},
		{
			name:       "Usecase Error",
			supplierID: "supplier-id",
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					DeleteSupplier(gomock.Eq("supplier-id")).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"code": "internal_error",
				"message": "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockSupplierUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewSupplierHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.DELETE("/suppliers/:id", handler.DeleteSupplier)

			// Create request
			req, _ := http.NewRequest(http.MethodDelete, "/suppliers/"+tt.supplierID, nil)

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

func TestGetAllSuppliers(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup
	suppliers := []*model.Supplier{
		{
			ID:    uuid.New().String(),
			Name:  "Supplier 1",
			Email: "supplier1@example.com",
		},
		{
			ID:    uuid.New().String(),
			Name:  "Supplier 2",
			Email: "supplier2@example.com",
		},
	}

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func(*mock.MockSupplierUsecase)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "Success",
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					GetAllSuppliers().
					Return(suppliers, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase Error",
			mockSetup: func(mockUsecase *mock.MockSupplierUsecase) {
				mockUsecase.EXPECT().
					GetAllSuppliers().
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"code": "internal_error",
				"message": "Failed to retrieve suppliers",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock usecase
			mockUsecase := mock.NewMockSupplierUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// Create handler with mock usecase
			handler := NewSupplierHandler(mockUsecase)

			// Create router and register handler
			router := gin.New()
			router.GET("/suppliers", handler.GetAllSuppliers)

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/suppliers", nil)

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
