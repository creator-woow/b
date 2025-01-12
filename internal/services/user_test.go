package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"happy-server/internal/db"
	"testing"
)

var (
	mockCreate = CreateUserData{Email: "new.email@gmail.com", FirstName: "User"}
	mockUpdate = UpdateUserData{Email: "another@gmail.com", FirstName: "Another name"}
)

func TestCreateUser(t *testing.T) {
	testDb, mock, conn := db.NewMock()
	defer conn.Close()
	s := UserService{DB: testDb}

	t.Run("create user succeed", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WithArgs(mockCreate.Email, 1).WillReturnRows(
			sqlmock.NewRows([]string{}),
		)
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1),
		)
		mock.ExpectCommit()
		createdUser, creationErr := s.CreateUser(mockCreate)
		if creationErr != nil {
			t.Error(creationErr)
			return
		}
		if !(createdUser.Email == mockCreate.Email && createdUser.FirstName == mockCreate.FirstName) {
			t.Error("created user record does not match with mock data")
		}
	})

	t.Run("create existed user fails", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WithArgs(mockCreate.Email, 1).WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1),
		)
		_, existErr := s.CreateUser(mockCreate)
		// Check if request fails on already existing user
		if existErr == nil {
			t.Error(existErr)
		}
	})
}

func TestGetUser(t *testing.T) {
	testDb, mock, conn := db.NewMock()
	defer conn.Close()
	s := UserService{DB: testDb}
	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "email", "first_name"}).AddRow(
			1, mockCreate.Email, mockCreate.FirstName,
		),
	)
	foundUser, err := s.ReadUserById(1)
	if err != nil {
		t.Error(err)
		return
	}
	if !(foundUser.Email == mockCreate.Email && foundUser.FirstName == mockCreate.FirstName) {
		t.Error("got user record which does not match with mock data")
	}
}

func TestUpdateUser(t *testing.T) {
	testDb, mock, conn := db.NewMock()
	defer conn.Close()
	s := UserService{DB: testDb}

	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "email", "first_name"}).AddRow(
			1, mockCreate.Email, mockCreate.FirstName,
		),
	)
	mock.ExpectBegin()
	mock.ExpectQuery("UPDATE").WithArgs(
		mockUpdate.Email, mockUpdate.FirstName, 1,
	).WillReturnRows(
		sqlmock.NewRows([]string{"email", "first_name", "id"}).AddRow(
			mockUpdate.Email, mockUpdate.FirstName, 1,
		),
	)
	mock.ExpectCommit()
	updatedUser, updateErr := s.UpdateUser(1, mockUpdate)
	if updateErr != nil {
		t.Error(updateErr)
		return
	}
	if !(updatedUser.Email == mockUpdate.Email && updatedUser.FirstName == mockUpdate.FirstName) {
		t.Error("updated user record has expired data")
	}
}

func TestDeleteUser(t *testing.T) {
	testDb, mock, conn := db.NewMock()
	defer conn.Close()
	s := UserService{DB: testDb}
	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "email", "first_name"}).AddRow(
			1, mockCreate.Email, mockCreate.FirstName,
		),
	)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WithArgs(1).WillReturnResult(
		sqlmock.NewResult(0, 1),
	)
	mock.ExpectCommit()
	err := s.DeleteUser(1)
	if err != nil {
		t.Error(err)
	}
}
