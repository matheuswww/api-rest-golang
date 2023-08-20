package mysql

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
)


func NewMysqlConnection() (*sql.DB,error) {
	err := godotenv.Load();
	if err != nil {
		logger.Error("ENV LOADING ERROR!!!",err)
		return nil,err
	}
	db_name := os.Getenv("MYSQL_NAME")
	db_pass := os.Getenv("MYSQL_PASS")
	db, err := sql.Open("mysql", "root:" + db_pass + "@tcp(172.17.0.2)/" + db_name)
	if err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!",err)
		return nil,err
	}

	if err := db.Ping();err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!",err)
		return nil,err
	}

	logger.Info("DB MYSQL IS RUNNING!!!!!")
	db.SetMaxOpenConns(10)   
	db.SetConnMaxLifetime(time.Minute * 3)
	return db,nil
}