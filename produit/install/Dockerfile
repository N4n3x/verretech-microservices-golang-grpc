FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .


RUN go mod download


RUN CGO_ENABLED=0 go build -o produit .
CMD [ "./produit"]
EXPOSE 4747
