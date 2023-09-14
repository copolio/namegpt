package database

import (
	"fmt"
	"log"
)

type Datasource struct {
	Ddl         Ddl    `yaml:"ddl"`
	Host        string `yaml:"host"`
	Protocol    string `yaml:"protocol"`
	Schema      string `yaml:"schema"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

func (c *Datasource) GetMysqlDSN() string {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Protocol, c.Host, c.Schema)
	log.Default().Println(dsn)
	return dsn
}
