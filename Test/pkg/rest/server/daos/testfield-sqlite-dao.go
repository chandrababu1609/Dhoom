package daos

import (
	"database/sql"
	"errors"
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/daos/clients/sqls"
	"github.com/chandrababu1609/Dhoom/test/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type TestfieldDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTestfields(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS testfields(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Username TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTestfieldDao() (*TestfieldDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTestfields(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TestfieldDao{
		sqlClient,
	}, nil
}

func (testfieldDao *TestfieldDao) CreateTestfield(m *models.Testfield) (*models.Testfield, error) {
	insertQuery := "INSERT INTO testfields(Username)values(?)"
	res, err := testfieldDao.sqlClient.DB.Exec(insertQuery, m.Username)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("testfield created")
	return m, nil
}

func (testfieldDao *TestfieldDao) UpdateTestfield(id int64, m *models.Testfield) (*models.Testfield, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	testfield, err := testfieldDao.GetTestfield(id)
	if err != nil {
		return nil, err
	}
	if testfield == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE testfields SET Username = ? WHERE Id = ?"
	res, err := testfieldDao.sqlClient.DB.Exec(updateQuery, m.Username, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("testfield updated")
	return m, nil
}

func (testfieldDao *TestfieldDao) DeleteTestfield(id int64) error {
	deleteQuery := "DELETE FROM testfields WHERE Id = ?"
	res, err := testfieldDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("testfield deleted")
	return nil
}

func (testfieldDao *TestfieldDao) ListTestfields() ([]*models.Testfield, error) {
	selectQuery := "SELECT * FROM testfields"
	rows, err := testfieldDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var testfields []*models.Testfield
	for rows.Next() {
		m := models.Testfield{}
		if err = rows.Scan(&m.Id, &m.Username); err != nil {
			return nil, err
		}
		testfields = append(testfields, &m)
	}
	if testfields == nil {
		testfields = []*models.Testfield{}
	}

	log.Debugf("testfield listed")
	return testfields, nil
}

func (testfieldDao *TestfieldDao) GetTestfield(id int64) (*models.Testfield, error) {
	selectQuery := "SELECT * FROM testfields WHERE Id = ?"
	row := testfieldDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Testfield{}
	if err := row.Scan(&m.Id, &m.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("testfield retrieved")
	return &m, nil
}
