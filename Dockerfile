FROM golang:1.21.12

COPY . /app/

WORKDIR /app

RUN mv .env .env.dev

RUN mv .env.production .env

RUN go build

RUN chmod +x ./outsourcing-api

EXPOSE 8080

CMD ["./outsourcing-api"]