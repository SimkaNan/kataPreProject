FROM golang:1.21.1-alpine

WORKDIR /usr/local/src

COPY ./ ./

RUN go mod init go-app && go get github.com/joho/godotenv
RUN go build -o main .
CMD ["./main"]