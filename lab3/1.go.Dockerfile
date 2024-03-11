FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY ./starter-project-golang/go.mod .
RUN go mod download

COPY ./starter-project-golang .
RUN go build -o build/fizzbuzz

EXPOSE 8080

CMD ["./build/fizzbuzz", "serve"]
