package driver

import (
	"log"
	"regexp"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/driver/mock"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/gateway/repository"
)

func TestGetAll(t *testing.T) {
	db, mockdb, err := mock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	suuid1 := uuid.New().String()
	suuid2 := uuid.New().String()
	rows := mockdb.NewRows([]string{"id","userid","username"}).
					AddRow(1,suuid1,"testuser1").
					AddRow(2,suuid2,"testuser2")
	
	mockdb.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).WillReturnRows(rows)

	driver := NewDBdriver(db)
	users, err := driver.GetAll()
	expected := repository.RepositoryJsons{
		{Id: 1,Userid: suuid1,Username: "testuser1"},
		{Id: 2,Userid: suuid2,Username: "testuser2"},
	}
	assert.Equal(t , users, expected)
}

func TestGetById(t *testing.T) {
	db, mockdb, err := mock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	suuid1 := uuid.New()
	rows := mockdb.NewRows([]string{"id","userid","username"}).
					AddRow(1,suuid1.String(),"testuser1")
	
	mockdb.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE userid = ?")).WithArgs(suuid1).WillReturnRows(rows)

	driver := NewDBdriver(db)
	user, err := driver.GetById(suuid1)
	
	expected := repository.RepositoryJson{Id: 1,Userid: suuid1.String(),Username: "testuser1"}
	assert.Equal(t, user, expected)
	assert.Equal(t, err, nil)
}

func TestCreate(t *testing.T) {
	t.Skip()
	db, mockdb, err := mock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	suuid := uuid.New().String()
	username := "testuser1"

	mockdb.ExpectBegin()
	mockdb.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user" ("userid", "username") VALUES ($1, $2)`)).WillReturnError(nil)
	mockdb.ExpectCommit()

	driver := NewDBdriver(db)
	param := repository.RepositoryJson{Userid: suuid,Username: username}
	err1 := driver.Create(param)
	err2 := mockdb.ExpectationsWereMet()
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
}

func TestUpdate(t *testing.T) {
	t.Skip()
	db, mockdb, err := mock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	suuid := uuid.New()
	username := "testuser1"

	mockdb.ExpectBegin()
	mockdb.ExpectQuery(regexp.QuoteMeta(`UPDATE "user" SET "userid" = ?, "username" = ? WHERE userid = ?`)).WillReturnError(nil)
	mockdb.ExpectCommit()

	driver := NewDBdriver(db)
	param := repository.RepositoryParamJson{Userid: suuid.String(),Username: username}
	err1 := driver.Update(suuid,param)
	err2 := mockdb.ExpectationsWereMet()
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
}

func TestDelete(t *testing.T) {
	t.Skip()
	db, mockdb, err := mock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	suuid := uuid.New()

	mockdb.ExpectBegin()
	mockdb.ExpectQuery(regexp.QuoteMeta(`DELETE FROM "user" WHERE "userid" = ?`)).WillReturnError(nil)
	mockdb.ExpectCommit()

	driver := NewDBdriver(db)
	err1 := driver.Delete(suuid)
	err2 := mockdb.ExpectationsWereMet()
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
}