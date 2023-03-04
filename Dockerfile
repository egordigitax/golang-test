FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./cmd/app

EXPOSE 8000

CMD [ "./app" ]