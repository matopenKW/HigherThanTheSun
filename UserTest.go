package main

import (
	"HigherThanTheSun/pkg/dto"
	"encoding/json"
	"fmt"
)

type JSONSample struct {
	field     string `json:"field"`
	Omit      string `json:"-"`
	OmitEmpty string `json:"omit_empty,omitempty"`
	Num       int    `json:"num,string"`
}

func main() {
	user := dto.NewUser("0001", "kazu", "asdasd")
	fmt.Println(user)
	jsonval, _ := json.Marshal(&user)
	fmt.Println(string(jsonval))

	sample := JSONSample{
		field:     "field",
		Omit:      "omit",
		OmitEmpty: "",
		Num:       1,
	}
	fmt.Println(&sample)
	jsonval2, _ := json.Marshal(&sample)
	fmt.Println(string(jsonval2))

}
