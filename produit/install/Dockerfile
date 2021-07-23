
FROM golang:latest
RUN mkdir -p /go/src/verretech-microservices/install
RUN CGO_ENABLED=0

ADD . /go/src/verretech-microservices/install
ENV GOPATH /go
WORKDIR /go/src/verretech-microservices/install


RUN go get
RUN go build

FROM alpine
RUN apk update
RUN apk add curl

RUN mkdir -p /app
COPY --from=builder /go/src/verretech-microservices/install /app/
RUN chmod +x /app/api
WORKDIR /app
EXPOSE 5001
ENTRYPOINT ["/app/api"]