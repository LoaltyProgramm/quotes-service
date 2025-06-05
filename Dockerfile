FROM golang:1.24.2

WORKDIR /app

COPY go.mod /app/

RUN go mod tidy

COPY cmd /app/cmd

COPY internal /app/internal

RUN GOARCH=amd64 GOOS=linux go build -o quotes_service ./cmd/main.go

CMD [ "./quotes_service" ]