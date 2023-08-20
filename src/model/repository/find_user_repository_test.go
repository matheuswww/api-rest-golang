package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matheuswww/api-rest-golang/src/configuration/database/sql_mock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery("SELECT email,password,name,age,id FROM users WHERE email = ?").
			WithArgs("test@test.com").
			WillReturnRows(sqlmock.NewRows([]string{"email", "password", "name", "age", "id"}).
			AddRow("test@test.com", "test", "test", uint8(90), 1))

		repo := NewUserRepository(db)
		userDomain, err := repo.FindUser("email","test@test.com")

		assert.Nil(t, err)
		assert.NotNil(t, userDomain.GetId())
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetName(), "test")
		assert.EqualValues(t, userDomain.GetPassword(), "test")
		assert.EqualValues(t, userDomain.GetAge(), uint8(90))
	})

	t.Run("return_error_from_database", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery("SELECT email,password,name,age,id FROM users WHERE email = ?").
			WithArgs("test@test.com").
			WillReturnError(errors.New("database error"))

		repo := NewUserRepository(db)
		userDomain, err := repo.FindUser("email","test@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	t.Run("error_user_notfound", func(t *testing.T) {
		db, mock, err := sql_mock.NewMockedMysqlConnection()
		if err != nil {
			t.Fatalf("Erro ao criar conexão falsa com o banco de dados: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery("SELECT email,password,name,age,id FROM users WHERE email = ?").
			WithArgs("test@test.com").
			WillReturnRows(sqlmock.NewRows([]string{}))
	
		repo := NewUserRepository(db)
		userDomain, err := repo.FindUser("email","test@test.com")

		assert.NotNil(t,err)
		assert.Nil(t,userDomain)
	})
}