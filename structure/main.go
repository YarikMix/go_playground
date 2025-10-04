package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

type (
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data,omitempty"`
	}

	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Data struct {
		Type       string     `json:"type"`
		Id         int        `json:"id"`
		Attributes Attributes `json:"attributes"`
	}

	Attributes struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

func main() {
	man := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}

	fmt.Println()

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}

	reqRaw, _ := json.Marshal(req)
	fmt.Println(string(reqRaw))

	const rawResp = `
		{
			"header": {
				"code": 0,
				"message": ""
			},
			"data": [{
				"type": "user",
				"id": 100,
				"attributes": {
					"email": "bob@yandex.ru",
					"article_ids": [10, 11, 12]
				}
			}]
		}
	`

	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}
