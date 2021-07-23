FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .


RUN go mod download

#go build -a -installsuffix cgo -o /go/bin/cimongo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 
CMD [ "app" ]
EXPOSE 4747
