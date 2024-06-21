package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stu/controllers"
	"stu/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSchoolService is a mock implementation of SchoolService for testing purposes.
type MockSchoolService struct {
	mock.Mock
}

func (m *MockSchoolService) GetAllSchools() ([]models.School, error) {
	args := m.Called()
	return args.Get(0).([]models.School), args.Error(1)
}

func (m *MockSchoolService) GetSchoolByID(id uint) (*models.School, error) {
	args := m.Called(id)
	return args.Get(0).(*models.School), args.Error(1)
}

func (m *MockSchoolService) CreateSchool(school *models.School) error {
	args := m.Called(school)
	return args.Error(0)
}

func (m *MockSchoolService) UpdateSchool(school *models.School) error {
	args := m.Called(school)
	return args.Error(0)
}

func (m *MockSchoolService) DeleteSchool(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestSchoolController_GetAllSchools(t *testing.T) {
	mockService := new(MockSchoolService)
	controller := controllers.NewSchoolController(mockService)

	// Mock data
	mockSchools := []models.School{
		{Name: "School A", ClassID: 1},
		{Name: "School B", ClassID: 2},
	}

	// Mock the service method
	mockService.On("GetAllSchools").Return(mockSchools, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/schools", controller.GetAllSchools)
	req, _ := http.NewRequest("GET", "/schools", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var schools []models.School
	err := json.NewDecoder(resp.Body).Decode(&schools)
	assert.Nil(t, err)
	assert.Equal(t, mockSchools, schools)

	mockService.AssertExpectations(t)
}

func TestSchoolController_GetSchoolByID(t *testing.T) {
	mockService := new(MockSchoolService)
	controller := controllers.NewSchoolController(mockService)

	// Mock data
	mockSchool := &models.School{Name: "School A", ClassID: 1}

	// Mock the service method
	mockService.On("GetSchoolByID", mock.AnythingOfType("uint")).Return(mockSchool, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/schools/:id", controller.GetSchoolByID)
	req, _ := http.NewRequest("GET", "/schools/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var school models.School
	err := json.NewDecoder(resp.Body).Decode(&school)
	assert.Nil(t, err)
	assert.Equal(t, mockSchool, &school)

	mockService.AssertExpectations(t)
}

func TestSchoolController_CreateSchool(t *testing.T) {
	mockService := new(MockSchoolService)
	controller := controllers.NewSchoolController(mockService)

	// Mock data
	newSchool := models.School{Name: "New School"}

	// Mock the service method
	mockService.On("CreateSchool", &newSchool).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.POST("/schools", controller.CreateSchool)
	body, _ := json.Marshal(newSchool)
	req, _ := http.NewRequest("POST", "/schools", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.Code)

	mockService.AssertExpectations(t)
}

// controllers/controllers_test.go

func TestSchoolController_UpdateSchool(t *testing.T) {
	mockService := new(MockSchoolService)
	controller := controllers.NewSchoolController(mockService)

	// Mock data
	updatedSchool := models.School{
		Name:    "Updated School",
		ClassID: 1, // Adjust ClassID as per your test scenario
		// Add other fields as necessary for your update operation
	}

	// Mock the service method
	mockService.On("UpdateSchool", mock.AnythingOfType("*models.School")).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.PUT("/schools/:id", controller.UpdateSchool)
	body, _ := json.Marshal(updatedSchool)
	req, _ := http.NewRequest("PUT", "/schools/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}

func TestSchoolController_DeleteSchool(t *testing.T) {
	mockService := new(MockSchoolService)
	controller := controllers.NewSchoolController(mockService)

	// Mock the service method
	mockService.On("DeleteSchool", uint(1)).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.DELETE("/schools/:id", controller.DeleteSchool)
	req, _ := http.NewRequest("DELETE", "/schools/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}

