FROM golang:latest

RUN mkdir /app
RUN CGO_ENABLED=0
ADD . /app
ENV GOPATH /go
WORKDIR /app
COPY . .
RUN go get
RUN go build

FROM alpine
RUN apk update
RUN apk add curl

RUN mkdir -p /app
COPY --from=builder /go/src/verretech-microservices/produit/ /app/
RUN chmod +x /app/api
WORKDIR /app
EXPOSE 4747
ENTRYPOINT ["/app/api"]