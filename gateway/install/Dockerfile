FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

#telechargement des dépendances
COPY go.mod ./
COPY go.sum ./
RUN  go mod download

#Copie de l'executable
COPY /gateway/gateway.go ./
RUN go build -o main .

CMD ["/app/main"]



