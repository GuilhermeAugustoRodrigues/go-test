package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"gin-api-rest/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getTestRoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := routes.GetRoutesSetup()

	return r
}

func TestHello(t *testing.T) {
	r := getTestRoutesSetup()
	request, _ := http.NewRequest("GET", "/gui", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)

	expectedResponseBody := `{"API says":"Hello, gui"}`
	actualResponseBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, expectedResponseBody, string(actualResponseBody), "Status should be: %d but was %d", expectedResponseBody, string(actualResponseBody))
}

var student = models.Student{
	Name: "Mock Student",
	RG:   "123456789",
	CPF:  "12345678910",
}

func TestCreateMockStudent(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()

	studentJson, _ := json.Marshal(student)

	request, _ := http.NewRequest("POST", "/alunos", bytes.NewBuffer(studentJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)

	actualResponseBody := models.Student{}
	err := json.Unmarshal(response.Body.Bytes(), &actualResponseBody)

	fmt.Println(actualResponseBody)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, student.Name, actualResponseBody.Name, "Response body 'Name' should be: %v but was %v", student.Name, actualResponseBody.Name)
	assert.Equal(t, student.CPF, actualResponseBody.CPF, "Response body 'CPF' should be: %v but was %v", student.CPF, actualResponseBody.CPF)
	assert.Equal(t, student.RG, actualResponseBody.RG, "Response body 'RG' should be: %v but was %v", student.RG, actualResponseBody.RG)

	student.ID = actualResponseBody.ID
}

func TestGetStudents(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()
	request, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)
}

func TestGetStudentById(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()
	request, _ := http.NewRequest("GET", "/alunos/"+strconv.FormatUint(uint64(student.ID), 10), nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)

	actualResponseBody := models.Student{}
	err := json.Unmarshal(response.Body.Bytes(), &actualResponseBody)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, student.Name, actualResponseBody.Name, "Response body 'Name' should be: %v but was %v", student.Name, actualResponseBody.Name)
	assert.Equal(t, student.CPF, actualResponseBody.CPF, "Response body 'CPF' should be: %v but was %v", student.CPF, actualResponseBody.CPF)
	assert.Equal(t, student.RG, actualResponseBody.RG, "Response body 'RG' should be: %v but was %v", student.RG, actualResponseBody.RG)
}

func TestGetStudentBySearch(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()
	request, _ := http.NewRequest("GET", "/alunos/search/Mock", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)

	actualResponseBody := []models.Student{}
	err := json.Unmarshal(response.Body.Bytes(), &actualResponseBody)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, student.Name, actualResponseBody[0].Name, "Response body 'Name' should be: %v but was %v", student.Name, actualResponseBody[0].Name)
	assert.Equal(t, student.CPF, actualResponseBody[0].CPF, "Response body 'CPF' should be: %v but was %v", student.CPF, actualResponseBody[0].CPF)
	assert.Equal(t, student.RG, actualResponseBody[0].RG, "Response body 'RG' should be: %v but was %v", student.RG, actualResponseBody[0].RG)
}

func TestEditStudent(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()
	editedStudent := models.Student{
		Name: student.Name + " Edited",
		RG:   student.RG,
		CPF:  student.CPF,
	}

	editedStudentJson, _ := json.Marshal(editedStudent)

	request, _ := http.NewRequest("PATCH", "/alunos/"+strconv.FormatUint(uint64(student.ID), 10), bytes.NewBuffer(editedStudentJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)

	actualResponseBody := models.Student{}
	err := json.Unmarshal(response.Body.Bytes(), &actualResponseBody)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, editedStudent.Name, actualResponseBody.Name, "Response body 'Name' should be: %v but was %v", editedStudent.Name, actualResponseBody.Name)
	assert.Equal(t, editedStudent.CPF, actualResponseBody.CPF, "Response body 'CPF' should be: %v but was %v", editedStudent.CPF, actualResponseBody.CPF)
	assert.Equal(t, editedStudent.RG, actualResponseBody.RG, "Response body 'RG' should be: %v but was %v", editedStudent.RG, actualResponseBody.RG)
}

func TestDeleteMockStudent(t *testing.T) {
	database.StartDatabaseConnection()
	r := getTestRoutesSetup()
	request, _ := http.NewRequest("DELETE", "/alunos/"+strconv.FormatUint(uint64(student.ID), 10), nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be: %d but was %d", http.StatusOK, response.Code)
}
