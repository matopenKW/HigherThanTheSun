package main

import (
	"HigherThanTheSun/internal/apps"
	"HigherThanTheSun/pkg/db"
	"HigherThanTheSun/pkg/ssh"
	"fmt"
	"log"
	"net/http"

	"github.com/go-ini/ini"
)

func forward(w http.ResponseWriter, r *http.Request) {

	cfg, _ := ini.Load("../pkg/conf/config.ini")
	sshSection := cfg.Section("ssh")
	conf := &ssh.SshConfig{
		Key:      sshSection.Key("key").String(),
		Host:     sshSection.Key("host").MustString("localhost"),
		Port:     sshSection.Key("port").MustString("22"),
		User:     sshSection.Key("user").String(),
		Password: sshSection.Key("password").String(),
	}

	sshConn, err := ssh.OpenSSH(conf)
	if err != nil {
		log.Fatalf("%s error: %v", "Dial", err)
	}
	defer sshConn.Close()

	session, err := sshConn.NewSession()
	if err != nil {
		log.Fatalf("%s error: %v", "NewSession", err)
	}
	defer session.Close()

	dbSection := cfg.Section("db")
	// DB接続
	dbConf := &db.DbConfig{
		Host:     dbSection.Key("host").MustString("localhost"),
		Port:     dbSection.Key("port").MustString("3306"),
		User:     dbSection.Key("user").String(),
		Password: dbSection.Key("password").String(),
		DBName:   dbSection.Key("dbname").String(),
	}
	dbClient, err := db.OpenDB(dbConf, sshConn)
	if err != nil {
		log.Fatalf("erro in new db client. reason : %v\n", err)
	}
	defer dbClient.Close()

	urlPath := r.URL.Path
	url := urlPath[1:len(urlPath)]

	if url == "" || url == "login" {
		apps.LoginExecute(w, r)
	} else {
		// token check
		cookie, err := r.Cookie("userinfo")
		if err != nil {
			log.Fatalf("error user not found: %v\n", err)
		}

		fmt.Println(cookie.Value)

		switch {
		case url == "top":
			apps.TopExecute(w, r, dbClient)
		default:
			apps.LoginExecute(w, r)
		}
	}
}

func main() {
	fmt.Println("------------アクセス------------")
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("../web/"))))
	http.HandleFunc("/", forward)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("%s error: %v", "ListenAndServe", err)
	}
}
