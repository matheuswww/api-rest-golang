package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/matheuswww/api-rest-golang/src/configuration/database/sql_mock"
	"github.com/matheuswww/api-rest-golang/src/model"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery("SELECT email,password,name,age,id FROM users WHERE email = ?").
			WithArgs("test@test.com").
			WillReturnRows(sqlmock.NewRows([]string{}))

		mock.ExpectExec("INSERT INTO users (.*)").
			WithArgs("test@test.com", "test", "test", 90).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewUserRepository(db)
		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com",
			"test",
			"test",
			90,
			0,
		))

		assert.Nil(t, err)
		assert.NotNil(t, userDomain.GetId())
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetName(), "test")
		assert.EqualValues(t, userDomain.GetPassword(), "test")
		assert.EqualValues(t, userDomain.GetAge(), 90)
	})

	t.Run("return_error_from_database",func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery("SELECT email,password,name,age,id FROM users WHERE email = ?").
			WithArgs("test@test.com").
			WillReturnRows(sqlmock.NewRows([]string{}))

		mock.ExpectExec("INSERT INTO users (.*)").
			WithArgs("test@test.com", "test", "test", 90).
			WillReturnError(errors.New("database error"))

		repo := NewUserRepository(db)
		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com",
			"test",
			"test",
			90,
			0,
		))

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
