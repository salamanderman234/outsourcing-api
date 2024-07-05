FROM golang:1.21.12

COPY . /app/

WORKDIR /app

RUN go build

RUN chmod +x ./outsourcing-api

EXPOSE 8080

CMD ["./outsourcing-api"]