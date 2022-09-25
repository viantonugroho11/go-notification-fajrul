package config

import (
	"database/sql"
	"fmt"
)



type DatabaseConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	NameDatabase  string
	DB sql.DB
	// SSLMode   string
	// ConnMaxLifetime    int
	// MaxOpenConnections int
	// MaxIdleConnections int

}

func InitDb(conf Config) DatabaseConfig{

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.DBName)
	db, err := sql.Open("mysql",conn)
	if err != nil {
		fmt.Println(err)
	}
	return DatabaseConfig{
		Host:     conf.Database.Host,
		Port:     conf.Database.Port,
		Username: conf.Database.User,
		Password: conf.Database.Password,
		NameDatabase: conf.Database.DBName,
		DB: *db,
	}
}