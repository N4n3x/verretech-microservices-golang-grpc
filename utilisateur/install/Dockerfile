FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/cimongo
#TODO trouver le bon exec
ENTRYPOINT [ "./verretech-microservices/utilisateur" ]
EXPOSE 4747
