package main

import (
	"HigherThanTheSun/util/db"
	"HigherThanTheSun/util/ssh"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	conf := &ssh.SshConfig{
		Key:      "",
		Host:     "localhost",
		Port:     "22",
		User:     "",
		Password: "",
	}

	// ssh接続
	sshConn, err := ssh.OpenSSH(conf)
	if err != nil {
		log.Fatalf("%s error: %v", "Dial", err)
	}
	defer sshConn.Close()

	session, err := sshConn.NewSession()
	if err != nil {
		log.Fatalf("%s error: %v", "NewSession", err)
	}

	log.Println("接続成功")
	defer session.Close()

	// DB接続
	dbConf := &db.DbConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "",
		Password: "",
		DBName:   "testDB",
	}
	dbClient, err := db.OpenDB(dbConf, sshConn)
	if err != nil {
		log.Fatalf("erro in new db client. reason : %v\n", err)
	}
	defer dbClient.Close()

	rows, err := dbClient.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))
	for i := range values {
		args[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(args...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, val := range values {
			if val == nil {
				value = "NULL"
			} else {
				value = string(val)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
