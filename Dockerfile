FROM golang:1.18

LABEL base.name="Groupie tracker"

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/

EXPOSE 8080

ENTRYPOINT ["./main"]
