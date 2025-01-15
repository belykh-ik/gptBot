FROM golang:latest

WORKDIR /gptBot

COPY . .

RUN wget https://gu-st.ru/content/lending/russian_trusted_sub_ca_pem.crt && cp russian_trusted_sub_ca_pem.crt /usr/local/share/ca-certificates/ && update-ca-certificates

CMD [ "go","run","." ]
