package controllers

import (
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/models"
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TestfieldController struct {
	testfieldService *services.TestfieldService
}

func NewTestfieldController() (*TestfieldController, error) {
	testfieldService, err := services.NewTestfieldService()
	if err != nil {
		return nil, err
	}
	return &TestfieldController{
		testfieldService: testfieldService,
	}, nil
}

func (testfieldController *TestfieldController) CreateTestfield(context *gin.Context) {
	// validate input
	var input models.Testfield
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger testfield creation
	if _, err := testfieldController.testfieldService.CreateTestfield(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Testfield created successfully"})
}

func (testfieldController *TestfieldController) UpdateTestfield(context *gin.Context) {
	// validate input
	var input models.Testfield
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfield update
	if _, err := testfieldController.testfieldService.UpdateTestfield(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Testfield updated successfully"})
}

func (testfieldController *TestfieldController) FetchTestfield(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfield fetching
	testfield, err := testfieldController.testfieldService.GetTestfield(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfield)
}

func (testfieldController *TestfieldController) DeleteTestfield(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfield deletion
	if err := testfieldController.testfieldService.DeleteTestfield(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Testfield deleted successfully",
	})
}

func (testfieldController *TestfieldController) ListTestfields(context *gin.Context) {
	// trigger all testfields fetching
	testfields, err := testfieldController.testfieldService.ListTestfields()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfields)
}

func (*TestfieldController) PatchTestfield(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TestfieldController) OptionsTestfield(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TestfieldController) HeadTestfield(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
