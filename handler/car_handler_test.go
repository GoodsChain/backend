package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoodsChain/backend/mock" // Assuming mock package is at this path
	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/repository" // For repository.ErrNotFound
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"                 // Corrected import path
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupCarRouter(t *testing.T) (*gin.Engine, *mock.MockCarUsecase) {
	ctrl := gomock.NewController(t)
	mockUsecase := mock.NewMockCarUsecase(ctrl)
	carHandler := NewCarHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	// No global error middleware for these specific unit tests, or add it if its behavior is tested.
	// router.Use(ErrorHandlingMiddleware()) // If you want to test with it

	carRoutes := router.Group("/cars")
	{
		carRoutes.POST("/", carHandler.CreateCar)
		carRoutes.GET("/:id", carHandler.GetCar)
		carRoutes.GET("/", carHandler.GetAllCars)
		carRoutes.PUT("/:id", carHandler.UpdateCar)
		carRoutes.DELETE("/:id", carHandler.DeleteCar)
	}
	return router, mockUsecase
}

func TestCarHandler_CreateCar(t *testing.T) {
	router, mockUsecase := setupCarRouter(t)

	t.Run("Success", func(t *testing.T) {
		carInput := model.Car{Name: "New Car", SupplierID: "supp1", Price: 30000}
		carOutput := carInput
		carOutput.ID = uuid.New().String() // Usecase/Repo would set this

		mockUsecase.EXPECT().CreateCar(gomock.Any()).DoAndReturn(
			func(c *model.Car) error {
				// Simulate ID generation and CreatedBy/UpdatedBy if usecase does it
				c.ID = carOutput.ID
				// c.CreatedBy = "system"
				// c.UpdatedBy = "system"
				return nil
			}).Times(1)

		jsonValue, _ := json.Marshal(carInput)
		req, _ := http.NewRequest(http.MethodPost, "/cars/", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var resultCar model.Car
		err := json.Unmarshal(rr.Body.Bytes(), &resultCar)
		assert.NoError(t, err)
		assert.Equal(t, carOutput.ID, resultCar.ID) // Check if ID is returned
		assert.Equal(t, carInput.Name, resultCar.Name)
	})

	t.Run("BindError", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/cars/", bytes.NewBufferString(`{"name": "Test", "price": "notanumber"}`)) // Invalid price type
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		// Check error response if needed
	})

	t.Run("UsecaseError", func(t *testing.T) {
		carInput := model.Car{Name: "Error Car", SupplierID: "supp_err", Price: 1}
		mockUsecase.EXPECT().CreateCar(gomock.Any()).Return(errors.New("usecase create error")).Times(1)

		jsonValue, _ := json.Marshal(carInput)
		req, _ := http.NewRequest(http.MethodPost, "/cars/", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		// Check error response
	})
}

func TestCarHandler_GetCar(t *testing.T) {
	router, mockUsecase := setupCarRouter(t)
	carID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		expectedCar := &model.Car{ID: carID, Name: "Fetched Car"}
		mockUsecase.EXPECT().GetCar(carID).Return(expectedCar, nil).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resultCar model.Car
		err := json.Unmarshal(rr.Body.Bytes(), &resultCar)
		assert.NoError(t, err)
		assert.Equal(t, expectedCar.Name, resultCar.Name)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockUsecase.EXPECT().GetCar(carID).Return(nil, repository.ErrNotFound).Times(1)
		req, _ := http.NewRequest(http.MethodGet, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("UsecaseError", func(t *testing.T) {
		mockUsecase.EXPECT().GetCar(carID).Return(nil, errors.New("some other error")).Times(1)
		req, _ := http.NewRequest(http.MethodGet, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestCarHandler_GetAllCars(t *testing.T) {
	router, mockUsecase := setupCarRouter(t)

	t.Run("Success", func(t *testing.T) {
		expectedCars := []model.Car{
			{ID: uuid.New().String(), Name: "Car A"},
			{ID: uuid.New().String(), Name: "Car B"},
		}
		mockUsecase.EXPECT().GetAllCars().Return(expectedCars, nil).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/cars/", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resultCars []model.Car
		err := json.Unmarshal(rr.Body.Bytes(), &resultCars)
		assert.NoError(t, err)
		assert.Equal(t, expectedCars, resultCars)
	})

	t.Run("SuccessEmpty", func(t *testing.T) {
		mockUsecase.EXPECT().GetAllCars().Return([]model.Car{}, nil).Times(1)
		req, _ := http.NewRequest(http.MethodGet, "/cars/", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		var resultCars []model.Car
		err := json.Unmarshal(rr.Body.Bytes(), &resultCars)
		assert.NoError(t, err)
		assert.Empty(t, resultCars)
	})
	
	t.Run("UsecaseError", func(t *testing.T) {
		mockUsecase.EXPECT().GetAllCars().Return(nil, errors.New("failed to fetch")).Times(1)
		req, _ := http.NewRequest(http.MethodGet, "/cars/", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}


func TestCarHandler_UpdateCar(t *testing.T) {
	router, mockUsecase := setupCarRouter(t)
	carID := uuid.New().String()
	carInput := model.Car{Name: "Updated Car", SupplierID: "supp_upd", Price: 35000}

	t.Run("Success", func(t *testing.T) {
		mockUsecase.EXPECT().UpdateCar(carID, gomock.Any()).Return(nil).Times(1)
		
		jsonValue, _ := json.Marshal(carInput)
		req, _ := http.NewRequest(http.MethodPut, "/cars/"+carID, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp model.SuccessResponse
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Car updated successfully", resp.Message)
	})

	t.Run("BindError", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/cars/"+carID, bytes.NewBufferString(`{"name": "Test", "price": "notanumber"}`))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockUsecase.EXPECT().UpdateCar(carID, gomock.Any()).Return(repository.ErrNotFound).Times(1)
		jsonValue, _ := json.Marshal(carInput)
		req, _ := http.NewRequest(http.MethodPut, "/cars/"+carID, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	
	t.Run("UsecaseError", func(t *testing.T) {
		mockUsecase.EXPECT().UpdateCar(carID, gomock.Any()).Return(errors.New("update failed badly")).Times(1)
		jsonValue, _ := json.Marshal(carInput)
		req, _ := http.NewRequest(http.MethodPut, "/cars/"+carID, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestCarHandler_DeleteCar(t *testing.T) {
	router, mockUsecase := setupCarRouter(t)
	carID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		mockUsecase.EXPECT().DeleteCar(carID).Return(nil).Times(1)
		req, _ := http.NewRequest(http.MethodDelete, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		var resp model.SuccessResponse
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Car deleted successfully", resp.Message)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockUsecase.EXPECT().DeleteCar(carID).Return(repository.ErrNotFound).Times(1)
		req, _ := http.NewRequest(http.MethodDelete, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("UsecaseError", func(t *testing.T) {
		mockUsecase.EXPECT().DeleteCar(carID).Return(errors.New("delete failed badly")).Times(1)
		req, _ := http.NewRequest(http.MethodDelete, "/cars/"+carID, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
