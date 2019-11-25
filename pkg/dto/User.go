package dto

import "encoding/json"

type User struct {
	id    string
	pass  string
	token string
}

func (u User) GetId() string {
	return u.id
}

func (u User) GetPass() string {
	return u.pass
}

func (u User) GetToken() string {
	return u.token
}

func NewUser(id, name, token string) *User {
	return &User{id, name, token}
}

func (u *User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Id    string
		Pass  string
		Token string
	}{
		Id:    u.GetId(),
		Pass:  u.GetPass(),
		Token: u.GetToken(),
	})
	return v, err
}
