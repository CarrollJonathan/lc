package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
}

type studentAPI struct {
	studentRepo repo.StudentRepository
}

func NewStudentAPI(studentRepo repo.StudentRepository) *studentAPI {
	return &studentAPI{studentRepo}
}

func (s *studentAPI) AddStudent(c *gin.Context) {
	var newStudent model.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := s.studentRepo.Store(&newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add student success"})
}

func (s *studentAPI) DeleteStudent(c *gin.Context) {

	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse := model.ErrorResponse{Error: "Invalid student ID"}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err = s.studentRepo.Delete(studentID)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error: "Failed to delete the student",
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// If the deletion is successful, return a success response
	successResponse := model.SuccessResponse{
		Message: "student delete success",
	}
	c.JSON(http.StatusOK, successResponse)
	// TODO: answer here
}
