package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func getQueryFromFile(token string, command string) (string, error) {

	// Выполнение запроса к GigaChat API

	requestData := map[string]interface{}{
		"model": "GigaChat",
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": command,
			},
		},
		"n":                  1,
		"stream":             false,
		"max_tokens":         512,
		"repetition_penalty": 1,
		"update_interval":    0,
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(requestData).
		Post("https://gigachat.devices.sberbank.ru/api/v1/chat/completions")

	if err != nil {
		log.Fatalf("Ошибка получения токена: %v", err)
	}

	// Проверка статуса ответа

	if response.StatusCode() == 401 {
		err = fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode())
		return response.String(), err
	}

	if response.StatusCode() == 404 {
		log.Fatalf("Что-то не так")
	}
	if response.StatusCode() != 200 {
		log.Fatalf("Error: Received status code %d", response.StatusCode())
	}

	var w1 Api

	err = json.Unmarshal(response.Body(), &w1)
	if err != nil {
		fmt.Println(err)
	}

	return w1.Choices[0].Message.Content, nil
}
