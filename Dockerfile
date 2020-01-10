FROM golang:latest
LABEL maintainer="Adam Thomason <adam.atdl@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 80

CMD ["./main"]