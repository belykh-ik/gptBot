package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

func getToken() (string, error) {

	// Выполнение запроса к GigaChat API

	//Получение токена
	// Уникальный идентификатор запроса
	rqUID := uuid.New().String()
	response, err := resty.New().R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetHeader("RqUID", rqUID).
		SetHeader("Authorization", "Bearer YmViNTRhMTgtZDhhYS00NGUwLTk4MGMtY2YyNDlkNWIyZGM5OjNlNDlmYzM2LTc5NWItNGY4ZC1iZjQwLWYzOTY0NDFkY2E2Yg==").
		SetFormData(map[string]string{
			"scope": "GIGACHAT_API_PERS",
		}).
		Post("https://ngw.devices.sberbank.ru:9443/api/v2/oauth")

	if err != nil {
		log.Fatalf("Ошибка получения токена: %v", err)
	}

	// Проверка статуса ответа
	if response.StatusCode() == 404 {
		log.Fatalf("Что-то не так")
	}
	if response.StatusCode() != 200 {
		log.Fatalf("Error: Received status code %d", response.StatusCode())
	}

	// Десериализация ответа в структуру
	type oauth struct {
		Access_token       string `json:"access_token"`
		Hexpires_atumidity int    `json:"expires_at"`
	}

	var w1 oauth

	err = json.Unmarshal(response.Body(), &w1)
	if err != nil {
		fmt.Println(err)
	}

	return w1.Access_token, nil
}
