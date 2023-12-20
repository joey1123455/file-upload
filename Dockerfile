FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 8000

CMD ["./file-upload"]