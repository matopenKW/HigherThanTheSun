package apps

import (
	"HigherThanTheSun/pkg/dto"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func LoginExecute(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	user := getLoginUser(r)
	if user != nil {
		jsonval, _ := json.Marshal(&user)
		fmt.Println(user, " → ", string(jsonval))
		cookie := &http.Cookie{
			Name:  "userinfo",
			Value: string(jsonval),
		}
		http.SetCookie(w, cookie)
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("../web/templates/login.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		t, _ := template.ParseFiles("../web/templates/login.gtpl")
		t.Execute(w, nil)

	}
}

func getLoginUser(r *http.Request) *dto.User {
	if idToken := r.Form["idToken"]; idToken != nil {
		token := getToken(idToken[0])
		return dto.NewUser(token.UID, "", idToken[0])
	} else {
		return nil
	}
}

func getToken(idToken string) *auth.Token {
	opt := option.WithCredentialsFile("key.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)

	return token
}
