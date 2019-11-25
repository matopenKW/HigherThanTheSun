package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Profile struct {
	Name          string
	Favouritefood string
	NckName       string
	Mymotto       string
}

type TopForm struct {
	Profile Profile
}

func TopExecute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("method:", r.Method)

	fmt.Println("-------------cookiechk")
	cookie, err := r.Cookie("userinfo")
	if err != nil {
		log.Fatal("Cookie: ", err)
	}
	fmt.Println(cookie.Value)

	//	if r.Method == "GET" {
	t := template.Must(template.ParseFiles("../web/templates/top.gtpl"))
	topForm := &TopForm{Profile{"おの", "メロンパン", "わろん", "町田の切れたナイフ"}}
	t.ExecuteTemplate(w, "top.gtpl", topForm)
	//}
}
