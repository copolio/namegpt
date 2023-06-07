package database

import (
	"fmt"
	"log"
)

type Datasource struct {
	Ddl         Ddl    `yaml:"ddl"`
	Url         string `yaml:"url"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

func (c *Datasource) GetMysqlDSN() string {
	dsn := fmt.Sprintf("%s:%s@%s", c.User, c.Password, c.Url)
	log.Default().Println(dsn)
	return dsn
}
