FROM golang:alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go ./

RUN go build -o /Go-Scrapper1

EXPOSE 3000

CMD ["/Go-Scrapper1"]

