# syntax=docker/dockerfile:1

FROM golang:1.17.2

#makes dir inside the image
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /hello-app

CMD [ "/hello-app" ]

