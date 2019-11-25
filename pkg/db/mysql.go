package db

import (
	"database/sql"
	"net"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func OpenDB(conf *DbConfig, sshc *ssh.Client) (*sql.DB, error) {
	mysqlNet := "tcp"
	if sshc != nil {
		mysqlNet = "mysql+tcp"
		dialFunc := func(addr string) (net.Conn, error) {
			return sshc.Dial("tcp", addr)
		}
		mysql.RegisterDial(mysqlNet, dialFunc)
	}
	dbConf := &mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Addr:                 conf.Host + ":" + conf.Port,
		Net:                  mysqlNet,
		DBName:               conf.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return sql.Open("mysql", dbConf.FormatDSN())
}
