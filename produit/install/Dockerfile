FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

EXPOSE 4747
