FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN go build -o /myapp

EXPOSE 1323

CMD [ "/myapp" ]