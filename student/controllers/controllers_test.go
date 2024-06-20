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

// controllers/controllers_test.go
type MockClassService struct {
	mock.Mock
}

// GetAllClasses mocks the GetAllClasses method of ClassService interface
func (m *MockClassService) GetAllClasses() ([]models.Class, error) {
	args := m.Called()
	return args.Get(0).([]models.Class), args.Error(1)
}

// GetClassByID mocks the GetClassByID method of ClassService interface
func (m *MockClassService) GetClassByID(id uint) (*models.Class, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Class), args.Error(1)
}

// CreateClass mocks the CreateClass method of ClassService interface
func (m *MockClassService) CreateClass(class *models.Class) error {
	args := m.Called(class)
	return args.Error(0)
}

// UpdateClass mocks the UpdateClass method of ClassService interface
func (m *MockClassService) UpdateClass(class *models.Class) error {
	args := m.Called(class)
	return args.Error(0)
}

// DeleteClass mocks the DeleteClass method of ClassService interface
func (m *MockClassService) DeleteClass(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
func TestClassController_GetAllClasses(t *testing.T) {
	mockService := new(MockClassService)
	controller := controllers.NewClassController(mockService)

	// Mock data
	mockClasses := []models.Class{
		{ClassName: "Class A", StudentID: 1},
		{ClassName: "Class B", StudentID: 2},
	}

	// Mock the service method
	mockService.On("GetAllClasses").Return(mockClasses, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/classes", controller.GetAllClasses)
	req, _ := http.NewRequest("GET", "/classes", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var classes []models.Class
	err := json.NewDecoder(resp.Body).Decode(&classes)
	assert.Nil(t, err)
	assert.Equal(t, mockClasses, classes)

	mockService.AssertExpectations(t)
}

func TestClassController_GetClassByID(t *testing.T) {
	mockService := new(MockClassService)
	controller := controllers.NewClassController(mockService)

	// Mock data
	mockClass := &models.Class{ClassName: "Class A", StudentID: 1}

	// Mock the service method
	mockService.On("GetClassByID", mock.AnythingOfType("uint")).Return(mockClass, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/classes/:id", controller.GetClassByID)
	req, _ := http.NewRequest("GET", "/classes/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var classObj models.Class
	err := json.NewDecoder(resp.Body).Decode(&classObj)
	assert.Nil(t, err)
	assert.Equal(t, mockClass, &classObj)

	mockService.AssertExpectations(t)
}

func TestClassController_CreateClass(t *testing.T) {
	mockService := new(MockClassService)
	controller := controllers.NewClassController(mockService)

	// Mock data
	newClass := models.Class{ClassName: "New Class"}

	// Mock the service method
	mockService.On("CreateClass", &newClass).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.POST("/classes", controller.CreateClass)
	body, _ := json.Marshal(newClass)
	req, _ := http.NewRequest("POST", "/classes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.Code)

	mockService.AssertExpectations(t)
}

func TestClassController_UpdateClass(t *testing.T) {
	mockService := new(MockClassService)
	controller := controllers.NewClassController(mockService)

	// Mock data
	updatedClass := models.Class{
		ClassName: "Updated Class",
		StudentID: 1, // Adjust SchoolID as per your test scenario
		// Add other fields as necessary for your update operation
	}

	// Mock the service method
	mockService.On("UpdateClass", mock.AnythingOfType("*models.Class")).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.PUT("/classes/:id", controller.UpdateClass)
	body, _ := json.Marshal(updatedClass)
	req, _ := http.NewRequest("PUT", "/classes/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}

func TestClassController_DeleteClass(t *testing.T) {
	mockService := new(MockClassService)
	controller := controllers.NewClassController(mockService)

	// Mock the service method
	mockService.On("DeleteClass", uint(1)).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.DELETE("/classes/:id", controller.DeleteClass)
	req, _ := http.NewRequest("DELETE", "/classes/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}

type MockStudentService struct {
	mock.Mock
}

// GetAllStudents mocks the GetAllStudents method of StudentService interface
func (m *MockStudentService) GetAllStudents() ([]models.Student, error) {
	args := m.Called()
	return args.Get(0).([]models.Student), args.Error(1)
}

// GetStudentByID mocks the GetStudentByID method of StudentService interface
func (m *MockStudentService) GetStudentByID(id uint) (*models.Student, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Student), args.Error(1)
}

// CreateStudent mocks the CreateStudent method of StudentService interface
func (m *MockStudentService) CreateStudent(student *models.Student) error {
	args := m.Called(student)
	return args.Error(0)
}

// UpdateStudent mocks the UpdateStudent method of StudentService interface
func (m *MockStudentService) UpdateStudent(student *models.Student) error {
	args := m.Called(student)
	return args.Error(0)
}

// DeleteStudent mocks the DeleteStudent method of StudentService interface
func (m *MockStudentService) DeleteStudent(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// controllers/controllers_test.go

func TestStudentController_GetAllStudents(t *testing.T) {
	mockService := new(MockStudentService)
	controller := controllers.NewStudentController(mockService)

	// Mock data
	mockStudents := []models.Student{
		{Name: "Student A", Marks: 98},
		{Name: "Student B", Marks: 96},
	}

	// Mock the service method
	mockService.On("GetAllStudents").Return(mockStudents, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/students", controller.GetAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var students []models.Student
	err := json.NewDecoder(resp.Body).Decode(&students)
	assert.Nil(t, err)
	assert.Equal(t, mockStudents, students)

	mockService.AssertExpectations(t)
}

func TestStudentController_GetStudentByID(t *testing.T) {
	mockService := new(MockStudentService)
	controller := controllers.NewStudentController(mockService)

	// Mock data
	mockStudent := &models.Student{Name: "Student A", Marks: 98}

	// Mock the service method
	mockService.On("GetStudentByID", mock.AnythingOfType("uint")).Return(mockStudent, nil)

	// Create a gin context
	router := gin.Default()
	router.GET("/students/:id", controller.GetStudentByID)
	req, _ := http.NewRequest("GET", "/students/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var studentObj models.Student
	err := json.NewDecoder(resp.Body).Decode(&studentObj)
	assert.Nil(t, err)
	assert.Equal(t, mockStudent, &studentObj)

	mockService.AssertExpectations(t)
}

func TestStudentController_CreateStudent(t *testing.T) {
	mockService := new(MockStudentService)
	controller := controllers.NewStudentController(mockService)

	// Mock data
	newStudent := models.Student{Name: "New Student"}

	// Mock the service method
	mockService.On("CreateStudent", &newStudent).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.POST("/students", controller.CreateStudent)
	body, _ := json.Marshal(newStudent)
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.Code)

	mockService.AssertExpectations(t)
}

func TestStudentController_UpdateStudent(t *testing.T) {
	mockService := new(MockStudentService)
	controller := controllers.NewStudentController(mockService)

	// Mock data
	updatedStudent := models.Student{
		Name:  "Updated Student",
		Marks: 94, // Adjust ClassID as per your test scenario
		// Add other fields as necessary for your update operation
	}

	// Mock the service method
	mockService.On("UpdateStudent", mock.AnythingOfType("*models.Student")).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.PUT("/students/:id", controller.UpdateStudent)
	body, _ := json.Marshal(updatedStudent)
	req, _ := http.NewRequest("PUT", "/students/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}

func TestStudentController_DeleteStudent(t *testing.T) {
	mockService := new(MockStudentService)
	controller := controllers.NewStudentController(mockService)

	// Mock the service method
	mockService.On("DeleteStudent", uint(1)).Return(nil)

	// Create a gin context
	router := gin.Default()
	router.DELETE("/students/:id", controller.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	mockService.AssertExpectations(t)
}
