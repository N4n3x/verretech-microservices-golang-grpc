FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#TODO trouver le bon exec

EXPOSE 4747
