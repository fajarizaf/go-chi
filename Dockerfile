FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o hello-go

EXPOSE 8080

CMD ./hello-go