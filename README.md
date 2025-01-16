# Бот на основе API Sber GigaChat

**Для запуска бота необходимо установить 2 переменных окружения**

1. ***TOKEN*** - Токен вашего бота.
2. ***KEY*** - Authorization key от GigaChat API 

***Пример установки переменной на один сеанс***

```
export TOKEN=<Ваш токен>
export TOKEN=<Ваш Authorization key>
```

***Необходимо собрать образ и запустить контейнер***

***Пример сборки и запуска***

```
docker build -t gptbot .
docker run -d gptbot
```