package sql_mock

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
)

func NewMockedMysqlConnection() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	logger.Info("DB MYSQL IS RUNNING!!!!!")
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return db, mock, nil
}	
