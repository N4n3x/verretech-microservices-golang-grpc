FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o ./gateway ./gateway

EXPOSE 10000

CMD [ "/gateway" ]ay/gateway.go ./
RUN go build -o main .

CMD ["/app/main"]



