FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/cimongo
ENTRYPOINT [ "/go/bin/cimongo" ]
EXPOSE 4747
