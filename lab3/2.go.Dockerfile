# -------------------------------------------------- #
# 0
# -------------------------------------------------- #
FROM golang:1.22.1-alpine3.19 AS builder 

WORKDIR /app

COPY ./starter-project-golang/go.mod .
RUN go mod download

COPY ./starter-project-golang .
RUN go build -o build/fizzbuzz


# -------------------------------------------------- #
# 1
# -------------------------------------------------- #
FROM scratch
COPY --from=builder /app/build/fizzbuzz ./fizzbuzz

COPY ./starter-project-golang/templates/index.html ./templates/index.html

CMD ["./fizzbuzz", "serve"]
