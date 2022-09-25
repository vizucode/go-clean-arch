FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o echo-boilerplate

CMD ["./echo-boilerplate"]
