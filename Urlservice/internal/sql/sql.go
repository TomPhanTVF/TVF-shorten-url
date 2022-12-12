package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_"github.com/lib/pq"
)

type Sql struct {
	Db *sqlx.DB
	Host string
	Port string
	Username string
	Password string
	DbName string
}


func (s *Sql) ConnectDB(){
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.DbName)
	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return 
	}
}

func (s *Sql) Close(){
	s.Db.Close()
}