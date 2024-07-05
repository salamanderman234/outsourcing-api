FROM golang:1.21.12

COPY . /app/

# RUN mv /app/build/.env .env
RUN cp /app/.env .env

RUN cd /app && go build /app/

RUN chmod +x /app/outsourcing-api

RUN /app/migrate

EXPOSE 8080

CMD ["/app/outsourcing-api"]