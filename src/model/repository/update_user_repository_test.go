package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matheuswww/api-rest-golang/src/configuration/database/sql_mock"
	"github.com/matheuswww/api-rest-golang/src/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET 
		name = CASE WHEN ? <> '' THEN ? ELSE name END,
		age = CASE WHEN ? > 0 THEN ? ELSE age END WHERE id = ?`)).
			WithArgs("test", "test",90,90,"1").
			WillReturnResult(sqlmock.NewResult(1,1))

		repo := NewUserRepository(db)
		userDomain := model.NewUserDomain(
			"test@test.com",
			"test",
			"test",
			90,
			0,
		)
		err = repo.UpdateUser("1",userDomain)

		assert.Nil(t, err)
	})
	
	t.Run("return_error_from_database",func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
		}
		defer db.Close()		
		
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET 
		name = CASE WHEN ? <> '' THEN ? ELSE name END,
		age = CASE WHEN ? > 0 THEN ? ELSE age END WHERE id = ?`)).
			WithArgs("test", "test",90,90,"1").
			WillReturnError(errors.New("database error"))

		repo := NewUserRepository(db)
		userDomain := model.NewUserDomain(
			"test@test.com",
			"test",
			"test",
			90,
			0,
		)
		err = repo.UpdateUser("1",userDomain)

		assert.NotNil(t, err)
	})
}