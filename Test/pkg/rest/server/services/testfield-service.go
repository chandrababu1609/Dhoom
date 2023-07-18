package services

import (
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/daos"
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/models"
)

type TestfieldService struct {
	testfieldDao *daos.TestfieldDao
}

func NewTestfieldService() (*TestfieldService, error) {
	testfieldDao, err := daos.NewTestfieldDao()
	if err != nil {
		return nil, err
	}
	return &TestfieldService{
		testfieldDao: testfieldDao,
	}, nil
}

func (testfieldService *TestfieldService) CreateTestfield(testfield *models.Testfield) (*models.Testfield, error) {
	return testfieldService.testfieldDao.CreateTestfield(testfield)
}

func (testfieldService *TestfieldService) UpdateTestfield(id int64, testfield *models.Testfield) (*models.Testfield, error) {
	return testfieldService.testfieldDao.UpdateTestfield(id, testfield)
}

func (testfieldService *TestfieldService) DeleteTestfield(id int64) error {
	return testfieldService.testfieldDao.DeleteTestfield(id)
}

func (testfieldService *TestfieldService) ListTestfields() ([]*models.Testfield, error) {
	return testfieldService.testfieldDao.ListTestfields()
}

func (testfieldService *TestfieldService) GetTestfield(id int64) (*models.Testfield, error) {
	return testfieldService.testfieldDao.GetTestfield(id)
}
