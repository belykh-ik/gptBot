# 🤖 gptBot

A lightweight **Telegram bot** built in **Go** that connects to OpenAI's GPT API and responds to user messages with AI-generated answers — directly in chat!

Whether you're asking questions, chatting casually, or testing prompts, gptBot delivers GPT-powered responses within Telegram.

---

## ✨ Features

- 💬 Chat with GPT through Telegram  
- 🔐 Secure API key usage with environment variables  
- ⚡ Fast & lightweight (built in Go)  
- 🧠 AI replies powered by OpenAI GPT (supports gpt-3.5-turbo and others)  
- 🛠️ Easy to deploy with Docker or locally  

---

## 📌 Technologies

- **Go 1.18+**  
- **Telegram Bot API**  
- **OpenAI API** (ChatGPT)  
- **Docker (optional)**  

---

## 🚀 Getting Started

### 1. Prerequisites

- A **Telegram bot token** from [@BotFather](https://t.me/BotFather)  
- An **OpenAI API key** (https://platform.openai.com/account/api-keys)  
- Go installed (or Docker)

---

### 2. Clone the Repository

```bash
git clone https://github.com/belykh-ik/gptBot.git
cd gptBot
```
---

### 3. Set Environment Variables

```bash
export TELEGRAM_BOT_TOKEN="your_telegram_bot_token"
export OPENAI_API_KEY="your_openai_api_key"
```

### 4. Run the Bot

```bash
go run main.go
```
- **Or use Docker:**
```bash
docker build -t gptbot .
docker run --rm \
  -e TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN \
  -e OPENAI_API_KEY=$OPENAI_API_KEY \
  gptbot
```

