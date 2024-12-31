package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"happy-server/db"
	"testing"
)

var (
	md = createUserData{Email: "new.email@gmail.com", FirstName: "User"}
	// todo: how to sync mock rows creation and models like User? Sync after search.
	mr = sqlmock.NewRows([]string{"id", "email", "first_name"}).AddRow(1, md.Email, md.FirstName)
)

func TestCreateUser(t *testing.T) {
	testDb, _ := db.NewMockConnection()
	createdUser, creationErr := CreateUser(testDb, md)
	if creationErr != nil {
		t.Error(creationErr)
		return
	}
	if createdUser.Email != md.Email || createdUser.FirstName != md.FirstName {
		t.Error("created user record does not match with mock data")
	}
}

func TestGetUser(t *testing.T) {
	testDb, mock := db.NewMockConnection()
	mock.ExpectQuery("SELECT").WillReturnRows(mr)
	foundUser, err := GetUser(testDb, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if foundUser.Email != md.Email || foundUser.FirstName != md.FirstName {
		t.Error("got user record does not match with mock data")
	}
}

func TestUpdateUser(t *testing.T) {
	testDb, mock := db.NewMockConnection()
	mock.ExpectQuery("SELECT").WillReturnRows(mr)
	mock.ExpectExec("UPDATE").WillReturnResult(
		sqlmock.NewResult(1, 1),
	)
	updateData := updateUserData{Email: "another@gmail.com", FirstName: "Another name"}
	updatedUser, updateErr := UpdateUser(testDb, 1, updateData)
	if updateErr != nil {
		t.Error(updateErr)
		return
	}
	if updatedUser.Email != updateData.Email || updatedUser.FirstName == updateData.FirstName {
		t.Error("updated user record still much previous data")
	}
}

func TestDeleteUser(t *testing.T) {
	testDb, mock := db.NewMockConnection()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeleteUser(testDb, 1)
	if err != nil {
		t.Error(err)
	}
}
