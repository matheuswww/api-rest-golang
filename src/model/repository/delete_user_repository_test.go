package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matheuswww/api-rest-golang/src/configuration/database/sql_mock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectExec("DELETE FROM users WHERE id = ?").
			WithArgs(uint64(1)).
			WillReturnResult(sqlmock.NewResult(1,1))

		repo := NewUserRepository(db)
		err = repo.DeleteUser("1")

		assert.Nil(t, err)
	})
	t.Run("return_error_from_database",func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(uint64(1)).
		WillReturnError(errors.New("database error"))

		repo := NewUserRepository(db)
		err = repo.DeleteUser("1")

		assert.NotNil(t, err)
	})
}
